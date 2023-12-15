package charge

import "fmt"

type Mobile interface {
	ChargeAppleMobile()
}

type Apple struct{}

func (a *Apple) ChargeAppleMobile() {
	fmt.Println("charge APPLE mobile")
}

type Client struct{}

func (c *Client) ChargeMobile(mob Mobile) {
	fmt.Println("connecting charging adapter to socket")

	fmt.Println("charging adapter connected")

	mob.ChargeAppleMobile()
}
