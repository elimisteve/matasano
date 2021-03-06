package padding

import (
    "testing"
)

func TestPad(t * testing.T) {
    // data less than block size
    out, err := Pad(3, []byte{'a'})
    if err != nil {
        t.Fatal()
    }
    exp := []byte{'a', 2, 2}
    for i, v := range exp {
        if v != out[i] {
            t.Fatal()
        }
    }

    // full block
    data := make([]byte, 50)
    out, err = Pad(5, data)
    if err != nil {
        t.Fatal()
    }
    for i := 50; i < 55; i++ {
        if out[i] != uint8(5) {
            t.Fatal()
        }
    }

    // pad end of long data
    data = make([]byte, 50)
    out, err = Pad(7, data)
    if err != nil {
        t.Fatal()
    }
    for i := 50; i < 52; i++ {
        if out[i] != uint8(6) {
            t.Fatal()
        }
    }

    // matasano test
    in := []byte("YELLOW SUBMARINE")
    exp = []byte("YELLOW SUBMARINE\x04\x04\x04\x04")
    out, err = Pad(20, in)
    for i, v:= range out {
        if v != exp[i] {
            t.Fatal()
        }
    }
}
func TestUnPad(t * testing.T) {
    good := []byte("ICE ICE BABY\x04\x04\x04\x04")
    bad1 := []byte("ICE ICE BABY\x05\x05\x05\x05")
    bad2 := []byte("ICE ICE BABY\x01\x02\x03\x04")
    bs := 16

    _, err := UnPad(bs, good)
    if err != nil {
        t.Fatal()
    }

    _, err = UnPad(bs, bad1)
    if err != InvalidPad {
        t.Fatal()
    }

    _, err = UnPad(bs, bad2)
    if err != InvalidPad {
        t.Fatal()
    }
}
