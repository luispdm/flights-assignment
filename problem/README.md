# Take-Home Programming Assignment

## Story
There are over 100,000 flights a day, with millions of people and cargo being transferred around the world. With so many people and different carrier/agency groups, it can be hard to track where a person might be. In order to determine the flight path of a person, we must sort through all of their flight records.

## Goal
To create a simple microservice API that can help us understand and track how a particular person's flight path may be queried. The API should accept a request that includes a list of flights, which are defined by a source and destination airport code. These flights may not be listed in order and will need to be sorted to find the total flight paths starting and ending airports.

The solution should be delivered within 24 hours from the receipt of the challenge description.

### Required Structure
```
[["SFO", "EWR"]] => ["SFO", "EWR"]

[["ATL", "EWR"], ["SFO", "ATL"]] => ["SFO", "EWR"]

[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]] => ["SFO", "EWR"]
```

## Specifications
- Define and document the format of the API endpoint in the README.

- Use Golang and/or any tools that you think will help you best accomplish the task at hand.
