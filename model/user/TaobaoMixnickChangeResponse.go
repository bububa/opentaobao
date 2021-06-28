package user

import (
    "github.com/bububa/opentaobao/model"
)

/* 
新旧mixnick互转 APIResponse
taobao.mixnick.change

如果采用老的Appkey获取MixNick A’， 发现DB里 new MixNick为null，则使用新的Appkey调API来换取A’’；如果采用新的Appkey获取A’’，发现DB里不存在当前A’’ 的记录时，先用老Appkey调API得到A’ 查询是否存在A用户的记录，如果已经存在，则关联，否则新增一条MixNick为null的新用户记录。
*/
type TaobaoMixnickChangeAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoMixnickChangeResponse `json:"mixnick_change_response,omitempty"` 
    TaobaoMixnickChangeResponse
}

/* model for simplify = false
type TaobaoMixnickChangeResponse struct {

    // 是否成功
    
    RetSuccess   bool `json:"ret_success,omitempty"`
    

    // 根据dstAppkey算出的mixnick
    
    Mixnick   string `json:"mixnick,omitempty"`
    

}
*/

type TaobaoMixnickChangeResponse struct {

    // 是否成功
    RetSuccess   bool `json:"ret_success,omitempty"`

    // 根据dstAppkey算出的mixnick
    Mixnick   string `json:"mixnick,omitempty"`

}
