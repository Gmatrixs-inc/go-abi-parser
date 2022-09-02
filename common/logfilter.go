package common

type Logfilter interface {
	Scan(startI, endI int64)
}
