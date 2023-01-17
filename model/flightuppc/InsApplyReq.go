package flightuppc

// InsApplyReq 结构体
type InsApplyReq struct {
	// 投保参数列表，通过险种聚类
	InsProductParams []InsProductBaseParam `json:"ins_product_params,omitempty" xml:"ins_product_params>ins_product_base_param,omitempty"`
}
