package gnuuid

import (
	"bufio"
	"crypto/sha1"
	"io"
	"os"

	u "github.com/google/uuid"
)

// Domain is UUIDv5 seed for `globalnames.org` domain identifiers. It was
// generated originally with
//
// u.NewSHA1(u.NameSpaceDNS, []byte("globalnames.org"))
var (
	GNDomain = u.Must(u.Parse("90181196-fecf-5082-a4c1-411d4f314cda"))
	Nil      = u.Nil
)

// New creates new UUIDv5 identifier for globalnames. It takes a string
// as an argument and returns a UUIDv5 based on the GNDomain and the
// provided string.
func New(name string) u.UUID {
	return u.NewSHA1(GNDomain, []byte(name))
}

// FromFile reads the content of a file and generates UUIDv5 from
// its content. It takes a file path as an argument and returns a
// UUIDv5 based on the GNDomain and the file's content.
func FromFile(path string) (u.UUID, error) {
	file, err := os.Open(path)
	if err != nil {
		return u.Nil, err
	}
	defer file.Close()

	hash := sha1.New()
	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024*1024) // 1MB buffer

	for {
		bytesRead, err := reader.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			return u.Nil, err
		}

		hash.Write(buffer[:bytesRead])
	}

	hashed := hash.Sum(nil)
	return u.NewSHA1(GNDomain, hashed), nil
}
