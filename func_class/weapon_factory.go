package func_class

type WeaponFactory struct {
	Cutter
	Gun
}

func (factory WeaponFactory) CreateFactory() {
	factory.Cutter = Cutter{}
	factory.Gun = Gun{}
}

func (factory WeaponFactory) Hit(weapon string) {
	if weapon == "gun" {
		factory.Gun.shoot(weapon)
	} else if weapon == "cutter" {
		factory.Cutter.shoot(weapon)
	}
}
