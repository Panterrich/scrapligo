package channel

import "fmt"

// Write writes the given bytes b to the channel.
func (c *Channel) Write(b []byte, r bool) error {
	lm := string(b)
	if r {
		lm = redacted
	}

	c.l.Debugf("channel write %#v", lm)

	fmt.Printf("Write bytes: \"%s\"\n", string(b))
	fmt.Printf("Write bytes (hex):\n")
	for i, byteVal := range b {
		fmt.Printf("%02x ", byteVal)
		if (i+1)%16 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()

	return c.t.Write(b)
}

// WriteReturn writes the channel ReturnChar to the channel.
func (c *Channel) WriteReturn() error {
	return c.Write(c.ReturnChar, false)
}

// WriteAndReturn writes the given bytes b and then sends the channel ReturnChar.
func (c *Channel) WriteAndReturn(b []byte, r bool) error {
	err := c.Write(b, r)
	if err != nil {
		return err
	}

	return c.WriteReturn()
}
