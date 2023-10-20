package btrip

import (
	"sync"
)

// BaggageSubContentVisualizesBean 结构体
type BaggageSubContentVisualizesBean struct {
}

var poolBaggageSubContentVisualizesBean = sync.Pool{
	New: func() any {
		return new(BaggageSubContentVisualizesBean)
	},
}

// GetBaggageSubContentVisualizesBean() 从对象池中获取BaggageSubContentVisualizesBean
func GetBaggageSubContentVisualizesBean() *BaggageSubContentVisualizesBean {
	return poolBaggageSubContentVisualizesBean.Get().(*BaggageSubContentVisualizesBean)
}

// ReleaseBaggageSubContentVisualizesBean 释放BaggageSubContentVisualizesBean
func ReleaseBaggageSubContentVisualizesBean(v *BaggageSubContentVisualizesBean) {
	poolBaggageSubContentVisualizesBean.Put(v)
}
