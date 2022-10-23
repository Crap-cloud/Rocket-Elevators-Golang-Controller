# Rocket-Elevators-Golang-Controller

# Usage

To use the program, you need to donwload the latest golang version. After that you need to type :

`go mod init Rocket-Elevators-Commercial-Controller`

The code to run the scenarios is included, and can be executed with:

`go run . <SCENARIO-NUMBER>`

### Running the tests

To launch the tests:

`go test`

# Description

Explain the purpose of your code/what it does.

Example:

This program controls a Column of elevators.

It sends an elevator when a user presses a button on a floor and it takes
a user to its desired floor when a button is pressed from the inside the elevator goes to the requested floor.

Elevator selection is based on score assigned from 1 to 5. The first elevator we check will always be the best at the beginnig (like a reference). The bestScore we can get is 1. If elevators have the same score, we do some calculations on the gap between elevators user position.

# Video Link

Here is the link to watch my code explanation video on youtube : https://www.youtube.com/watch?v=h1v8tgQJK5o