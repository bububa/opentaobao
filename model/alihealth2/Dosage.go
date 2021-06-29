package alihealth2

// Dosage 
type Dosage struct {
    // 用法用量
    Items   []DosageItem `json:"items,omitempty" xml:"items>dosage_item,omitempty"`
    // 用法
    Way   string `json:"way,omitempty" xml:"way,omitempty"`
}
