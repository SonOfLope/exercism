package elon

import "fmt"

// TODO: define the 'Drive()' method
func (car *Car) Drive() {
	if !(car.battery < car.batteryDrain) {
		car.distance += car.speed
		car.battery -= car.batteryDrain
	}
}

// 'DisplayDistance() string' method
func (car *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

// 'DisplayBattery() string' method
func (car *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

// 'CanFinish(trackDistance int) bool' method
func (car *Car) CanFinish(trackDistance int) bool {
	return (trackDistance-car.distance)/car.speed <= car.battery/car.batteryDrain
}
