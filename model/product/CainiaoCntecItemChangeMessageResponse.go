package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品变更消息 APIResponse
cainiao.cntec.item.change.message

供货商商品信息变更消息
*/
type CainiaoCntecItemChangeMessageAPIResponse struct {
    model.CommonResponse
    CainiaoCntecItemChangeMessageResponse
}

type CainiaoCntecItemChangeMessageResponse struct {
    XMLName xml.Name `xml:"cainiao_cntec_item_change_message_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 调用返回的result结构体
    
    Result   *CainiaoCntecItemChangeMessageResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
