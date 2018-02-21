package contenttypes

type ContentType interface {
	FillBuffer(buffer []byte)
}

func GetContentTypeImplementation(contenttypeName string) ContentType {
	implementation := newPlaintext()
	return implementation
}
