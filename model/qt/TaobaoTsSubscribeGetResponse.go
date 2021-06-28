package qt

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝服务订购关系查询 APIResponse
taobao.ts.subscribe.get

ts订购关系状态查询. 暂只支持1口价服务.
*/
type TaobaoTsSubscribeGetAPIResponse struct {
    model.CommonResponse
    TaobaoTsSubscribeGetResponse
}

type TaobaoTsSubscribeGetResponse struct {
    XMLName xml.Name `xml:"ts_subscribe_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 订购关系对象
    
    ServiceSubscribe   *ServiceSubscribe `json:"service_subscribe,omitempty" xml:"service_subscribe,omitempty"`

    
}
