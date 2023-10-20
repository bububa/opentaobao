package alihealth2

// DosageItem 结构体
type DosageItem struct {
	// 使用量单位
	Unit string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 用量类型（SINGLEDAILYDOSE(&#34;单日剂量&#34;),     SINGLEDOSE(&#34;单次剂量&#34;),     DRUGCOURSE(&#34;药品疗程&#34;),     FREQUENCYADMINISTRATION(&#34;给药频次&#34;)）
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 使用量
	Value float64 `json:"value,omitempty" xml:"value,omitempty"`
}
