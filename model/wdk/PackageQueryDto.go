package wdk

// PackageQueryDto 
type PackageQueryDto struct {
    // 同城令牌号
    TokenCode   string `json:"token_code,omitempty" xml:"token_code,omitempty"`
    // 仓Code
    NodeCode   string `json:"node_code,omitempty" xml:"node_code,omitempty"`
    // 作业单元ID
    WorkUnitId   string `json:"work_unit_id,omitempty" xml:"work_unit_id,omitempty"`
}
