curl -X POST \
  http://us-west-bidder.mathtag.com:9180/bid/gor \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -H 'postman-token: 46378f59-7a1e-3c39-582d-b646b32980e9' \
  -d '{
        "id": "83187405165",
        "at": 2,
        "imp": [{
                "id": "2d567fb4ab21",
                "banner": {
                        "w": 300,
                        "h": 250,
                        "pos": 0,
                        "mimes": ["image/gif", "image/jpeg", "image/png", "text/javascript"],
                        "name": "AUM"
                },
                "instl": 0,
                "ext": {
                        "nex_screen": 0
                }
        }],
        "site": {
                "name": "BidderTestMobileWEB",
                "domain": "atlassian.com",
                "cat": ["IAB3"],
                "page": "http://www.atlassian.com/mobile",
                "ref": "http://www.iab.net",
                "search": "radiation",
                "publisher": {
                        "id": "123",
                        "name": "Hind"
                },
                "keywords": "radiation"
        },
        "device": {
                "ua": "Mozilla/5.0 (iPhone; U; CPU iPhone OS 4_2_1 like Mac OS X; el-gr) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8C148 Safari/6533.18.5",
                "ip": "72.229.28.185",
                "geo": {
                        "lat": 40.7605,
                        "lon": -73.9933,
                        "country": "TUR",
                        "region": "NY",
                        "city": "New York",
                        "utcoffset": -300
                },
                "didsha1": "132079238ec783b0b89dff308e1f9bdd08576273",
                "dpidsha1": "f22711a823044bb9ce7ace097955de0286eb0182",
                "carrier": "ATT",
                "connectiontype": 3,
                "devicetype": 2
        },
        "ext": {
                "ssp": "SelfService1"
        }
}'

# Response for above request
# {
#     "id": "83187405165",
#     "seatbid": [
#         {
#             "seat": "101209",
#             "bid": [
#                 {
#                     "id": "7234462437741189118",
#                     "impid": "2d567fb4ab21",
#                     "price": 0.007,
#                     "adid": "4734078",
#                     "adm": "<script language='JavaScript' src='https://tags.mathtag.com/notify/js?exch=gor&id=5aW95q2jLzEzLyAvWmpJeU56RXhZVGd0TWpNd05DMDBZbUk1TFdObE4yRXRZMlV3T1RjNU5UVmtaVEF5LzcyMzQ0NjI0Mzc3NDExODkxMTgvNDczNDA3OC8yNjA0NjEzLzU3L0puVHVkSVBDTTZIWDA1M2NhdW1CY0JacFdxVl9fLVVUMUFMeEZoZW8yUXMvMS81Ny8xNTA0NzE3MTc1LzAvNDkwMzg0LzEyMjI5NzQ2NDkvMTk0Mzg4LzQwMTQ0Ni80LzAvMC9NREF3TURBd01EQXRNREF3TUMwd01EQXdMVEF3TURBdE1EQXdNREF3TURBd01EQXcvMC8wLzAvMC8wLw/DCxRzoVM7vrq6jvq29iR0JQmZk0&sid=2604613&cid=4734078&nodeid=401&price=${AUCTION_PRICE}&group=pao&auctionid=7234462437741189118&bid=pao&pbs_id=7234462437741189118&bp=a_aagjfh'></script>",
#                     "adomain": [
#                         "iwitnessbullying.org"
#                     ],
#                     "cid": "401446",
#                     "crid": "4734078"
#                 }
#             ]
#         },
#         {
#             "seat": "101117",
#             "bid": [
#                 {
#                     "id": "7234462437741189118",
#                     "impid": "2d567fb4ab21",
#                     "price": 0.007,
#                     "adid": "4725160",
#                     "adm": "<script language='JavaScript' src='https://tags.mathtag.com/notify/js?exch=gor&id=5aW95q2jLzEzLyAvWmpJeU56RXhZVGd0TWpNd05DMDBZbUk1TFdObE4yRXRZMlV3T1RjNU5UVmtaVEF5LzcyMzQ1MDE0NzM2MjUyMDA2MzgvNDcyNTE2MC8yNjAwMTE4LzU3L1BadDFBOHBFYmhWbzl4SHhwblF3MW4xc1hNeUVzc3JhME9Gbk9PMExEb0EvMS81Ny8xNTA0NTQ1MjM2LzAvNTEyMTk3LzEyMjI5NzQ2NDkvMTQ4NzUwLzQxMTUyNy80LzAvMC9NREF3TURBd01EQXRNREF3TUMwd01EQXdMVEF3TURBdE1EQXdNREF3TURBd01EQXcvMC8wLzAvMC8wLw/lfwhG2foMhEJyc-9HNDhu3O_qAU&sid=2600118&cid=4725160&nodeid=401&price=${AUCTION_PRICE}&group=pao&auctionid=7234501473625200638&bid=pao&pbs_id=7234462437741189118&bp=a_aagidi'></script>",
#                     "adomain": [
#                         "pmlatam.com"
#                     ],
#                     "cid": "411527",
#                     "crid": "4725160"
#                 }
#             ]
#         }
#     ],
#     "bidid": "7234462437741189118",
#     "cur": "USD"
# }