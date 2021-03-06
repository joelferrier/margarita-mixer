# Margarita Mixer

Margarita Mixer is a LiME module build tool which compiles LiME kernel modules and volatility/rekall profiles for the following distributions.

- Amazon Linux 1
- Amazon Linux 2
- Centos 6
- Centos 7
- Debian 7
- Debian 8
- Ubuntu 12.04
- Ubuntu 14.04
- Ubuntu 15.10
- Ubuntu 16.04
- Ubuntu 16.10

# Developing

This project is developed with Go version 1.10.  Previous go releases may function but are not tested.

[dep](https://github.com/golang/dep) is used to manage dependencies.  Ensure it is installed and run `dep ensure` prior to beginning development.

Docker api version v1.35 is currently targeted (ships with 17.12.1-ce).
This will be revised with compatability in mind prior to an initial release.

# License

The MIT License (MIT)

Copyright (c) 2018 Joel Ferrier

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
