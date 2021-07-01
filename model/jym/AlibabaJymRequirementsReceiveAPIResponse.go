package jym

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaJymRequirementsReceiveAPIResponse
交易猫需求接单接口 API返回值
alibaba.jym.requirements.receive

交易猫需求接单接口 */
type AlibabaJymRequirementsReceiveAPIResponse struct {
	model.CommonResponse
	AlibabaJymRequirementsReceiveAPIResponseModel
}

// AlibabaJymRequirementsReceiveAPIResponseModel is 交易猫需求接单接口 成功返回结果
type AlibabaJymRequirementsReceiveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_jym_requirements_receive_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	ResultDto *AlibabaJymRequirementsReceiveResultDto `json:"result_dto,omitempty" xml:"result_dto,omitempty"`
}
