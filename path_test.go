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

package binpath

import (
	"bytes"
	"testing"
)

func TestPath(t *testing.T) {
	p, err := FromString("/home/keks/go/src/cryptoscope.co/binpath")
	if err != nil {
		t.Fatal(err)
	}
	exp := []byte{
		4, 'h', 'o', 'm', 'e',
		4, 'k', 'e', 'k', 's',
		2, 'g', 'o',
		3, 's','r','c',
		14, 'c','r','y','p','t','o','s','c','o','p','e','.','c','o',
		7,'b','i','n','p','a','t','h',
	}
	if !bytes.Equal([]byte(p), exp) {
		t.Fatalf("expected \n%x but got \n%x", exp, []byte(p))
	}
}

func TestJoin(t *testing.T) {
	p1, err := FromString("/home")
	if err != nil {
		t.Fatal(err)
	}
	p2, err := FromString("keks")
	if err != nil {
		t.Fatal(err)
	}
	p := Join(p1, p2)
	if p.String() != "/home/keks" {
		t.Fatalf("got %#v", p)
	}
}


func TestBinary(t *testing.T) {
	p1, err := FromString("bin")
	if err != nil {
		t.Fatal(err)
	}
	p2 := Path{4, 60, 90, 129, 37}
	
	p := Join(p1, p2)
	t.Log(p)
	if p.String() != "/bin/b:PFqBJQ==" {
		t.Fatalf("got %v", p)
	}
	
	pOld := p
	
	p, err = FromString(p.String())
	if err != nil {
		t.Fatal(err)
	}
	
	if !bytes.Equal(p, pOld) {
		t.Fatalf("%x != %x", []byte(p), []byte(pOld))
	}
}