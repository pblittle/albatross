# Albatross - Golf Shot Data Processor

Albatross is a specialized tool designed for processing and analyzing golf shot data. It works with data from launch monitors, standardizing the output for consistent analysis.

## Features

- Parse shot data from CSV files
- Support for MLM2Pro launch monitor
- Normalize club types (e.g., "3 wood" -> "3W")
- Determine shot types (Tee or Approach)
- Calculate median target distances for each club type
- Export processed data to CSV

## Packages

### Models (`internal/models`)

Defines core data structures and interfaces:

- `LaunchMonitor` interface for different launch monitor types
- `RawShotData` struct for storing raw shot data
- `ProcessedShotData` struct for storing processed shot information

### Reader (`internal/reader`)

Provides implementations for different types of launch monitors:

- `MLM2ProLaunchMonitor` for processing data from MLM2Pro launch monitors

### Parsers (`internal/parsers`)

Contains functions for reading and parsing shot data:

- `ProcessShotData` function for reading and processing CSV files

### Processors (`internal/processors`)

Provides functionality for processing golf-related data:

- `NormalizeClubType` function for standardizing club type notation
- `DetermineShotType` function for categorizing shots as Tee or Approach

### Calculators (`internal/calculators`)

Performs statistical analysis of golf shot data:

- `CalculateTargets` function for computing median target distances for each club type

### Utils (`utils`)

Offers utility functions for common tasks:

- `ReplaceFileExtension` function for manipulating file names
- `WriteCSV` function for exporting processed data to CSV files

### Writer (`internal/writer`)

Handles writing processed data to output files:

- `ShotPatternWriter` for writing processed shot data to CSV

## Installation

To install this project, make sure you have Go installed on your system, then clone the repository:

```shell
git clone https://github.com/pblittle/albatross.git
cd albatross
```

## Usage

Albatross is a command-line application. Here's how to use it:

```shell
go run main.go <launch_monitor_type> <input_csv_file>
```

For example:

```shell
go run main.go mlm2pro input_data.csv
```

This will process the `input_data.csv` file using the MLM2Pro launch monitor type and output a file named `input_data_processed.csv` in the same directory.

## Testing

To run the tests, use the following command:

```shell
make test
```

This will run all tests in the project, including unit tests and end-to-end tests.

## Contributing

1. Fork it
2. Checkout the `main` branch (`git checkout -b develop`)
3. Create your feature branch (`git checkout -b my-new-feature`)
4. Sign and commit your changes (`git commit -sam 'Add some feature'`)
5. Push to the branch (`git push origin my-new-feature`)
6. Create new Pull Request

## License

This project is distributed under the [MIT License](LICENSE.md).
