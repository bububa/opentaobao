package btrip

// BtripUserSyncRq 
type BtripUserSyncRq struct {

    // 第三方企业ID
    
    CorpId   string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
    

    // 人员列表，最大长度5000
    
    UserList   []UserSyncRq `json:"user_list,omitempty" xml:"user_list,omitempty"`
    

}
