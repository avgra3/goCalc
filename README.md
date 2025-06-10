# goCalc

![goCalc example](images/image.png)

Calculates the period-over-period (PoP) and year-over-year (YoY) change in counts.

## Prequisites

- Go version 1.23.0 (although newer versions will likely work)

### Optional

- Make -- allows the use of the Makefile to make working with easier

## Installation

### Simple

From your terminal run the following:

```bash
go install https://github.com/avgra3/goCalc@latest
```

Once the installation has completed, you should be able to use `gocalc` from anywhere in your shell.

### From Source

If you would like to install this from source, that is not harder, just more steps. Please use the followingas a guide. **_This assumes you have Make installed on your machine._**

```bash
git clone https://github.com/avgra3/goCalc
cd goCalc
make install
```

Once the installation has completed, you should be able to use `gocalc` from anywhere in your shell.

## Why

This was something I needed to use often enough that a simple tool was needed.
