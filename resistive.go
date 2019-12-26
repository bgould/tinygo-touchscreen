package touchscreen

import (
	"machine"
)

type Resistive4Wire struct {
	YP machine.Pin
	YM machine.Pin
	XP machine.Pin
	XM machine.Pin

	RxPlate int

	yp machine.ADC
	ym machine.ADC
	xp machine.ADC

	samples []uint16
}

func (res *Resistive4Wire) Configure() {
	res.yp = machine.ADC{res.YP}
	res.ym = machine.ADC{res.YM}
	res.xp = machine.ADC{res.XP}
	res.samples = make([]uint16, 2)
}

func (res *Resistive4Wire) GetTouchPoint() (p Point) {
	p.X = int(res.ReadX())
	p.Y = int(res.ReadY())
	p.Z = int(res.ReadZ())
	return
}

func (res *Resistive4Wire) ReadX() uint16 {
	//res.YP.Configure(input)
	//res.YP.Low()
	res.YM.Configure(input)
	res.YM.Low()

	res.XP.Configure(output)
	res.XM.Configure(output)
	res.XP.High()
	res.XM.Low()

	res.yp.Configure()
	delayMicroseconds(20)

	//return res.yp.Get()
	res.samples[0] = res.yp.Get()
	res.samples[1] = res.yp.Get()
	return 1023 - (((res.samples[0] + res.samples[1]) / 2) >> 4)
}

func (res *Resistive4Wire) ReadY() uint16 {
	res.XM.Configure(input)
	res.XM.Low()

	res.YP.Configure(output)
	res.YM.Configure(output)
	res.YP.High()
	res.YM.Low()

	res.xp.Configure()
	delayMicroseconds(20)

	//return res.xp.Get()
	res.samples[0] = res.xp.Get()
	res.samples[1] = res.xp.Get()
	return 1023 - (((res.samples[0] + res.samples[1]) / 2) >> 4)
}

func (res *Resistive4Wire) ReadZ() uint16 {
	// set x- to ground
	res.XM.Configure(output)
	res.XM.Low()

	// set y+ to VCC
	res.YP.Configure(output)
	res.YP.High()

	// Hi-Z x+ and y-
	res.xp.Configure()
	res.ym.Configure()

	z1 := res.xp.Get()
	z2 := res.yp.Get()

	//if (res.rxplate != 0) {

	//} else {
	return (1023 - (z2>>4 - z1>>4))
	//}

}
