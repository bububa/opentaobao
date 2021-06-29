package tmc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里金融用户授权 APIResponse
alipay.xiaodai.user.permit

阿里金融为用户开通消息通道接口
*/
type AlipayXiaodaiUserPermitAPIResponse struct {
    model.CommonResponse
    AlipayXiaodaiUserPermitResponse
}

type AlipayXiaodaiUserPermitResponse struct {
    XMLName xml.Name `xml:"alipay_xiaodai_user_permit_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
