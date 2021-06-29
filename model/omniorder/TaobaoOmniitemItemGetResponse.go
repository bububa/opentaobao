package omniorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取全渠道门店商品 APIResponse
taobao.omniitem.item.get

通过门店id/类目id/商品id单个或多个参数组合查询全渠道门店商品信息
*/
type TaobaoOmniitemItemGetAPIResponse struct {
    model.CommonResponse
    TaobaoOmniitemItemGetResponse
}

type TaobaoOmniitemItemGetResponse struct {
    XMLName xml.Name `xml:"omniitem_item_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *PageResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
