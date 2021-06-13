package client

type Client interface {
	Translate(original map[string]string, sl, tl string) (map[string]string, error)
}
