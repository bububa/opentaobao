package servicecenter

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商查询前置补贴红包的最新数据 API返回值 
taobao.recycle.ofnpreredpacket.get

服务商查询前置补贴红包的最新数据
*/
type TaobaoRecycleOfnpreredpacketGetAPIResponse struct {
    model.CommonResponse
    TaobaoRecycleOfnpreredpacketGetAPIResponseModel
}

// 服务商查询前置补贴红包的最新数据 成功返回结果
type TaobaoRecycleOfnpreredpacketGetAPIResponseModel struct {
    XMLName xml.Name `xml:"recycle_ofnpreredpacket_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 前置补贴红包
    Data   *OfnPreRedPacketDto `json:"data,omitempty" xml:"data,omitempty"`
}
