package tmallcar

// XCarSysModelDto 结构体
type XCarSysModelDto struct {
	// 车型最高降幅
	CarModelMaximumDecline string `json:"car_model_maximum_decline,omitempty" xml:"car_model_maximum_decline,omitempty"`
	// 燃油费用
	FuelCost string `json:"fuel_cost,omitempty" xml:"fuel_cost,omitempty"`
	// 保险费用
	InsuranceFree string `json:"insurance_free,omitempty" xml:"insurance_free,omitempty"`
	// 养车费用
	KeepCarFree string `json:"keep_car_free,omitempty" xml:"keep_car_free,omitempty"`
	// 本地参考价
	LocalRefPriceRange string `json:"local_ref_price_range,omitempty" xml:"local_ref_price_range,omitempty"`
	// 保养费用
	MaintainCost string `json:"maintain_cost,omitempty" xml:"maintain_cost,omitempty"`
	// 厂商指导价
	ManuGuiPrice string `json:"manu_gui_price,omitempty" xml:"manu_gui_price,omitempty"`
	// 媒体库，超大json
	MediaConfig string `json:"media_config,omitempty" xml:"media_config,omitempty"`
	// 品牌pid
	BrandPid int64 `json:"brand_pid,omitempty" xml:"brand_pid,omitempty"`
	// 品牌vid
	BrandVid int64 `json:"brand_vid,omitempty" xml:"brand_vid,omitempty"`
	// 车系属性id
	LinePid int64 `json:"line_pid,omitempty" xml:"line_pid,omitempty"`
	// 车系值id
	LineVid int64 `json:"line_vid,omitempty" xml:"line_vid,omitempty"`
	// 型号属性id
	ModelPid int64 `json:"model_pid,omitempty" xml:"model_pid,omitempty"`
	// 型号值id
	ModelVid int64 `json:"model_vid,omitempty" xml:"model_vid,omitempty"`
	// 1.新车上市 0.非新车上市
	NewCarModel int64 `json:"new_car_model,omitempty" xml:"new_car_model,omitempty"`
	// 状态0.无效 1有效
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 年款属性id
	YearPid int64 `json:"year_pid,omitempty" xml:"year_pid,omitempty"`
	// 年款值id
	YearVid int64 `json:"year_vid,omitempty" xml:"year_vid,omitempty"`
}
