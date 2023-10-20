package openmall

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenmallrefundcloseAPIResponse 关闭OpenMall退款单 API返回值
// taobao.openmall.refund.close
//
// 关闭OpenMall退款单
type TaobaoopenmallrefundcloseAPIResponse struct {
	model.CommonResponse
	TaobaoopenmallrefundcloseAPIResponseModel
}

// TaobaoopenmallrefundcloseAPIResponseModel is 关闭OpenMall退款单 成功返回结果
type TaobaoopenmallrefundcloseAPIResponseModel struct {
	XMLName xml.Name `xml:"openmall_refund_close_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否关闭成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
