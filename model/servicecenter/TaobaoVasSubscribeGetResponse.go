package servicecenter

import (
    "github.com/bububa/opentaobao/model"
)

/* 
订购关系查询 APIResponse
taobao.vas.subscribe.get

用于ISV根据登录进来的淘宝会员名查询该为该会员开通哪些收费项目，ISV只能查询自己名下的应用及收费项目的订购情况
*/
type TaobaoVasSubscribeGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoVasSubscribeGetResponse `json:"vas_subscribe_get_response,omitempty"` 
    TaobaoVasSubscribeGetResponse
}

/* model for simplify = false
type TaobaoVasSubscribeGetResponse struct {

    // 用户订购信息
    
    ArticleUserSubscribes  struct {
        ArticleUserSubscribe  []ArticleUserSubscribe `json:"article_user_subscribe,omitempty"`
    } `json:"article_user_subscribes,omitempty"`
    

}
*/

type TaobaoVasSubscribeGetResponse struct {

    // 用户订购信息
    ArticleUserSubscribes   []ArticleUserSubscribe `json:"article_user_subscribes,omitempty"`

}
