package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
零售商品发货 API返回值 
alibaba.mj.oc.outbound

用于接收发货的数据
*/
type AlibabaMjOcOutboundAPIResponse struct {
    model.CommonResponse
    AlibabaMjOcOutboundAPIResponseModel
}

// 零售商品发货 成功返回结果
type AlibabaMjOcOutboundAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_mj_oc_outbound_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // success
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
