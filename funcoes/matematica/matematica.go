package matematica

func Aquivo() string {
	return "pacote matematica"
}

func Delta(a, b, c float64) float64 {
	var d float64
	d = (b * b) - (4 * a * c)
	return d
}

func MySqrt(num float64) float64 {
	var raiz, ant, y float64
	raiz = num / 2
	for i := 1.0000; i > 0.0001; i = y {
		ant = raiz
		raiz = (raiz*raiz + num) / (2 * raiz)
		y = raiz - ant
		if y < 0 {
			y *= -1
		}
	}
	return raiz
}

func Raizes(a, b, c float64) (float64, float64) {
	var de, x1, x2 float64
	de = Delta(a, b, c)
	de = MySqrt(de)

	if de < 0 {
		return -1.0, -1.0
	} else {
		b *= -1
		x1 = (b - de) / (2 * a)
		x2 = (b + de) / (2 * a)
		if x1 < 0 {
			return 0, x2
		} else if x2 < 0 {
			return x1, 0
		}
	}
	return x1, x2

}
