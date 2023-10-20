package omniorder

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoomniorderdtdconsumeAPIResponse 门店自送对码进行核销 API返回值
// taobao.omniorder.dtd.consume
//
// 该接口根据传入的码及订单信息，如果码与订单一致，则对门店自送服务进行核销。
type TaobaoomniorderdtdconsumeAPIResponse struct {
	model.CommonResponse
	TaobaoomniorderdtdconsumeAPIResponseModel
}

// TaobaoomniorderdtdconsumeAPIResponseModel is 门店自送对码进行核销 成功返回结果
type TaobaoomniorderdtdconsumeAPIResponseModel struct {
	XMLName xml.Name `xml:"omniorder_dtd_consume_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码，为0表示成功，非0表示失败
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 错误西溪
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}
