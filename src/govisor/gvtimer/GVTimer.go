package gvtimer

import (

)
//MAX_TIMEOUT long = 5000

type GVTimer struct{
	max_timeout int64
	min_timeout int64
}
func NewGVTimer() *GVTimer {
	return &GVTimer{max_timeout:5000, min_timeout:1}
}	

