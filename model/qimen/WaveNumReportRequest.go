package qimen

import (
	"sync"
)

// WaveNumReportRequest 结构体
type WaveNumReportRequest struct {
	// 发货单号
	Orders []Order `json:"orders,omitempty" xml:"orders>order,omitempty"`
	// 波次号
	WaveNum string `json:"waveNum,omitempty" xml:"waveNum,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoQimenWavenumReportMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}

var poolWaveNumReportRequest = sync.Pool{
	New: func() any {
		return new(WaveNumReportRequest)
	},
}

// GetWaveNumReportRequest() 从对象池中获取WaveNumReportRequest
func GetWaveNumReportRequest() *WaveNumReportRequest {
	return poolWaveNumReportRequest.Get().(*WaveNumReportRequest)
}

// ReleaseWaveNumReportRequest 释放WaveNumReportRequest
func ReleaseWaveNumReportRequest(v *WaveNumReportRequest) {
	v.Orders = v.Orders[:0]
	v.WaveNum = ""
	v.ExtendProps = nil
	poolWaveNumReportRequest.Put(v)
}
