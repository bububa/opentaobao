package tbk

// TaobaoTbkScPublisherInfoSaveData 
/* model for simplify = false
type TaobaoTbkScPublisherInfoSaveData struct {

    // 渠道关系ID
    
    RelationId   int64 `json:"relation_id,omitempty"`
    

    // 渠道昵称
    
    AccountName   string `json:"account_name,omitempty"`
    

    // 会员运营ID
    
    SpecialId   int64 `json:"special_id,omitempty"`
    

    // 如果重复绑定会提示：”重复绑定渠道“或”重复绑定粉丝“
    
    Desc   string `json:"desc,omitempty"`
    

}
*/

// TaobaoTbkScPublisherInfoSaveData 
type TaobaoTbkScPublisherInfoSaveData struct {

    // 渠道关系ID
    RelationId   int64 `json:"relation_id,omitempty"`

    // 渠道昵称
    AccountName   string `json:"account_name,omitempty"`

    // 会员运营ID
    SpecialId   int64 `json:"special_id,omitempty"`

    // 如果重复绑定会提示：”重复绑定渠道“或”重复绑定粉丝“
    Desc   string `json:"desc,omitempty"`

}
