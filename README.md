# discourse-protobuf

[![Go Reference](https://pkg.go.dev/badge/github.com/michaellindman/discourse-protobuf.svg)](https://pkg.go.dev/github.com/michaellindman/discourse-protobuf)

Protocol Buffers for the Discourse forum API. This is a work in progress largly compiled from the [Discourse API Documentation](https://docs.discourse.org/) and reverse engineering a live instance and should not be concidered production ready, please feel free to submit pull requests.

## Building

### Installing protoc

to build the protocol buffers protoc a protobuf compiler needs to be installed. The latest releases can be found [here](https://github.com/protocolbuffers/protobuf/releases).

#### Linux

Run the following commands:

```sh
curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.19.1/protoc-3.19.1-linux-x86_64.zip
sudo unzip -o protoc-3.19.1-linux-x86_64.zip -d /usr/local bin/protoc
sudo unzip -o protoc-3.19.1-linux-x86_64.zip -d /usr/local 'include/*'
```

and to build the latest proto files run:

```sh
make proto
```

## License

```text
MIT License Copyright (c) 2021 Michael Lindman

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is furnished
to do so, subject to the following conditions:

The above copyright notice and this permission notice (including the next
paragraph) shall be included in all copies or substantial portions of the
Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS
OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF
OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
```
