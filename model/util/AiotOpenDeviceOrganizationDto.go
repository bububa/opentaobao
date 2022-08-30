package util

// AiotOpenDeviceOrganizationDto 结构体
type AiotOpenDeviceOrganizationDto struct {
	// 1级组织架构code
	L1Code string `json:"l1_code,omitempty" xml:"l1_code,omitempty"`
	// 1级组织架构名
	L1Name string `json:"l1_name,omitempty" xml:"l1_name,omitempty"`
	// 2级组织架构code
	L2Code string `json:"l2_code,omitempty" xml:"l2_code,omitempty"`
	// 2级组织架构名
	L2Name string `json:"l2_name,omitempty" xml:"l2_name,omitempty"`
	// 3级组织架构code
	L3Code string `json:"l3_code,omitempty" xml:"l3_code,omitempty"`
	// 3级组织架构名
	L3Name string `json:"l3_name,omitempty" xml:"l3_name,omitempty"`
	// 4级组织架构code
	L4Code string `json:"l4_code,omitempty" xml:"l4_code,omitempty"`
	// 4级组织架构名
	L4Name string `json:"l4_name,omitempty" xml:"l4_name,omitempty"`
	// 5级组织架构code
	L5Code string `json:"l5_code,omitempty" xml:"l5_code,omitempty"`
	// 5级组织架构名
	L5Name string `json:"l5_name,omitempty" xml:"l5_name,omitempty"`
}
