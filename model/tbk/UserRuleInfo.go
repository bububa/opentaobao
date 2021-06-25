package tbk

// UserRuleInfo 
type UserRuleInfo struct {

    // 用户在TOP上的openId
    OpenId   string `json:"open_id,omitempty"`

    // 每条记录离线任务生成，代表当时离线任务的时间戳
    Version   int64 `json:"version,omitempty"`

    // 用户对应的商品详细信息
    ItemList   []MapData `json:"item_list,omitempty"`

}
