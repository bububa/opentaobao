package maitix

// OpenApiPostFeeParam 
type OpenApiPostFeeParam struct {
    // 地址国标
    Address   *AddressDto `json:"address,omitempty" xml:"address,omitempty"`
    // 项目ID-必填
    ProjectId   int64 `json:"project_id,omitempty" xml:"project_id,omitempty"`
}
