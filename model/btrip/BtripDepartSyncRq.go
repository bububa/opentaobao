package btrip

// BtripDepartSyncRq 
type BtripDepartSyncRq struct {

    // 部门列表
    
    DepartList   []DepartSyncRq `json:"depart_list,omitempty" xml:"depart_list,omitempty"`
    

    // 第三方企业ID
    
    CorpId   string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
    

}
