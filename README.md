# Medical App

Please develop a RESTful API that conforms to the following specification.
 
The system is responsible for scheduling medical appointments over the internet. In order to do so it exposes a RESTful API that allows you to get a list of available slots for a given day, request a slot to be reserved and cancel a reserved slot.  All calls are made with a valid token id given previously to the client by other part of the system not relevant to this piece.

## Language

I chose Go as the programming language to support this RESTful service.

Go has solid network support and the tooling around it makes it the best fit.

## Installation

```
docker-compose up -d
```

* Wait for Docker to download the Go image
* Wait for Docker-Compose to setup the containers
* Wait for `entrypoint.sh` to install the dependencies
* Wait 1-2 mins for the RESTful to be fully available
* Open `medicalapp.postman_collection.json` using [Postman](https://www.getpostman.com)
* ???
* Profit

## Data Storage

Data is stored in memory, data is lost during restarts.

An quick improvement would be to write the data into a plain text file during the graceful shutdown of the web server. However, persistency is out of the scope of this project and, as discussed with the interviewer, the current implementation already showcases the expected features.

## Unit-Tests

```sh
docker exec medicalapp go test -v
```

**Note:** Wait 1-2 mins for the RESTful service to be fully available before running the integration tests.

## API Definition

Returns a list of free slots for a given date

```
GET /appointments/<token-id>/<date>/free
{"slots": ["10:00", "10:30", …]}
```

Requests a slot to be reserved on the date and time for the patient name.

```
POST /appointments/<token-id>/<date>/<time>/<patient name>
{"appointmentId": xxx}

{"error": "Unable to reserve the appointment"}
```
 
Deletes an appointment

```
DELETE /appointments/<token-id>/<appointment-id>
{"success"}

{"error": "Unable to cancel the appointment"}
```

For the previously defined system, write the automated tests to validate the system.
