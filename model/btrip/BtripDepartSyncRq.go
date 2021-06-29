package btrip

// BtripDepartSyncRq 
type BtripDepartSyncRq struct {
    // 部门列表
    DepartList   []DepartSyncRq `json:"depart_list,omitempty" xml:"depart_list>depart_sync_rq,omitempty"`
    // 第三方企业ID
    CorpId   string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
}
