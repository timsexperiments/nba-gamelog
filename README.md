# NBA Game Log Scraper CLI

## Introduction

This CLI tool, named `gamelog`, scrapes NBA game logs from [Basketball Reference](https://www.basketball-reference.com/) for specified seasons and outputs the data to a CSV file. You can specify single seasons or a range of seasons and choose the output file location.

## Installation

To install the `gamelog` CLI, use the following Go command:

```sh
go install gamelog
```

Make sure your Go environment is properly set up and `$GOPATH/bin` is in your system's PATH.

## Usage

- **Single Season**: `gamelog --season 2023` (Required if no start season is provided)
- **Season Range**: `gamelog --start 2020 --end 2023` (If no end season is provided, defaults to the current year)
- **Output Location**: gamelog --season 2023 --output /path/to/output.csv (Defaults to nba_gamelog.csv in the user's home directory if not specified)

## Input Parameters

- ` --season`: Specifies a single NBA season (either in YY or YYYY format). This is required if no start season (`--start`) is provided.
- `--start`: The starting season of the range to scrape. Required if `--season` is not provided.
- `--end`: The ending season of the range. If not specified, it defaults to the current year.
- `--output`: Path for the output CSV file. Defaults to `nba_gamelog.csv` in the user's home directory.

## Output

The tool generates a CSV file with NBA game logs. By default, the file is saved in the user's home directory as `nba_gamelog.csv`. You can specify a custom path with the `--output` flag.

## Data Source

This project uses data from [Basketball Reference](https://www.basketball-reference.com/). Please cite them and provide a link to their website when using this data.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
