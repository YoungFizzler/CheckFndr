# CheckFndr

This is a small command line tool that finds the check digit algorithm used for a given barcode.

## Installation

```
go install github.com/YoungFizzler/CheckFndr@latest
```

## Usage

```
Usage: CheckFndr <barcode>
```

## Examples

```
$ CheckFndr 12345678901
The barcode 12345678901 uses the EAN-13 check digit algorithm.


To add...

- Allow multiple barcodes for better confidence scores.
- Allow for more barcode types.
- Allow for PNG/SVG/PDF/etc input.