package legalcase

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegalcaseentrustgetAPIResponse 委托 API返回值
// alibaba.legal.case.entrust.get
//
// 获取委托案件的基本信息
type AlibabalegalcaseentrustgetAPIResponse struct {
	model.CommonResponse
	AlibabalegalcaseentrustgetAPIResponseModel
}

// AlibabalegalcaseentrustgetAPIResponseModel is 委托 成功返回结果
type AlibabalegalcaseentrustgetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_case_entrust_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
