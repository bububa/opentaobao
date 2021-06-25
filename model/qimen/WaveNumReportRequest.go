package qimen

// WaveNumReportRequest 
type WaveNumReportRequest struct {

    // 波次号
    WaveNum   string `json:"waveNum,omitempty"`

    // 发货单号
    Orders   []Order `json:"orders,omitempty"`

    // 扩展属性
    ExtendProps   *Map `json:"extendProps,omitempty"`

}
