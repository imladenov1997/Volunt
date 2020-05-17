package model

func (exchangePair *ExchangePair) UpdateFromValue(value *float64) {
	exchangePair.FromValue = value
}

func (exchangePair *ExchangePair) UpdateToValue(value *float64) {
	exchangePair.ToValue = value
}

