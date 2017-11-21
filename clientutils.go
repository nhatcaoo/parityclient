package parityclient

var client *Client

// GetClient connect to the local parity node - assume localhost:8545
func GetClient(endPoint string) (client *Client, err error) {
	if client != nil {
		return client, nil
	}
	client, err = Dial(endPoint)
	return
}
