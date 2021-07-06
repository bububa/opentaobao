package alicom

// EcardItemDo 结构体
type EcardItemDo struct {
	// begin_of_validity
	BeginOfValidity string `json:"begin_of_validity,omitempty" xml:"begin_of_validity,omitempty"`
	// cz_activity_desc
	CzActivityDesc string `json:"cz_activity_desc,omitempty" xml:"cz_activity_desc,omitempty"`
	// err_msg
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// face_price
	FacePrice string `json:"face_price,omitempty" xml:"face_price,omitempty"`
	// flow_type
	FlowType string `json:"flow_type,omitempty" xml:"flow_type,omitempty"`
	// is_invalid_monthly
	IsInvalidMonthly string `json:"is_invalid_monthly,omitempty" xml:"is_invalid_monthly,omitempty"`
	// item_id
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// item_name
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// mobile
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// price_ext_content
	PriceExtContent string `json:"price_ext_content,omitempty" xml:"price_ext_content,omitempty"`
	// promotion_price
	PromotionPrice string `json:"promotion_price,omitempty" xml:"promotion_price,omitempty"`
	// seller
	Seller string `json:"seller,omitempty" xml:"seller,omitempty"`
	// send_desc
	SendDesc string `json:"send_desc,omitempty" xml:"send_desc,omitempty"`
	// simple_desc
	SimpleDesc string `json:"simple_desc,omitempty" xml:"simple_desc,omitempty"`
	// std_price
	StdPrice string `json:"std_price,omitempty" xml:"std_price,omitempty"`
	// term_of_validity
	TermOfValidity string `json:"term_of_validity,omitempty" xml:"term_of_validity,omitempty"`
	// time_limit
	TimeLimit string `json:"time_limit,omitempty" xml:"time_limit,omitempty"`
	// title
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// result
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
	// is_activity
	IsActivity bool `json:"is_activity,omitempty" xml:"is_activity,omitempty"`
	// is_guan_f_seller
	IsGuanFSeller bool `json:"is_guan_f_seller,omitempty" xml:"is_guan_f_seller,omitempty"`
	// is_hui_item
	IsHuiItem bool `json:"is_hui_item,omitempty" xml:"is_hui_item,omitempty"`
	// is_white
	IsWhite bool `json:"is_white,omitempty" xml:"is_white,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
