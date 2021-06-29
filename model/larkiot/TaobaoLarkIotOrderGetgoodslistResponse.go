package larkiot

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
iot渠道获取卖品信息 APIResponse
taobao.lark.iot.order.getgoodslist

iot无人超市服务商通过接口获取影院的可售卖品数据
*/
type TaobaoLarkIotOrderGetgoodslistAPIResponse struct {
    model.CommonResponse
    TaobaoLarkIotOrderGetgoodslistResponse
}

type TaobaoLarkIotOrderGetgoodslistResponse struct {
    XMLName xml.Name `xml:"lark_iot_order_getgoodslist_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 卖品信息列表
    
    Data   *BizSingleResult `json:"data,omitempty" xml:"data,omitempty"`

    
}
