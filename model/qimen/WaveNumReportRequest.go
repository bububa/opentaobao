package qimen

// WaveNumReportRequest 结构体
type WaveNumReportRequest struct {
	// 发货单号
	Orders []Order `json:"orders,omitempty" xml:"orders>order,omitempty"`
	// 波次号
	WaveNum string `json:"waveNum,omitempty" xml:"waveNum,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoqimenwavenumreportMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}
