# syr-sugoku
Term project for CSE682

## Introduction

Provides the backing RESTful service for a class project focused on software engineering practices.
The project is authored in Go, with dependence on `opencv` for Sudoku puzzle image parsing.

[Link to SRS Document](https://docs.google.com/document/d/1i3jTwvAZrSgjs6TnRywGZNwpDQP4tRTtefdVMilIm94/edit?usp=sharing)

[Link to deployed development instance](https://sudoku-dev.herokuapp.com)

### Partners
- Carl Poole, Syracuse University
- Kevin Wren, Syracuse University

## Deployment

Deployments are facilitated via Heroku.  Deploys are done via the included `Makefile`. The
app can also be run locally via the Makefile, on port 8080.

To run locally, you must populate an `.env` file with a populated and valid `DATABASE_URL` variable.

## Dependencies

This web app utilzes a library written by Kevin W that is focused on solving the puzzles that this app
is charged with parsing.

This web app also utilizes parsing code re-used under the MIT license.  Said code is located in `pkg/sudokuparser`
and was cloned locally to work around `go get` issues.  Credit to James Andersen.  His original repo is located 
[here](https://github.com/jamesandersen/go-sudoku).

