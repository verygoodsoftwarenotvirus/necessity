package necessity

import "encoding/xml"

func MustUnmarshalXMLBytes[T any](s []byte) T {
	var x T

	if err := xml.Unmarshal(s, &x); err != nil {
		panic(err)
	}

	return x
}

func MustUnmarshalXMLString[T any](s string) T {
	return MustUnmarshalXMLBytes[T]([]byte(s))
}
