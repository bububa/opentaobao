package alihealth2

// PrescriptionItem 
type PrescriptionItem struct {

    // 用法用量，启用用法用量规则的条件下必传
    
    Dosage   *Dosage `json:"dosage,omitempty" xml:"dosage,omitempty"`
    

    // 购买数量，启用限购的规则条件下必传
    
    BuyAmount   *BuyAmount `json:"buy_amount,omitempty" xml:"buy_amount,omitempty"`
    

    // 购买的药品
    
    Drug   *Drug `json:"drug,omitempty" xml:"drug,omitempty"`
    

}
