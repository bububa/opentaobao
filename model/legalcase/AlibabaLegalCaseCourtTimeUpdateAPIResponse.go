package legalcase

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegalcasecourttimeupdateAPIResponse 开庭时间变更 API返回值
// alibaba.legal.case.court.time.update
//
// 修改案件的开庭时间
type AlibabalegalcasecourttimeupdateAPIResponse struct {
	model.CommonResponse
	AlibabalegalcasecourttimeupdateAPIResponseModel
}

// AlibabalegalcasecourttimeupdateAPIResponseModel is 开庭时间变更 成功返回结果
type AlibabalegalcasecourttimeupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_case_court_time_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
