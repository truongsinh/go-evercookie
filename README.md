# Evercookie Middleware for Golang

[![Build Status](https://travis-ci.org/truongsinh/go-evercookie.png?branch=master)](https://travis-ci.org/truongsinh/go-evercookie)

[Evercookie](http://samy.pl/evercookie/) is a Javascript API that produces extremely persistent cookies in a browser.
It is written in JavaScript and additionally uses a SWF (Flash) object for the Local Shared Objects and,
originally, PHPs for the server-side generation of cached PNGs and ETags.

This middleware port original PHP script to Golang

# Go version support
1.0+

# Install
```bash
go get go-evercookie
```

# Usage
`@todo`

# Settings
Customized settings can be used, but up to this moment, it makes no sense to change the default one,
as all these values are hardcoded in (frontend) `evercookie.js`.
`@todo`

# Test
```bash
go test
```

# Acknowledgement
- [Samy Kamkar] (https://github.com/samyk) for his awesome idea and implementation of [Evercookie](http://samy.pl/evercookie/)
- [Golang community] (https://github.com/golang/go) for our awesome [Golang](http://expressjs.com/)

# License
The MIT License (MIT)
Copyright (c) 2015 [TruongSinh Tran-Nguyen](i@truongsinh.pro)

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"),
to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense,
and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
IN THE SOFTWARE.