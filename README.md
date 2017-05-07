Report Portal Client
=======

This is a Go wrapper for working with Report Portal WEB API
[Web API](https://rp.epam.com/ui/swagger-ui.html).

By using this library you agree to Report Portal
[Developer Terms of Use](http://reportportal.io/).

## Installation

To install the library, simply

`go get github.com/darkowl91/rp-client`

## Authentication

Web API functionality is available with authenticating.
To get your access token you should be a Report Portal user, if so take token
from user profile page

## API Examples

Examples of the API can be found in the [examples](examples) directory.

JUnit XML Reporter for Report portal
=======
[JUnit xml format](http://stackoverflow.com/questions/4922867/junit-xml-format-specification-that-hudson-supports)

## Installation

Download the [latest binary release](https://github.com/darkowl91/rp-client/releases) and unpack it.

## Usage
```bash
rp-client [OPTIONS] (DIR|FILE)

Options:
        -r      --rp            Report Portal host
	-d	--debug		Report Portal debug mode
	-p	--project	Report Portal project
	-l	--launch	Report Portal launch name
	-t	--tags		Report Portal launch tags
	-id	--uuid		Report Portal user id
	-h,	--help		Print usage
	-v,	--version	Print version information and quit
Example:
    rp-client -r http://example.com/api/v1/ -p PROJECT -l LAUNCH -t tag1,tag2,tag3 -id your_id ./examples/report
```

