package user

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询appstore应用订购关系 APIResponse
taobao.appstore.subscribe.get

查询appstore应用订购关系(对于新上架的多版本应用，建议使用taobao.vas.subscribe.get)
*/
type TaobaoAppstoreSubscribeGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoAppstoreSubscribeGetResponse `json:"appstore_subscribe_get_response,omitempty"` 
    TaobaoAppstoreSubscribeGetResponse
}

/* model for simplify = false
type TaobaoAppstoreSubscribeGetResponse struct {

    // 用户订购信息
    
    UserSubscribe  *struct {
        UserSubscribe  *UserSubscribe `json:"user_subscribe,omitempty"`
    } `json:"user_subscribe,omitempty"`
    

}
*/

type TaobaoAppstoreSubscribeGetResponse struct {

    // 用户订购信息
    UserSubscribe   *UserSubscribe `json:"user_subscribe,omitempty"`

}
