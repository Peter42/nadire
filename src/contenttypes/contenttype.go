package contenttypes

type ContentType interface {
	FillBuffer(buffer []byte)
}

func GetContentTypeImplementation(contenttypeName string) ContentType {
	var implementation ContentType

	if contenttypeName == "text/plain" {
		implementation = newPlaintext()
	} else if contenttypeName == "application/octet-stream" {
		implementation = newRawdata()
	}

	return implementation
}
