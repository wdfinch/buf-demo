## buf common commands

`buf build` - verfy the proto builds and there are no syntax errors (https://docs.buf.build/build/usage) \
`buf lint` - verfy the proto files are linted correctly (https://docs.buf.build/lint/overview) \
`buf format -w` - format the proto files (https://docs.buf.build/format/usage) \
`buf breaking --against '../../.git#branch=main'` - verify there are no breaking changes against a git branch (https://docs.buf.build/format/usage)