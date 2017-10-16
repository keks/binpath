/*
   This file is part of binpath.

   binpath is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   binpath is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with binpath.  If not, see <http://www.gnu.org/licenses/>.
*/

// binpath is a package providing tools for binary paths
package binpath

import (
	"encoding/base64"
	"strings"
)

// Path is a path that may contain binary data
type Path []byte

// FromBytes converts a byte slice to a path. Arguments should only contain a single path element.
func FromBytes(bs []byte) Path {
	p := make(Path, 1, len(bs)+1)

	p[0] = byte(len(bs))
	p = append(p, bs...)

	return p
}

// FromString converts a string to a path. Arguments should have the format "/home/keks/go", with the leading slash being optional.
// if an element in the path (i.e. "home" or "keks" in the example above) starts with "b:", the rest of the element is parsed as URL-compatible base64.
func FromString(sp string) (Path, error) {
	if len(sp) > 0 && sp[0] != '/' {
		sp = "/" + sp
	}

	parts := strings.Split(sp, "/")

	var (
		p   = make(Path, len(sp))
		cur = p
	)

	for _, el := range parts {
		var data []byte

		if strings.HasPrefix(el, "b:") {
			var err error

			b64 := el[2:]
			data, err = base64.URLEncoding.DecodeString(b64)
			if err != nil {
				return nil, err
			}
		} else {
			data = []byte(el)
		}

		l := len(data)
		if l == 0 {
			continue
		}

		cur[0] = byte(l)
		copy(cur[1:], data[:])
		cur = cur[l+1:]
	}

	return p[:len(p)-len(cur)], nil
}

func (p Path) isBinary() bool {
	p = p[1:]
	for _, c := range p {
		if c < ' ' || c > 127 {
			return true
		}
	}
	
	return false
}

// String returns the path's string representation
func (p Path) String() string {
	var (
		head Path
		out string
	)
	
	for len(p) > 0 {
		head, p = p.Pop()
		if len(head) == 0 {
			continue
		}
		out += "/"
		
		if head.isBinary() {
			out += "b:" + base64.URLEncoding.EncodeToString(head[1:])
		} else {
			out +=string(head[1:])
		}
	}
	
	return out
}
		

// Pop splits the Path in two parts: the first element ("head") and the rest ("tail")
func (p Path) Pop() (head Path, tail Path) {
	if len(p) == 0 {
		return Path{}, Path{}
	}
	
	l := p[0]
	
	return p[:l+1], p[l+1:]
}

// Join joins two paths and returns the result.
func Join(ps ...Path) Path {
	var l int

	for _, p := range ps {
		l += len(p)
	}

	out := make(Path, 0, l)
	for len(ps) > 0 {
		out = append(out, ps[0]...)
		ps = ps[1:]
	}

	return out
}

func Must(p Path, err error) Path {
	if err != nil {
		panic(err)
	}

	return p
}
