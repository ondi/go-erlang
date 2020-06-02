//
//
//

package erlang

import "errors"

// e - вероятность отказа
// y - нагрузка
// v - число линий
func E1(y float64, v float64) (e float64) {
	e = 1
	for i := float64(1); i <= v; i++ {
		e = y * e / (i + y*e)
	}
	return
}

// e - вероятность ожидания
// y - нагрузка
// v - число линий
func E2(y float64, v float64) (e float64) {
	e = E1(y, v)
	return v * e / (v - y*(1-e))
}

func V1(y float64, p float64) (v float64, err error) {
	if y < 0 || p < 0 || p > 1 {
		return 0, errors.New("invalid argument")
	}
	for e := float64(1); e > p; {
		v++
		e = y * e / (v + y*e)
	}
	return
}

func V2(y float64, p float64) (v float64, err error) {
	if y < 0 || p < 0 || p > 1 {
		return 0, errors.New("invalid argument")
	}
	v = 1
	for e := float64(1); (v * e / (v - y*(1-e))) > p; {
		v++
		e = (y * e / v) / (1 + y*e/v)
	}
	return
}

func Y1(p float64, v float64, precision float64) (y float64, err error) {
	if p < 0 || p > 1 {
		return 0, errors.New("invalid argument")
	}
	// find interval for bipartitioning
	y1 := float64(0)
	y2 := v
	for E1(y2, v) < p {
		y1 = y2
		y2 += y2 / 2
	}
	for y2-y1 > precision {
		y = (y1 + y2) / 2
		if E1(y, v) < p {
			y1 = y
		} else {
			y2 = y
		}
	}
	return
}

func Y2(p float64, v float64, precision float64) (y float64, err error) {
	if p < 0 || p > 1 {
		return 0, errors.New("invalid argument")
	}
	// find interval for bipartitioning
	y1 := float64(0)
	y2 := v
	for E2(y2, v) < p {
		y1 = y2
		y2 += y2 / 2
	}
	for y2-y1 > precision {
		y = (y1 + y2) / 2
		if E2(y, v) < p {
			y1 = y
		} else {
			y2 = y
		}
	}
	return
}

// v_utilization = y / v
// y_served = y - y * p
// y_lost = y * p
// y = number_of_calls * duration_in_hours
