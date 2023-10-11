package ascp

// SpecialSignLineRules 结构体
type SpecialSignLineRules struct {
	// 到货线路表达规则（组）
	PromiseDesRules []PromiseDesRule `json:"promise_des_rules,omitempty" xml:"promise_des_rules>promise_des_rule,omitempty"`
	// wms货主id
	WmsOwnerCode string `json:"wms_owner_code,omitempty" xml:"wms_owner_code,omitempty"`
}
