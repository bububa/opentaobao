package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpindustryworkermodifyAPIResponse 送货入户并安装修改师傅信息 API返回值
// alibaba.ascp.industry.worker.modify
//
// 送货入户并安装修改师傅信息
type AlibabaascpindustryworkermodifyAPIResponse struct {
	model.CommonResponse
	AlibabaascpindustryworkermodifyAPIResponseModel
}

// AlibabaascpindustryworkermodifyAPIResponseModel is 送货入户并安装修改师傅信息 成功返回结果
type AlibabaascpindustryworkermodifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_industry_worker_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *ResultWrapper `json:"result,omitempty" xml:"result,omitempty"`
}
