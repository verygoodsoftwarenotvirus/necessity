package necessity

import "encoding/json"

func MustUnmarshalJSONBytes[T any](s []byte) T {
	var x T

	if err := json.Unmarshal(s, &x); err != nil {
		panic(err)
	}

	return x
}

func MustUnmarshalJSONString[T any](s string) T {
	return MustUnmarshalJSONBytes[T]([]byte(s))
}
