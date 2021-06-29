package drugtrace

// CodeToBill 
type CodeToBill struct {

    // 追溯码
    
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    

    // codeToBill列表
    
    BillIdentityList   []BillIdentity `json:"bill_identity_list,omitempty" xml:"bill_identity_list,omitempty"`
    

}
