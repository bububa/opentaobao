package maitix

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimaitixordercancelAPIResponse 大麦-库存释放 API返回值
// alibaba.damai.maitix.order.cancel
//
// 库存释放
type AlibabadamaimaitixordercancelAPIResponse struct {
	model.CommonResponse
	AlibabadamaimaitixordercancelAPIResponseModel
}

// AlibabadamaimaitixordercancelAPIResponseModel is 大麦-库存释放 成功返回结果
type AlibabadamaimaitixordercancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_maitix_order_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果集
	Result *MxResult `json:"result,omitempty" xml:"result,omitempty"`
}
