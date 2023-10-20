package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpIndustryAnomalyRecourseStatusModifyAPIResponse 送货入户并安装投诉工单状态变更 API返回值
// alibaba.ascp.industry.anomaly.recourse.status.modify
//
// 送货入户并安装投诉工单状态变更
type AlibabaAscpIndustryAnomalyRecourseStatusModifyAPIResponse struct {
	model.CommonResponse
	AlibabaAscpIndustryAnomalyRecourseStatusModifyAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpIndustryAnomalyRecourseStatusModifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpIndustryAnomalyRecourseStatusModifyAPIResponseModel).Reset()
}

// AlibabaAscpIndustryAnomalyRecourseStatusModifyAPIResponseModel is 送货入户并安装投诉工单状态变更 成功返回结果
type AlibabaAscpIndustryAnomalyRecourseStatusModifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_industry_anomaly_recourse_status_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *ResultWrapper `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpIndustryAnomalyRecourseStatusModifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAscpIndustryAnomalyRecourseStatusModifyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpIndustryAnomalyRecourseStatusModifyAPIResponse)
	},
}

// GetAlibabaAscpIndustryAnomalyRecourseStatusModifyAPIResponse 从 sync.Pool 获取 AlibabaAscpIndustryAnomalyRecourseStatusModifyAPIResponse
func GetAlibabaAscpIndustryAnomalyRecourseStatusModifyAPIResponse() *AlibabaAscpIndustryAnomalyRecourseStatusModifyAPIResponse {
	return poolAlibabaAscpIndustryAnomalyRecourseStatusModifyAPIResponse.Get().(*AlibabaAscpIndustryAnomalyRecourseStatusModifyAPIResponse)
}

// ReleaseAlibabaAscpIndustryAnomalyRecourseStatusModifyAPIResponse 将 AlibabaAscpIndustryAnomalyRecourseStatusModifyAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpIndustryAnomalyRecourseStatusModifyAPIResponse(v *AlibabaAscpIndustryAnomalyRecourseStatusModifyAPIResponse) {
	v.Reset()
	poolAlibabaAscpIndustryAnomalyRecourseStatusModifyAPIResponse.Put(v)
}
