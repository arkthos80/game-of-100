# Game-Of-100 Puzzle Solver

## Overview
This Go project provides a fast and efficient solver for the "100" solitaire number game. It employs three distinct strategies to fill a 10x10 grid with numbers from 1 to 100, adhering to the game's placement rules. The solver is designed to execute multiple runs within seconds, making it highly effective for finding solutions.

## Game Rules
The "100" solitaire number game involves placing numbers from 1 to 100 on a 10x10 grid. The rules are as follows:
- Place the number **1** in any grid position.
- Subsequent numbers must be placed:
  - Horizontally or vertically, skipping two squares from the last number.
  - Diagonally, skipping one square from the last number.
- The game ends when no more numbers can be placed following these rules. The goal is to fill the entire grid.

Rules reference[ITA]: https://www.attivitiamo.it/gioco-solitario-numeri-100/

Example of a solved grid, found in 40 msec: 


|||||||||||
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
|**1**| | 34| |  6| | 21| | 33| |  7| | 22| | 32| |  8| | 23|
| 37| | 17| |  3| | 36| | 16| | 13| | 10| | 82| | 14| | 11|
|  5| | 20| | 43| | 99| | 71| | 88| | 45| | 24| | 69| | 31|
|  2| | 35| | 38| | 18| | 62| | 83| | 15| | 12| |  9| | 81|
| 42| | 98| |  4| | 89| | 44| |**100**| | 70| | 87| | 46| | 25|
| 39| | 19| | 54| | 78| | 72| | 91| | 61| | 80| | 68| | 30|
| 56| | 76| | 41| | 97| | 63| | 84| | 94| | 28| | 65| | 86|
| 53| | 59| | 73| | 90| | 60| | 79| | 67| | 50| | 47| | 26|
| 40| | 96| | 55| | 77| | 95| | 92| | 64| | 85| | 93| | 29|
| 57| | 75| | 52| | 58| | 74| | 51| | 48| | 27| | 66| | 49|
|||||||||||

Executions number and total number of move tries:  211 46907

### Moves executed
[Down UpRight Down UpLeft UpRight Right Right Down UpLeft Right DownLeft UpLeft Right DownLeft UpLeft Left DownRight DownLeft Up UpRight Right Right DownLeft DownRight Down DownLeft Up DownRight Up Up UpLeft Left Left Down UpRight Left DownRight DownLeft Down UpRight UpLeft UpRight DownRight UpRight DownRight Down DownLeft Right UpLeft DownLeft Left UpLeft UpRight Down UpLeft Down Right UpLeft Right UpRight UpLeft Down DownRight UpRight Down UpLeft UpRight Up DownLeft UpLeft Down DownLeft DownRight Left Up DownRight Up DownRight UpRight UpRight UpLeft DownLeft Down DownRight UpRight UpLeft UpLeft DownLeft Down UpRight Down Right UpLeft DownLeft Left UpRight UpLeft UpRight DownRight]

## Strategies
The solver implements the following strategies to navigate through the grid:
- **Random**: Places numbers randomly while respecting the game's rules.
- **CrossFirst**: Prioritizes horizontal and vertical moves before considering diagonal ones.
- **DiagonalFirst**: Prefers diagonal moves before exploring horizontal and vertical options.
- **CloseToBorder**: [BEST PERFORMANCE] Prefers a move that keeps the next one as close to one of the 4 borders as possible

## Usage
Run the solver with the desired strategy and optional parameters for maximum timeout and starting grid position:

```go
// Command line usage
game100 STRATEGY MAX_TIMEOUT_SEC [X] [Y]

// Example
game100 CrossFirst 5 0 0
```

Replace STRATEGY with Random, CrossFirst, or DiagonalFirst. MAX_TIMEOUT_SEC is the maximum time in seconds the solver will run before stopping. [X] [Y] are optional parameters to specify the starting position on the grid.


## Contributing
We welcome contributions to enhance the solver’s performance and add new features. Please fork the repository, make your changes, and submit a pull request.

## License
This project is licensed under the MIT License - see the LICENSE.md file for details.