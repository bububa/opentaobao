package idleitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
免费送商品发送 API返回值 
alibaba.idle.item.idlecoin.add

免费送商品发布
*/
type AlibabaIdleItemIdlecoinAddAPIResponse struct {
    model.CommonResponse
    AlibabaIdleItemIdlecoinAddResponse
}

// 免费送商品发送 成功返回结果
type AlibabaIdleItemIdlecoinAddResponse struct {
    XMLName xml.Name `xml:"alibaba_idle_item_idlecoin_add_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *IdleResultDO `json:"result,omitempty" xml:"result,omitempty"`
}
