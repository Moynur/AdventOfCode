Advent of Code Data Tool
-----

### Usage

AoC gives each user a different input file, we can get our own input by including session data when sending the HTTP 
request.

This tool will get your AoC input for the day and return it as an array of strings for each line.

- input[0] = line 1
- input[1] = line 2
- .
- .
- .
- input[i] = input i-1

### Setup 

- Open web inspector tools
- Go to an input page eg. https://adventofcode.com/2020/day/1
- Under Network:Headers: find Cookies:session=YourSessionID
- put this in the keys file should look like export = "YourSessionID"
- type source keys to write to env
- import package and call by using AoC.GetInput(dayNeeded, yearNeeded) to get the task needed for the day


