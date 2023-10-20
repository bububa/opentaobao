package medicalbase

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthmedicalbaseorderstatussyncAPIResponse 号源直连订单状态同步接口 API返回值
// alibaba.alihealth.medicalbase.order.status.sync
//
// 互联网医院isv批量通过接口批量导入
type AlibabaalihealthmedicalbaseorderstatussyncAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthmedicalbaseorderstatussyncAPIResponseModel
}

// AlibabaalihealthmedicalbaseorderstatussyncAPIResponseModel is 号源直连订单状态同步接口 成功返回结果
type AlibabaalihealthmedicalbaseorderstatussyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_medicalbase_order_status_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 和三方交互最外层model对象
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
