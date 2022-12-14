# Parsyl Backend Code Challenge

## Pre-requisites
- An IDE or text editor of your choice
- Docker

## Your Task

Parsyl ingests a lot of data from hardware in order to present that information to users.  Your task will be to create an API that ingests log messages, transforms them for a database and then aggregates that data for display via other endpoints.  Review the "user stories" below for instructions.

Our stack is written in Go and we’d love to see a solution in that but you are welcome to build your solution in whatever language and libraries you are most comfortable with.  Google and the internet are at your disposal.  

You are free to spend as much time as you’d like to solve this challenge, however, *we don’t expect more than 2-3 hours*.

### Considerations
When developing your solution, please consider the following:

- Structure of your endpoints
- Quality of your code
- Testability
- Can your solution be easily configured and deployed (deployment not actually required)?

**It’s okay if you don’t finish the entire challenge.**  We favor high quality code over completing all the stories!  One well built story is better than three poorly built ones.  It's not uncommon for the first story to take significantly longer than the second or third.

**It’s okay if you get stuck.**  Please reach out and we will answer questions or try to help you get moving again.  There is an [FAQ](#frequently-asked-questions) which may contain an answer but don't hesitate to reach out.


## Getting Started

1. Make sure Docker is running
2. Build and start the database Docker image: `make bootstrap`
3. Write and structure your project and code how you see fit
	* Reference [sample_log_messages.txt](./sample_log_messages.txt) for a bunch of sample log messages
	* [sample_requests](#sample-requests) contains a couple cURL requests that can be used to test your work
4. Stop the docker containers when you’re finished: `make stop`


## User Stories To Complete

You can complete the stories any way you see fit but we recommend they be completed in order with distinct commits (or a merge commit) for each.

### Story #1: Ingest log messages

- A [log message](#log-messages) is stored appropriately in the [database](#database)
- Logs are ingested through the `/ingest` endpoint
- Temperature values should be converted to farenheit before saving to the database
	- `TempF = (2*TempC) + 30`
- The request format should have the contract (see [sample_requests](#sample-requests) below):

	```json
	{
		"data": [
			// Log event string
			//"YYARKx|2022-03-22T21:42:02.362Z|-8"
		]
	}
	```

**Implementation Note:** See the detail about the [database and log structure](#data-structures) below.
<br />

### Story #2: Display aggregated data by device id and date

- Endpoint exists at `/device/{id}` to retrieve log messages aggregated by device id (see [sample_request](#sample-requests) below)
- Datetimes can be in UTC
- Log events should be sorted ASC by log date
- Response should have the contract:

	```json
	{
		"deviceId": String,
		"averageTemperature": Number,
		"mostRecentLogDate": Datetime,
		"logs": [
			{
				"logDate": Datetime,
				"temperature": Number,
				"humidity": Number
			}
		]
	}
	```
	<br />

### Story #3: Flag records according to business rules

- Log messages that have a temperature greater than 32 degrees Fahrenheit should be flagged
- Response should have the contract:
    
    ```json
    {
    	"deviceId": String,
    	"averageTemperature": Number,
    	"mostRecentLogDate": Datetime,
    	"totalAlerts": Number, // <-- New field
    	"logs": [
    		{
    			"logDate": Datetime,
    			"temperature": Number,
    			"alert": true // <-- New field
    		}
    	]
    }
    ```

<br /><br />
<hr />

## Data Structures
### Log Messages
The log messages are pipe delimited strings that have the following structure:

```
device_id|datetime|temperature_celsius
```

- Datetime is in UTC
- Temperature is in celsius

Note, this structure cannot be modified!

### Database
The database stores a row per log message.  The `Logs` table has the following schema:

| Column         | Type       | Constraints                 |
| -------------- | ---------- | --------------------------- |
| event_id       | INT        | Primary Key, auto increment |
| event_date     | DATETIME   | Not Null                    |
| device_id      | VARCHAR(6) | Not Null                    |
| temp_farenheit | INT        |                             |

Note, this schema cannot be changed!

Once Postgres is spun up in Docker, you connect to it with any postgres library you choose.  The schema will be in place but there will be no data.  The connection information should be:

- Host: `http://localhost`
- Port: `5432`
- User: `pguser`
- Password: `pgpassword`
- Database: `code_challenge`


<br /><br />
<hr />

## Frequently Asked Questions

### Can I run queries against the database with a client?
You can run queries directly.  If you use a client (e.g. Postico), you can use the connection information provided in the README.  If you prefer to use `psql` you can connect to the running docker container:

1. Connect to the container: `docker exec -it db bash`
2. Connect to the database: `psql --user pguser --db=code_challenge`
3. Submit queries directly (e.g. `SELECT * FROM logs`)

### I get the error: Error response from daemon: Conflict. The container name "/db" is already in use...
This is because you stopped the container without removing it.  If you want to start where you left off previously, you can simply run `docker start db`.  Alternately, if you want to start from a clean install of the database then you can run `docker rm db` and then run `make bootstrap` again.

This issue and solution will work for the generator container as well.


<br /><br />
<hr />

## Sample Requests
Some sample cURL requests are provided to help you interact with the API you build.  They default to `http://localhost:8000`.  You can also copy these commands into a REST client (e.g. Postman) if you prefer to use a GUI instead of the terminal.

### GET /ingest
```
curl --location --request POST 'http://localhost:8000/ingest' \
--header 'Content-Type: application/json' \
--data-raw '{
    "data": [
        "YYARKx|2022-03-22T21:42:02.362Z|-8",
        "YYARKx|2022-03-22T21:42:04.372Z|-1",
        "YYARKx|2022-03-22T21:42:50.572Z|7"
    ]
}'
```

### POST /device/{id}
```
curl --location --request GET 'http://localhost:8000/device/YYARKx'
```