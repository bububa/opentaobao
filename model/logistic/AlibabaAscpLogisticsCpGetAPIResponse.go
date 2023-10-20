package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascplogisticscpgetAPIResponse 快递公司资源列表查询接口 API返回值
// alibaba.ascp.logistics.cp.get
//
// 快递公司资源列表查询接口
type AlibabaascplogisticscpgetAPIResponse struct {
	model.CommonResponse
	AlibabaascplogisticscpgetAPIResponseModel
}

// AlibabaascplogisticscpgetAPIResponseModel is 快递公司资源列表查询接口 成功返回结果
type AlibabaascplogisticscpgetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_logistics_cp_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *ResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
