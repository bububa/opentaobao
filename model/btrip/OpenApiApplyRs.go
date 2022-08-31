package btrip

// OpenApiApplyRs 结构体
type OpenApiApplyRs struct {
	// 外部申请单id
	ThirdpartApplyId string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	// 第三方业务id
	ThirdpartBusinessId string `json:"thirdpart_business_id,omitempty" xml:"thirdpart_business_id,omitempty"`
	// 商旅申请单id
	ApplyId int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
}
