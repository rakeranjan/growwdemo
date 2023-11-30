`
{
	"info": {
		"_postman_id": "e92b1481-2f3f-4304-af9f-6270bc92bbb9",
		"name": "groww",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
		"_exporter_id": "3136197"
	},
	"item": [
		{
			"name": "New Request",
			"request": {
				"method": "POST",
				"header": [],
				"url": {
					"raw": "localhost:8000/addSeller",
					"host": [
						"localhost"
					],
					"port": "8000",
					"path": [
						"addSeller"
					]
				}
			},
			"response": []
		},
		{
			"name": "New Request",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\n    \"sellerId\": \"eea24fde-454f-4d32-b807-5e3c5dfe3f33\",\n    \"name\": \"notepad\",\n    \"price\": 123\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "localhost:8000/listItem",
					"host": [
						"localhost"
					],
					"port": "8000",
					"path": [
						"listItem"
					]
				}
			},
			"response": []
		},
		{
			"name": "New Request",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "localhost:8000/getAllListing",
					"host": [
						"localhost"
					],
					"port": "8000",
					"path": [
						"getAllListing"
					]
				}
			},
			"response": []
		},
		{
			"name": "New Request",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\n    \"sellerId\": \"5fd6e323-8a21-4672-830e-d813be883c8c\"\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "localhost:8000/getSellerListing",
					"host": [
						"localhost"
					],
					"port": "8000",
					"path": [
						"getSellerListing"
					]
				}
			},
			"response": []
		},
		{
			"name": "New Request",
			"request": {
				"method": "POST",
				"header": [],
				"url": {
					"raw": "localhost:8000/addBuyer",
					"host": [
						"localhost"
					],
					"port": "8000",
					"path": [
						"addBuyer"
					]
				}
			},
			"response": []
		},
		{
			"name": "New Request",
			"request": {
				"method": "GET",
				"header": []
			},
			"response": []
		}
	]
}
`
