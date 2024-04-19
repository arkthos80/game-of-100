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

Example of a solved grid, found in 253 msec: 


|||||||||||
|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
| 48 |  5 | 34 | 39 |  6 | 33 | 40 |  7 | 32 | 91 |
| 20 | 83 | 60 | 19 | 84 | 99 | 14 | 85 | **100** | 13 |
| 74 | 38 | 49 | 73 | 35 | 50 | 72 | 92 | 51 | 71 |
| 47 |  4 | 25 | 98 | 59 | 26 | 41 |  8 | 31 | 90 |
| 21 | 82 | 61 | 18 | 81 | 64 | 15 | 86 | 65 | 12 |
| 75 | 37 | 56 | 78 | 36 | 55 | 67 | 93 | 52 | 70 |
| 46 |  3 | 24 | 97 | 58 | 27 | 42 |  9 | 30 | 89 |
| 22 | 79 | 62 | 17 | 80 | 63 | 16 | 87 | 66 | 11 |
| 76 | 96 | 57 | 77 | 95 | 54 | 68 | 94 | 53 | 69 |
| 45 |  2 | 23 | 44 |  **1** | 28 | 43 | 10 | 29 | 88 |
|||||||||||

### Moves executed
 [Left Up Up Up Right Right Down Down Down UpRight Up Up Left Down Down Left Up Up Left Down Down DownRight Up Up Right Down Down Right Up Up Up Left Left DownRight Down Left Up UpRight Right Down Down Down Left Left Up Up Up DownRight Right Right Down Down Left Up Left Down UpRight Up UpLeft Down Down Right Up Right Down UpLeft Down Right Up Up Left Left Left Down Down Right Up DownLeft Right Up Left Up Right Right Down Down DownRight Up Up Up DownLeft Down Down Left Left UpRight Up UpRight Right]

## Strategies
The solver implements the following strategies to navigate through the grid:
- **Random**: Places numbers randomly while respecting the game's rules.
- **CrossFirst**: Prioritizes horizontal and vertical moves before considering diagonal ones.
- **DiagonalFirst**: Prefers diagonal moves before exploring horizontal and vertical options.

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