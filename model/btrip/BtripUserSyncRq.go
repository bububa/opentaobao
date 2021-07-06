package btrip

// BtripUserSyncRq 结构体
type BtripUserSyncRq struct {
	// 人员列表，最大长度5000
	UserList []UserSyncRq `json:"user_list,omitempty" xml:"user_list>user_sync_rq,omitempty"`
	// 第三方企业ID
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
}
