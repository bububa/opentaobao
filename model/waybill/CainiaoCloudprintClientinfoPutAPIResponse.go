package waybill

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaoCloudprintClientinfoPutAPIResponse 云打印客户端监控信息收集 API返回值
// cainiao.cloudprint.clientinfo.put
//
// 云打印客户端监控信息收集
type CainiaoCloudprintClientinfoPutAPIResponse struct {
	model.CommonResponse
	CainiaoCloudprintClientinfoPutAPIResponseModel
}

// CainiaoCloudprintClientinfoPutAPIResponseModel is 云打印客户端监控信息收集 成功返回结果
type CainiaoCloudprintClientinfoPutAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_cloudprint_clientinfo_put_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *CloudPrintBaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
