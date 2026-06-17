# TODO-CLI-GO

A command-line todo application written in Go.

This project was originally built by following a Go CLI tutorial and was later extended with:

- Multi item deleteion
- Improved timeformat

## Installation

git clone git@github.com:Teddinator/TODO-CLI-GO.git
cd TODO-CLI-GO
go build -o todo

## Usage

./todo -add "Buy milk"
./todo - list 
./todo -toggle 1
./todo -del 1 or 1,2,3