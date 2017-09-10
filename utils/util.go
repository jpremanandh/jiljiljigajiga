package utils

import (
	"fmt"
	"math"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var repeatingWhitespace *regexp.Regexp

func init() {
	repeatingWhitespace = regexp.MustCompile("\\s+")
}

/*
Substitute method substitues the %{..} placeholders with actual value from the
*/
func Substitute(baseString string, data map[string]interface{}) (string, error) {
	value := baseString
	re := regexp.MustCompile("%{.+?}")

	for match := re.FindString(value); match != ""; match = re.FindString(value) {
		jsonPath := match[2 : len(match)-1]
		jsonValue := JSONValue(data, jsonPath)

		if jsonValue == nil {
			return "", fmt.Errorf("Unable to find key %s", jsonPath)
		}
		value = strings.Replace(value, match, fmt.Sprintf("%v", jsonValue), -1)
	}

	return value, nil
}

/*
EventJSONValue method returns the JSON value from the given event either Data or Metadata based on path.
*/
// func EventJSONValue(event *common.Event, jsonPath string) interface{} {
// 	splits := strings.Split(jsonPath, ".")
// 	if len(splits) > 0 && splits[0] == "@metadata" {
// 		var modifiedJsonPath string
// 		if len(splits) > 1 {
// 			modifiedJsonPath = strings.Join(splits[1:], ".")
// 		}
// 		return JSONValue(event.Metadata, modifiedJsonPath)
// 	} else {
// 		return JSONValue(event.Data, jsonPath)
// 	}
// }

/*
JSONValue method returns the JSON value from the given map, for the given json path
*/
func JSONValue(data map[string]interface{}, jsonPath string) interface{} {
	if data == nil {
		return nil
	}
	if jsonPath == "" {
		return data
	}

	var lookup map[string]interface{}
	var mapValue interface{}
	var ok bool

	lookup = data
	splits := strings.Split(jsonPath, ".")
	for _, split := range splits {
		mapValue, ok = lookup[split]
		if !ok {
			return nil
		}

		kind := reflect.ValueOf(mapValue).Kind()
		if kind == reflect.Map {
			lookup = mapValue.(map[string]interface{})
		}
	}

	return mapValue
}

/*
ConvertPathToArray method returns the JSON value from the given map, for the given json path
*/
func ConvertPathToArray(data map[string]interface{}, jsonPath string) interface{} {
	if data == nil {
		return nil
	}

	if jsonPath == "" {
		return data
	}

	var lookup map[string]interface{}
	var mapValue interface{}
	var ok bool

	lookup = data
	splits := strings.Split(jsonPath, ".")
	split := splits[0]
	mapValue, ok = lookup[split]
	if !ok {
		return data
	}

	if len(splits) == 1 {
		var convertedData interface{}
		if mapValue != nil && reflect.ValueOf(mapValue).Kind() != reflect.Array && reflect.ValueOf(mapValue).Kind() != reflect.Slice {
			convertedDataArray := make([]interface{}, 1)
			convertedDataArray[0] = mapValue
			convertedData = convertedDataArray
		} else {
			convertedData = mapValue
		}
		lookup[split] = convertedData
		return lookup
	}

	kind := reflect.ValueOf(mapValue).Kind()
	if kind == reflect.Map {
		data := ConvertPathToArray(mapValue.(map[string]interface{}), strings.Join(splits[1:len(splits)], "."))
		lookup[split] = data
	} else if kind == reflect.Slice {
		mapValueArray := mapValue.([]interface{})
		for index, interfaceValue := range mapValueArray {
			if reflect.Map == reflect.ValueOf(interfaceValue).Kind() {
				data := ConvertPathToArray(interfaceValue.(map[string]interface{}), strings.Join(splits[1:len(splits)], "."))
				mapValueArray[index] = data
			}
		}
		lookup[split] = mapValueArray
	}
	return lookup
}

/*
JSONSetValue method sets the given value into the data JSON in the given json path
*/
func JSONSetValue(data map[string]interface{}, jsonPath string, value interface{}) {
	var lookup map[string]interface{}
	var mapValue interface{}
	var ok bool

	lookup = data
	splits := strings.Split(jsonPath, ".")
	for index, split := range splits {
		if index == len(splits)-1 {
			break
		}

		mapValue, ok = lookup[split]
		if !ok {
			lookup[split] = map[string]interface{}{}
			mapValue = lookup[split]
		}

		if reflect.ValueOf(mapValue).Kind() == reflect.Map {
			lookup = mapValue.(map[string]interface{})
		}
	}

	lookup[splits[len(splits)-1]] = value
}

func Contains(slice []int, element int) bool {
	if slice == nil {
		return false
	}

	for _, value := range slice {
		if value == element {
			return true
		}
	}

	return false
}

func StringOf(value interface{}) string {
	if value == nil {
		return ""
	}

	var text string
	kind := reflect.ValueOf(value).Kind()
	switch kind {
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		fallthrough
	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		text = fmt.Sprintf("%d", value)
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		text = fmt.Sprintf("%.2f", value)
	case reflect.String:
		text = value.(string)
	case reflect.Bool:
		text = strconv.FormatBool(value.(bool))
	}

	return strings.Trim(text, " ")
}

// func FillErrorIntoArray(events []*common.Event, errArray []error, errData error) []error {
// 	for index, _ := range errArray {
// 		if events[index].State == common.Skip {
// 			continue
// 		}
// 		if errArray[index] == nil {
// 			errArray[index] = errData
// 		}
// 	}
// 	return errArray
// }

/*
ToSnakeCase converts the input string to snake case
*/
func ToSnakeCase(in string) string {
	out := repeatingWhitespace.ReplaceAllString(in, "_")
	out = strings.ToLower(out)

	return out
}

func SplitArray(arrayTosplit []interface{}, chunkSize int) [][]interface{} {
	var splittedArray = [][]interface{}{}
	var arrayLen = len(arrayTosplit)
	if chunkSize > arrayLen {
		return [][]interface{}{arrayTosplit}
	}

	for i := 0; i < len(arrayTosplit); i += chunkSize {
		end := i + chunkSize
		if end > arrayLen {
			end = arrayLen
		}
		splittedArray = append(splittedArray, arrayTosplit[i:end])
	}
	return splittedArray
}

// func SplitEvents(arrayTosplit []*common.Event, chunkSize int) [][]*common.Event {
// 	var splittedArray = [][]*common.Event{}
// 	var arrayLen = len(arrayTosplit)
// 	if chunkSize > arrayLen {
// 		return [][]*common.Event{arrayTosplit}
// 	}

// 	for i := 0; i < len(arrayTosplit); i += chunkSize {
// 		end := i + chunkSize
// 		if end > arrayLen {
// 			end = arrayLen
// 		}
// 		splittedArray = append(splittedArray, arrayTosplit[i:end])
// 	}
// 	return splittedArray
// }

func GenerateMapper(inputMapper map[string]interface{}, f func(data interface{}) interface{}) map[string]interface{} {
	var outputMapper = map[string]interface{}{}
	for ikey, ivalue := range inputMapper {
		sliceData, sliceOk := ivalue.([]interface{})
		mapData, mapOk := ivalue.(map[string]interface{})
		if sliceOk {
			outputMapper[ikey] = GenerateArray(sliceData, f)
		} else if mapOk {
			outputMapper[ikey] = GenerateMapper(mapData, f)
		} else {
			outputMapper[ikey] = f(ivalue)
		}
	}
	return outputMapper
}

func GenerateArray(inputArray []interface{}, f func(data interface{}) interface{}) []interface{} {
	var outputArray = make([]interface{}, len(inputArray))
	for i, aVal := range inputArray {
		sliceData, sliceOk := aVal.([]interface{})
		mapData, mapOk := aVal.(map[string]interface{})
		if sliceOk {
			outputArray[i] = GenerateArray(sliceData, f)
		} else if mapOk {
			outputArray[i] = GenerateMapper(mapData, f)
		} else {
			outputArray[i] = f(aVal)
		}
	}
	return outputArray
}

/*
fmt.Println(Round(19.9967001, .5, 2))
20
fmt.Println(Round(19.9917001, .5, 2))
19.99
*/

func Round(val, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

/*
WithinAllowedDiff(19.9967001, 19.9917001, .01)
*/
func WithinAllowedDiff(val1, val2, diffAllowed float64) bool {
	diff := (val1 - val2)
	if diff < 0 {
		diff = -diff
	}
	return (Truncate(diff, 2) <= diffAllowed)
}

func Truncate(some float64, places int) float64 {
	pow := math.Pow(10, float64(places))
	return float64(int(some*pow)) / float64(int(pow))
}
