package ascpqcc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpqccsamplecancelitemrelationAPIResponse 魅力惠样品解除父子商品关系 API返回值
// alibaba.ascp.qcc.sample.cancel.item.relation
//
// 品控中心魅力惠样品解除父子商品关系
type AlibabaascpqccsamplecancelitemrelationAPIResponse struct {
	model.CommonResponse
	AlibabaascpqccsamplecancelitemrelationAPIResponseModel
}

// AlibabaascpqccsamplecancelitemrelationAPIResponseModel is 魅力惠样品解除父子商品关系 成功返回结果
type AlibabaascpqccsamplecancelitemrelationAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_qcc_sample_cancel_item_relation_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *SendResult `json:"result,omitempty" xml:"result,omitempty"`
}
