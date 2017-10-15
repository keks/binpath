# binpath

A library providing tools for binary paths. Instead of delimiting path elements with a slash, they are simply length-prefixed. This way it is safe to store binary data inside.

My use case: I wanted to have paths in a key-value store, so I just use binpaths as keys.

[Godoc]: https://godoc.org/cryptoscope.co/go/binpath