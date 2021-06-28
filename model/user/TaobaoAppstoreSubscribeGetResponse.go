package user

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询appstore应用订购关系 APIResponse
taobao.appstore.subscribe.get

查询appstore应用订购关系(对于新上架的多版本应用，建议使用taobao.vas.subscribe.get)
*/
type TaobaoAppstoreSubscribeGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"appstore_subscribe_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 用户订购信息
    
    UserSubscribe   *UserSubscribe `json:"user_subscribe,omitempty" xml:"