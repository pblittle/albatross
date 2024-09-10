# Albatross - Golf Shot Data Processor

Albatross is a specialized tool designed for processing and analyzing golf shot data. It works with data from launch monitors, standardizing the output for consistent analysis.

## Features

- Parse shot data from CSV files
- Support for MLM2Pro launch monitor
- Normalize club types (e.g., "3 wood" -> "3W")
- Determine shot types (Tee or Approach)
- Calculate median target distances for each club type
- Export processed data to CSV in Shot Pattern format

## Launch Monitor Support

Currently, Albatross supports the [MLM2Pro](https://rapsodo.com/pages/mlm2pro-golf-simulator) launch monitor. If you need support for additional launch monitors, please open an issue on our GitHub repository.

## Output Format

Albatross currently outputs processed data in the [Shot Pattern](https://shotpattern.app/) format. This is the default and only output format available.

## Installation

To install this project, make sure you have Go installed on your system, then clone the repository:

```shell
git clone https://github.com/pblittle/albatross.git
cd albatross
```

## Usage

Albatross is a command-line application. Here's how to use it:

```shell
go run main.go &lt;launch_monitor_type&gt; &lt;input_csv_file&gt;
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
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Sign and commit your changes (`git commit -sam 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

## License

This project is distributed under the [MIT License](LICENSE.md).
