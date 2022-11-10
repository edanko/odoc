## Paths

**GET** /store/inventory

Returns a map of status codes to quantities

### Request parameters

### Responses

200 OK

Content type: application/json

body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|


---

**POST** /user/createWithList

Creates list of users with given input array

### Request parameters

Content type: application/json

Body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| ||body|array/object/|||
| |lastName|body|string//|||
| |password|body|string//|||
| |phone|body|string//|||
| |userStatus|body|integer//int32|User Status||
| |username|body|string//|||
| |email|body|string//|||
| |firstName|body|string//|||
| |id|body|integer//int64|||


### Responses

200 OK: Successful operation

Content type: application/json

body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |userStatus|body|integer//int32|User Status||
| |username|body|string//|||
| |email|body|string//|||
| |firstName|body|string//|||
| |id|body|integer//int64|||
| |lastName|body|string//|||
| |password|body|string//|||
| |phone|body|string//|||


---

**POST** /user/createWithList

Creates list of users with given input array

### Request parameters

Content type: application/json

Body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| ||body|array/object/|||
| |username|body|string//|||
| |email|body|string//|||
| |firstName|body|string//|||
| |id|body|integer//int64|||
| |lastName|body|string//|||
| |password|body|string//|||
| |phone|body|string//|||
| |userStatus|body|integer//int32|User Status||


### Responses

200 OK: Successful operation

Content type: application/xml

body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |password|body|string//|||
| |phone|body|string//|||
| |userStatus|body|integer//int32|User Status||
| |username|body|string//|||
| |email|body|string//|||
| |firstName|body|string//|||
| |id|body|integer//int64|||
| |lastName|body|string//|||


---

**GET** /user/login

### Request parameters

query

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |username|query|string//|||
| |password|query|string//|||


### Responses

200 OK

Content type: application/json

headers

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |X-Expires-After||string//date-time|||
| |X-Rate-Limit||integer//int32|||


body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| ||body|string//|||


---

**GET** /user/login

### Request parameters

query

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |username|query|string//|||
| |password|query|string//|||


### Responses

200 OK

Content type: application/xml

headers

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |X-Expires-After||string//date-time|||
| |X-Rate-Limit||integer//int32|||


body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| ||body|string//|||


---

**GET** /pet/findByStatus

Multiple status values can be provided with comma separated strings

### Request parameters

query

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |status|query|string//||available, pending, sold|


### Responses

200 OK

Content type: application/xml

body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| ||body|array/object/|||
| |status|body|string//|pet status in the store|available, pending, sold|
| |tags|body|array/object/|||
| |category|body|object//|||
| |id|body|integer//int64|||
| |name|body|string//|||
| |photoUrls|body|array/string/|||


---

**GET** /pet/findByStatus

Multiple status values can be provided with comma separated strings

### Request parameters

query

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |status|query|string//||available, pending, sold|


### Responses

200 OK

Content type: application/json

body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| ||body|array/object/|||
| |id|body|integer//int64|||
| |name|body|string//|||
| |photoUrls|body|array/string/|||
| |status|body|string//|pet status in the store|available, pending, sold|
| |tags|body|array/object/|||
| |category|body|object//|||


---

**GET** /pet/findByTags

Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.

### Request parameters

query

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |tags|query|array/string/|||


### Responses

200 OK

Content type: application/xml

body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| ||body|array/object/|||
| |id|body|integer//int64|||
| |name|body|string//|||
| |photoUrls|body|array/string/|||
| |status|body|string//|pet status in the store|available, pending, sold|
| |tags|body|array/object/|||
| |category|body|object//|||


---

**GET** /pet/findByTags

Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.

### Request parameters

query

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |tags|query|array/string/|||


### Responses

200 OK

Content type: application/json

body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| ||body|array/object/|||
| |category|body|object//|||
| |id|body|integer//int64|||
| |name|body|string//|||
| |photoUrls|body|array/string/|||
| |status|body|string//|pet status in the store|available, pending, sold|
| |tags|body|array/object/|||


---

**GET** /pet/{petId}

Returns a single pet

### Request parameters

Path params

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |petId|path|integer//int64|||


### Responses

200 OK

Content type: application/json

body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |photoUrls|body|array/string/|||
| |status|body|string//|pet status in the store|available, pending, sold|
| |tags|body|array/object/|||
| |category|body|object//|||
| |id|body|integer//int64|||
| |name|body|string//|||


---

**GET** /pet/{petId}

Returns a single pet

### Request parameters

Path params

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |petId|path|integer//int64|||


### Responses

200 OK

Content type: application/xml

body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |category|body|object//|||
| |id|body|integer//int64|||
| |name|body|string//|||
| |photoUrls|body|array/string/|||
| |status|body|string//|pet status in the store|available, pending, sold|
| |tags|body|array/object/|||


---

**POST** /pet/{petId}/uploadImage

### Request parameters

Content type: application/octet-stream

Path params

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |petId|path|integer//int64|||


query

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |additionalMetadata|query|string//|||


Body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| ||body|string//binary|||


### Responses

200 OK: successful operation

Content type: application/json

body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |message|body|string//|||
| |type|body|string//|||
| |code|body|integer//int32|||


---

**POST** /store/order

Place a new order in the store

### Request parameters

Content type: application/json

Body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |complete|body|boolean//|||
| |id|body|integer//int64|||
| |petId|body|integer//int64|||
| |quantity|body|integer//int32|||
| |shipDate|body|string//date-time|||
| |status|body|string//|Order Status|placed, approved, delivered|


### Responses

200 OK: successful operation

Content type: application/json

body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |petId|body|integer//int64|||
| |quantity|body|integer//int32|||
| |shipDate|body|string//date-time|||
| |status|body|string//|Order Status|placed, approved, delivered|
| |complete|body|boolean//|||
| |id|body|integer//int64|||


---

**POST** /store/order

Place a new order in the store

### Request parameters

Content type: application/x-www-form-urlencoded

Body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |complete|body|boolean//|||
| |id|body|integer//int64|||
| |petId|body|integer//int64|||
| |quantity|body|integer//int32|||
| |shipDate|body|string//date-time|||
| |status|body|string//|Order Status|placed, approved, delivered|


### Responses

200 OK: successful operation

Content type: application/json

body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |status|body|string//|Order Status|placed, approved, delivered|
| |complete|body|boolean//|||
| |id|body|integer//int64|||
| |petId|body|integer//int64|||
| |quantity|body|integer//int32|||
| |shipDate|body|string//date-time|||


---

**POST** /store/order

Place a new order in the store

### Request parameters

Content type: application/xml

Body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |petId|body|integer//int64|||
| |quantity|body|integer//int32|||
| |shipDate|body|string//date-time|||
| |status|body|string//|Order Status|placed, approved, delivered|
| |complete|body|boolean//|||
| |id|body|integer//int64|||


### Responses

200 OK: successful operation

Content type: application/json

body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |shipDate|body|string//date-time|||
| |status|body|string//|Order Status|placed, approved, delivered|
| |complete|body|boolean//|||
| |id|body|integer//int64|||
| |petId|body|integer//int64|||
| |quantity|body|integer//int32|||


---

**GET** /store/order/{orderId}

For valid response try integer IDs with value <= 5 or > 10. Other values will generate exceptions.

### Request parameters

Path params

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |orderId|path|integer//int64|||


### Responses

200 OK

Content type: application/xml

body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |petId|body|integer//int64|||
| |quantity|body|integer//int32|||
| |shipDate|body|string//date-time|||
| |status|body|string//|Order Status|placed, approved, delivered|
| |complete|body|boolean//|||
| |id|body|integer//int64|||


---

**GET** /store/order/{orderId}

For valid response try integer IDs with value <= 5 or > 10. Other values will generate exceptions.

### Request parameters

Path params

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |orderId|path|integer//int64|||


### Responses

200 OK

Content type: application/json

body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |complete|body|boolean//|||
| |id|body|integer//int64|||
| |petId|body|integer//int64|||
| |quantity|body|integer//int32|||
| |shipDate|body|string//date-time|||
| |status|body|string//|Order Status|placed, approved, delivered|


---

**POST** /user

This can only be done by the logged in user.

### Request parameters

Content type: application/json

Body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |id|body|integer//int64|||
| |lastName|body|string//|||
| |password|body|string//|||
| |phone|body|string//|||
| |userStatus|body|integer//int32|User Status||
| |username|body|string//|||
| |email|body|string//|||
| |firstName|body|string//|||


### Responses

200 OK: successful operation

Content type: application/json

body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |username|body|string//|||
| |email|body|string//|||
| |firstName|body|string//|||
| |id|body|integer//int64|||
| |lastName|body|string//|||
| |password|body|string//|||
| |phone|body|string//|||
| |userStatus|body|integer//int32|User Status||


---

**POST** /user

This can only be done by the logged in user.

### Request parameters

Content type: application/x-www-form-urlencoded

Body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |id|body|integer//int64|||
| |lastName|body|string//|||
| |password|body|string//|||
| |phone|body|string//|||
| |userStatus|body|integer//int32|User Status||
| |username|body|string//|||
| |email|body|string//|||
| |firstName|body|string//|||


### Responses

200 OK: successful operation

Content type: application/json

body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |lastName|body|string//|||
| |password|body|string//|||
| |phone|body|string//|||
| |userStatus|body|integer//int32|User Status||
| |username|body|string//|||
| |email|body|string//|||
| |firstName|body|string//|||
| |id|body|integer//int64|||


---

**POST** /user

This can only be done by the logged in user.

### Request parameters

Content type: application/xml

Body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |lastName|body|string//|||
| |password|body|string//|||
| |phone|body|string//|||
| |userStatus|body|integer//int32|User Status||
| |username|body|string//|||
| |email|body|string//|||
| |firstName|body|string//|||
| |id|body|integer//int64|||


### Responses

200 OK: successful operation

Content type: application/json

body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |phone|body|string//|||
| |userStatus|body|integer//int32|User Status||
| |username|body|string//|||
| |email|body|string//|||
| |firstName|body|string//|||
| |id|body|integer//int64|||
| |lastName|body|string//|||
| |password|body|string//|||


---

**POST** /user

This can only be done by the logged in user.

### Request parameters

Content type: application/x-www-form-urlencoded

Body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |username|body|string//|||
| |email|body|string//|||
| |firstName|body|string//|||
| |id|body|integer//int64|||
| |lastName|body|string//|||
| |password|body|string//|||
| |phone|body|string//|||
| |userStatus|body|integer//int32|User Status||


### Responses

200 OK: successful operation

Content type: application/xml

body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |email|body|string//|||
| |firstName|body|string//|||
| |id|body|integer//int64|||
| |lastName|body|string//|||
| |password|body|string//|||
| |phone|body|string//|||
| |userStatus|body|integer//int32|User Status||
| |username|body|string//|||


---

**POST** /user

This can only be done by the logged in user.

### Request parameters

Content type: application/xml

Body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |password|body|string//|||
| |phone|body|string//|||
| |userStatus|body|integer//int32|User Status||
| |username|body|string//|||
| |email|body|string//|||
| |firstName|body|string//|||
| |id|body|integer//int64|||
| |lastName|body|string//|||


### Responses

200 OK: successful operation

Content type: application/xml

body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |userStatus|body|integer//int32|User Status||
| |username|body|string//|||
| |email|body|string//|||
| |firstName|body|string//|||
| |id|body|integer//int64|||
| |lastName|body|string//|||
| |password|body|string//|||
| |phone|body|string//|||


---

**POST** /user

This can only be done by the logged in user.

### Request parameters

Content type: application/json

Body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |email|body|string//|||
| |firstName|body|string//|||
| |id|body|integer//int64|||
| |lastName|body|string//|||
| |password|body|string//|||
| |phone|body|string//|||
| |userStatus|body|integer//int32|User Status||
| |username|body|string//|||


### Responses

200 OK: successful operation

Content type: application/xml

body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |phone|body|string//|||
| |userStatus|body|integer//int32|User Status||
| |username|body|string//|||
| |email|body|string//|||
| |firstName|body|string//|||
| |id|body|integer//int64|||
| |lastName|body|string//|||
| |password|body|string//|||


---

**GET** /user/{username}

### Request parameters

Path params

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |username|path|string//|||


### Responses

200 OK

Content type: application/json

body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |phone|body|string//|||
| |userStatus|body|integer//int32|User Status||
| |username|body|string//|||
| |email|body|string//|||
| |firstName|body|string//|||
| |id|body|integer//int64|||
| |lastName|body|string//|||
| |password|body|string//|||


---

**GET** /user/{username}

### Request parameters

Path params

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |username|path|string//|||


### Responses

200 OK

Content type: application/xml

body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |email|body|string//|||
| |firstName|body|string//|||
| |id|body|integer//int64|||
| |lastName|body|string//|||
| |password|body|string//|||
| |phone|body|string//|||
| |userStatus|body|integer//int32|User Status||
| |username|body|string//|||


---

**POST** /pet

Add a new pet to the store

### Request parameters

Content type: application/json

Body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |status|body|string//|pet status in the store|available, pending, sold|
| |tags|body|array/object/|||
| |category|body|object//|||
| |id|body|integer//int64|||
| |name|body|string//|||
| |photoUrls|body|array/string/|||


### Responses

200 OK: Successful operation

Content type: application/json

body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |status|body|string//|pet status in the store|available, pending, sold|
| |tags|body|array/object/|||
| |category|body|object//|||
| |id|body|integer//int64|||
| |name|body|string//|||
| |photoUrls|body|array/string/|||


---

**POST** /pet

Add a new pet to the store

### Request parameters

Content type: application/x-www-form-urlencoded

Body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |name|body|string//|||
| |photoUrls|body|array/string/|||
| |status|body|string//|pet status in the store|available, pending, sold|
| |tags|body|array/object/|||
| |category|body|object//|||
| |id|body|integer//int64|||


### Responses

200 OK: Successful operation

Content type: application/json

body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |photoUrls|body|array/string/|||
| |status|body|string//|pet status in the store|available, pending, sold|
| |tags|body|array/object/|||
| |category|body|object//|||
| |id|body|integer//int64|||
| |name|body|string//|||


---

**POST** /pet

Add a new pet to the store

### Request parameters

Content type: application/xml

Body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |category|body|object//|||
| |id|body|integer//int64|||
| |name|body|string//|||
| |photoUrls|body|array/string/|||
| |status|body|string//|pet status in the store|available, pending, sold|
| |tags|body|array/object/|||


### Responses

200 OK: Successful operation

Content type: application/json

body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |category|body|object//|||
| |id|body|integer//int64|||
| |name|body|string//|||
| |photoUrls|body|array/string/|||
| |status|body|string//|pet status in the store|available, pending, sold|
| |tags|body|array/object/|||


---

**POST** /pet

Add a new pet to the store

### Request parameters

Content type: application/x-www-form-urlencoded

Body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |category|body|object//|||
| |id|body|integer//int64|||
| |name|body|string//|||
| |photoUrls|body|array/string/|||
| |status|body|string//|pet status in the store|available, pending, sold|
| |tags|body|array/object/|||


### Responses

200 OK: Successful operation

Content type: application/xml

body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |photoUrls|body|array/string/|||
| |status|body|string//|pet status in the store|available, pending, sold|
| |tags|body|array/object/|||
| |category|body|object//|||
| |id|body|integer//int64|||
| |name|body|string//|||


---

**POST** /pet

Add a new pet to the store

### Request parameters

Content type: application/xml

Body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |category|body|object//|||
| |id|body|integer//int64|||
| |name|body|string//|||
| |photoUrls|body|array/string/|||
| |status|body|string//|pet status in the store|available, pending, sold|
| |tags|body|array/object/|||


### Responses

200 OK: Successful operation

Content type: application/xml

body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |category|body|object//|||
| |id|body|integer//int64|||
| |name|body|string//|||
| |photoUrls|body|array/string/|||
| |status|body|string//|pet status in the store|available, pending, sold|
| |tags|body|array/object/|||


---

**POST** /pet

Add a new pet to the store

### Request parameters

Content type: application/json

Body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |photoUrls|body|array/string/|||
| |status|body|string//|pet status in the store|available, pending, sold|
| |tags|body|array/object/|||
| |category|body|object//|||
| |id|body|integer//int64|||
| |name|body|string//|||


### Responses

200 OK: Successful operation

Content type: application/xml

body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |category|body|object//|||
| |id|body|integer//int64|||
| |name|body|string//|||
| |photoUrls|body|array/string/|||
| |status|body|string//|pet status in the store|available, pending, sold|
| |tags|body|array/object/|||


---

**PUT** /pet

Update an existing pet by Id

### Request parameters

Content type: application/json

Body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |status|body|string//|pet status in the store|available, pending, sold|
| |tags|body|array/object/|||
| |category|body|object//|||
| |id|body|integer//int64|||
| |name|body|string//|||
| |photoUrls|body|array/string/|||


### Responses

200 OK: Successful operation

Content type: application/json

body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |id|body|integer//int64|||
| |name|body|string//|||
| |photoUrls|body|array/string/|||
| |status|body|string//|pet status in the store|available, pending, sold|
| |tags|body|array/object/|||
| |category|body|object//|||


---

**PUT** /pet

Update an existing pet by Id

### Request parameters

Content type: application/x-www-form-urlencoded

Body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |category|body|object//|||
| |id|body|integer//int64|||
| |name|body|string//|||
| |photoUrls|body|array/string/|||
| |status|body|string//|pet status in the store|available, pending, sold|
| |tags|body|array/object/|||


### Responses

200 OK: Successful operation

Content type: application/json

body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |category|body|object//|||
| |id|body|integer//int64|||
| |name|body|string//|||
| |photoUrls|body|array/string/|||
| |status|body|string//|pet status in the store|available, pending, sold|
| |tags|body|array/object/|||


---

**PUT** /pet

Update an existing pet by Id

### Request parameters

Content type: application/xml

Body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |category|body|object//|||
| |id|body|integer//int64|||
| |name|body|string//|||
| |photoUrls|body|array/string/|||
| |status|body|string//|pet status in the store|available, pending, sold|
| |tags|body|array/object/|||


### Responses

200 OK: Successful operation

Content type: application/json

body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |photoUrls|body|array/string/|||
| |status|body|string//|pet status in the store|available, pending, sold|
| |tags|body|array/object/|||
| |category|body|object//|||
| |id|body|integer//int64|||
| |name|body|string//|||


---

**PUT** /pet

Update an existing pet by Id

### Request parameters

Content type: application/json

Body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |name|body|string//|||
| |photoUrls|body|array/string/|||
| |status|body|string//|pet status in the store|available, pending, sold|
| |tags|body|array/object/|||
| |category|body|object//|||
| |id|body|integer//int64|||


### Responses

200 OK: Successful operation

Content type: application/xml

body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |category|body|object//|||
| |id|body|integer//int64|||
| |name|body|string//|||
| |photoUrls|body|array/string/|||
| |status|body|string//|pet status in the store|available, pending, sold|
| |tags|body|array/object/|||


---

**PUT** /pet

Update an existing pet by Id

### Request parameters

Content type: application/x-www-form-urlencoded

Body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |photoUrls|body|array/string/|||
| |status|body|string//|pet status in the store|available, pending, sold|
| |tags|body|array/object/|||
| |category|body|object//|||
| |id|body|integer//int64|||
| |name|body|string//|||


### Responses

200 OK: Successful operation

Content type: application/xml

body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |category|body|object//|||
| |id|body|integer//int64|||
| |name|body|string//|||
| |photoUrls|body|array/string/|||
| |status|body|string//|pet status in the store|available, pending, sold|
| |tags|body|array/object/|||


---

**PUT** /pet

Update an existing pet by Id

### Request parameters

Content type: application/xml

Body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |photoUrls|body|array/string/|||
| |status|body|string//|pet status in the store|available, pending, sold|
| |tags|body|array/object/|||
| |category|body|object//|||
| |id|body|integer//int64|||
| |name|body|string//|||


### Responses

200 OK: Successful operation

Content type: application/xml

body

|Required|Name|In|Type / Format|Description|Enum|
|----|----|----|----|----|----|
| |id|body|integer//int64|||
| |name|body|string//|||
| |photoUrls|body|array/string/|||
| |status|body|string//|pet status in the store|available, pending, sold|
| |tags|body|array/object/|||
| |category|body|object//|||


---

