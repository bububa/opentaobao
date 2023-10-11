package drugtrace

// CodeInfo 结构体
type CodeInfo struct {
	// 码状态（I核注O核销A激活C注销E错误码）
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 追溯码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 码等级--展示等级 【相当于包装等级，1代表最大展示等级, 如：申请的包装比例是1:5:10, 对应的码展示等级就是 1、2、3, 代表大码、中码、小码】
	CodeLevel string `json:"code_level,omitempty" xml:"code_level,omitempty"`
	// 父码
	ParentCode string `json:"parent_code,omitempty" xml:"parent_code,omitempty"`
	// 码等级【1代表最小码   如：申请的包装比例是1:5:10, 对应的码等级就是3、2、1, 代表大码、中码、小码】
	CodePackLevel string `json:"code_pack_level,omitempty" xml:"code_pack_level,omitempty"`
}
