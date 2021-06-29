package servicecenter

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
订购关系查询 API返回值 
taobao.vas.subscribe.get

用于ISV根据登录进来的淘宝会员名查询该为该会员开通哪些收费项目，ISV只能查询自己名下的应用及收费项目的订购情况
*/
type TaobaoVasSubscribeGetAPIResponse struct {
    model.CommonResponse
    TaobaoVasSubscribeGetResponse
}

// 订购关系查询 成功返回结果
type TaobaoVasSubscribeGetResponse struct {
    XMLName xml.Name `xml:"vas_subscribe_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 用户订购信息
    ArticleUserSubscribes   []ArticleUserSubscribe `json:"article_user_subscribes,omitempty" xml:"article_user_subscribes>article_user_subscribe,omitempty"`
}
