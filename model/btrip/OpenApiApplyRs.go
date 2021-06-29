package btrip

// OpenApiApplyRs 
type OpenApiApplyRs struct {
    // 商旅申请单id
    ApplyId   int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
    // 外部申请单id
    ThirdpartApplyId   string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
    // 第三方业务id
    ThirdpartBusinessId   string `json:"thirdpart_business_id,omitempty" xml:"thirdpart_business_id,omitempty"`
}
