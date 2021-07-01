package maitix

// OpenApiPostFeeDto 结构体
type OpenApiPostFeeDto struct {
	// 运费金额，单位分
	PostFee int64 `json:"post_fee,omitempty" xml:"post_fee,omitempty"`
}
