package qimen

// WaveNumReportRequest 
/* model for simplify = false
type WaveNumReportRequest struct {

    // 波次号
    
    WaveNum   string `json:"waveNum,omitempty"`
    

    // 发货单号
    
    Orders  struct {
        Order  []Order `json:"order,omitempty"`
    } `json:"orders,omitempty"`
    

    // 扩展属性
    
    ExtendProps  *struct {
        Map  *Map `json:"map,omitempty"`
    } `json:"extendProps,omitempty"`
    

}
*/

// WaveNumReportRequest 
type WaveNumReportRequest struct {

    // 波次号
    WaveNum   string `json:"waveNum,omitempty"`

    // 发货单号
    Orders   []Order `json:"orders,omitempty"`

    // 扩展属性
    ExtendProps   *Map `json:"extendProps,omitempty"`

}
