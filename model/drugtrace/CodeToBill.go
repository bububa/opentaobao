package drugtrace

// CodeToBill 结构体
type CodeToBill struct {
	// codeToBill列表
	BillIdentityList []BillIdentity `json:"bill_identity_list,omitempty" xml:"bill_identity_list>bill_identity,omitempty"`
	// 追溯码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
}
