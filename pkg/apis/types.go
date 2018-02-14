package apis

type RequestDetails struct {
  Method string
  Headers map[string]string
  Path string
  Query map[string]string
}
