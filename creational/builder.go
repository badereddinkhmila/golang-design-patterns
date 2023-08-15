package creational

import "errors"

/*********************************************************************/
/*************************** Builder Interface ***********************/
/*********************************************************************/
type IHouseBuilder interface {
	setWindowsType()
	setNumberWindows()
	setDoorsType()
	setNumberDoors()
	setNumFloors()
	getHouse() House
}

type BuilderType string

const (
	Normal BuilderType = "normal"
	IGloo  BuilderType = "igloo"
)

func getHouseBuilder(builderType BuilderType) (IHouseBuilder, error) {
	if builderType == Normal {
		return nil, nil
	}
	if builderType == IGloo {
		return nil, nil
	}
	return nil, errors.New("builder type not valid")
}

/*********************************************************************/
/*************************** Builder Implementation ******************/
/*********************************************************************/

type IglooBuilder struct {
	windowsType   string
	numberWindows int
	doorsType     string
	numberDoors   int
	floor         int
}

func (b *IglooBuilder) setWindowsType() {
	b.windowsType = "snow windows"
}
func (b *IglooBuilder) setNumberWindows() {
	b.numberWindows = 15
}

func (b *IglooBuilder) setDoorsType() {
	b.doorsType = "snow doors"
}

func (b *IglooBuilder) setNumberDoors() {
	b.numberDoors = 10
}
func (b *IglooBuilder) setNumFloors() {
	b.floor = 2
}

func (b *IglooBuilder) getHouse() House {
	return House{
		doorsType:     b.doorsType,
		numberDoors:   b.numberDoors,
		windowsType:   b.windowsType,
		numberWindows: b.numberWindows,
		floors:        b.floor,
	}
}

type NormalBuilder struct {
	windowsType   string
	numberWindows int
	doorsType     string
	numberDoors   int
	floor         int
}

func (n *NormalBuilder) setWindowsType() {
	n.windowsType = "snow windows"
}
func (n *NormalBuilder) setNumberWindows() {
	n.numberWindows = 15
}

func (n *NormalBuilder) setDoorsType() {
	n.doorsType = "snow doors"
}

func (n *NormalBuilder) setNumberDoors() {
	n.numberDoors = 10
}
func (n *NormalBuilder) setNumFloors() {
	n.floor = 2
}

func (n *NormalBuilder) getHouse() House {
	return House{
		doorsType:     n.doorsType,
		numberDoors:   n.numberDoors,
		windowsType:   n.windowsType,
		numberWindows: n.numberWindows,
		floors:        n.floor,
	}
}

/*********************************************************************/
/*************************** Director ********************************/
/*********************************************************************/

type Director struct {
	builder IHouseBuilder
}

func newDirector(b IHouseBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b IHouseBuilder) {
	d.builder = b
}

func (d *Director) buildHouse() House {
	d.builder.setDoorsType()
	d.builder.setNumberDoors()
	d.builder.setWindowsType()
	d.builder.setNumberWindows()
	d.builder.setNumFloors()
	return d.builder.getHouse()
}

/*********************************************************************/
/*************************** Product *********************************/
/*********************************************************************/
type House struct {
	numberWindows int
	windowsType   string
	numberDoors   int
	doorsType     string
	floors        int
}
