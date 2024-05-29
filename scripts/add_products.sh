#!/bin/bash

# JSON array of products
products='[
    {"name":"T-Shirt","type":"Clothing","picture":"https://loremflickr.com/600/400/tshirt","price":19.99,"description":"Comfortable cotton T-Shirt"},
    {"name":"Jeans","type":"Clothing","picture":"https://loremflickr.com/600/500/jeans","price":49.99,"description":"Durable denim jeans"},
    {"name":"Sneakers","type":"Shoes","picture":"https://loremflickr.com/600/600/sneakers","price":59.99,"description":"Comfortable and stylish sneakers"},
    {"name":"Hoodie","type":"Clothing","picture":"https://loremflickr.com/600/500/hoodie","price":39.99,"description":"Warm and cozy hoodie"},
    {"name":"Dress","type":"Clothing","picture":"https://loremflickr.com/600/450/dress","price":69.99,"description":"Elegant and stylish dress"},
    {"name":"Skirt","type":"Clothing","picture":"https://loremflickr.com/600/400/skirt","price":39.99,"description":"Versatile and stylish skirt"},
    {"name":"Blouse","type":"Clothing","picture":"https://loremflickr.com/600/350/blouse","price":34.99,"description":"Comfortable and stylish blouse"},
    {"name":"Sweater","type":"Clothing","picture":"https://loremflickr.com/600/500/sweater","price":49.99,"description":"Warm and cozy sweater"},
    {"name":"Coat","type":"Clothing","picture":"https://loremflickr.com/600/600/coat","price":99.99,"description":"Warm and stylish coat"},
    {"name":"Boots","type":"Shoes","picture":"https://loremflickr.com/600/550/boots","price":79.99,"description":"Durable and stylish boots"},
    {"name":"Hat","type":"Accessories","picture":"https://loremflickr.com/600/450/hat","price":24.99,"description":"Stylish and functional hat"},
    {"name":"Scarf","type":"Accessories","picture":"https://loremflickr.com/600/400/scarf","price":19.99,"description":"Warm and stylish scarf"},
    {"name":"Gloves","type":"Accessories","picture":"https://loremflickr.com/600/350/gloves","price":29.99,"description":"Warm and stylish gloves"},
    {"name":"Sunglasses","type":"Accessories","picture":"https://loremflickr.com/600/500/sunglasses","price":39.99,"description":"Stylish and functional sunglasses"},
    {"name":"Bag","type":"Accessories","picture":"https://loremflickr.com/600/400/bag","price":59.99,"description":"Stylish and functional bag"},
    {"name":"Wallet","type":"Accessories","picture":"https://loremflickr.com/600/350/wallet","price":29.99,"description":"Stylish and functional wallet"},
    {"name":"Belt","type":"Accessories","picture":"https://loremflickr.com/600/400/belt","price":24.99,"description":"Stylish and functional belt"},
    {"name":"Jewelry","type":"Accessories","picture":"https://loremflickr.com/600/500/jewelry","price":49.99,"description":"Elegant and stylish jewelry"}
]'

# Loop through each product in the JSON array
for row in $(echo "${products}" | jq -c '.[]'); do
    # Send a POST request for each product
	# Do not forget to check the endpoint URL if it is correct with your local setup
    curl -X POST http://localhost:8080/products -H "Content-Type: application/json" -d "${row}"
    echo "Added product: ${row}"
done
