package qimen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
渠道间库存规则设置接口 API返回值 
taobao.qimen.inventoryrule.create

渠道间库存规则设置
*/
type TaobaoQimenInventoryruleCreateAPIResponse struct {
    model.CommonResponse
    TaobaoQimenInventoryruleCreateResponse
}

// 渠道间库存规则设置接口 成功返回结果
type TaobaoQimenInventoryruleCreateResponse struct {
    XMLName xml.Name `xml:"qimen_inventoryrule_create_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 
    Response   *ResponseDO `json:"response,omitempty" xml:"response,omitempty"`
}
