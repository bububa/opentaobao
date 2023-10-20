package omniorder

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoomniorderdtdconsignAPIResponse 门店自送发货 API返回值
// taobao.omniorder.dtd.consign
//
// 该接口触发门店自送发货，推进淘系订单状态为发货，为消费者发送核销码短信，并将物流信息写入订单
type TaobaoomniorderdtdconsignAPIResponse struct {
	model.CommonResponse
	TaobaoomniorderdtdconsignAPIResponseModel
}

// TaobaoomniorderdtdconsignAPIResponseModel is 门店自送发货 成功返回结果
type TaobaoomniorderdtdconsignAPIResponseModel struct {
	XMLName xml.Name `xml:"omniorder_dtd_consign_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码，为0表示成功，非0表示失败
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}
