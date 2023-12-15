package adapter

import (
	"design-pattern/1adapter/charge"
	"fmt"
)

type android struct {
	brand string
}

func (a *android) chargeAndroid() {
	fmt.Printf("Charge %s mobile\n", a.brand)
}

type androidAdapter struct {
	android *android
}

func (a *androidAdapter) ChargeAppleMobile() {
	a.android.chargeAndroid()
}

func Run() {
	client := &charge.Client{}
	apple := &charge.Apple{}
	client.ChargeMobile(apple)

	android := &android{
		brand: "huawei",
	}
	androidAdapter := &androidAdapter{
		android: android,
	}
	client.ChargeMobile(androidAdapter)
}
