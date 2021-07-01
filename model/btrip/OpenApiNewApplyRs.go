package btrip

// OpenApiNewApplyRs 
type OpenApiNewApplyRs struct {
    // 用户传入审批单id
    ThirdpartApplyId   string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
    // 商旅审批单id
    ApplyId   string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
}
