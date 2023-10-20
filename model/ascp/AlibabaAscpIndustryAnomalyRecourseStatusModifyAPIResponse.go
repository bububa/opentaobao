package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpindustryanomalyrecoursestatusmodifyAPIResponse 送货入户并安装投诉工单状态变更 API返回值
// alibaba.ascp.industry.anomaly.recourse.status.modify
//
// 送货入户并安装投诉工单状态变更
type AlibabaascpindustryanomalyrecoursestatusmodifyAPIResponse struct {
	model.CommonResponse
	AlibabaascpindustryanomalyrecoursestatusmodifyAPIResponseModel
}

// AlibabaascpindustryanomalyrecoursestatusmodifyAPIResponseModel is 送货入户并安装投诉工单状态变更 成功返回结果
type AlibabaascpindustryanomalyrecoursestatusmodifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_industry_anomaly_recourse_status_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *ResultWrapper `json:"result,omitempty" xml:"result,omitempty"`
}
