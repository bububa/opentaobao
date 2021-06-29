package alihealth2

// DosageItem 
type DosageItem struct {
    // 使用量单位
    Unit   string `json:"unit,omitempty" xml:"unit,omitempty"`
    // 用量类型（SINGLEDAILYDOSE("单日剂量"),     SINGLEDOSE("单次剂量"),     DRUGCOURSE("药品疗程"),     FREQUENCYADMINISTRATION("给药频次")）
    Type   string `json:"type,omitempty" xml:"type,omitempty"`
    // 使用量
    Value   int64 `json:"value,omitempty" xml:"value,omitempty"`
}
