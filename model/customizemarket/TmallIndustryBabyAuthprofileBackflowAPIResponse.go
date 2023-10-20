package customizemarket

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallIndustryBabyAuthprofileBackflowAPIResponse 孕校云回流档案 API返回值
// tmall.industry.baby.authprofile.backflow
//
// 孕校云回流档案
type TmallIndustryBabyAuthprofileBackflowAPIResponse struct {
	model.CommonResponse
	TmallIndustryBabyAuthprofileBackflowAPIResponseModel
}

// Reset 清空结构体
func (m *TmallIndustryBabyAuthprofileBackflowAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallIndustryBabyAuthprofileBackflowAPIResponseModel).Reset()
}

// TmallIndustryBabyAuthprofileBackflowAPIResponseModel is 孕校云回流档案 成功返回结果
type TmallIndustryBabyAuthprofileBackflowAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_industry_baby_authprofile_backflow_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	MsgMessage string `json:"msg_message,omitempty" xml:"msg_message,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 接口调用是否成功
	ResultStatus bool `json:"result_status,omitempty" xml:"result_status,omitempty"`
}

// Reset 清空结构体
func (m *TmallIndustryBabyAuthprofileBackflowAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgMessage = ""
	m.MsgCode = ""
	m.ResultStatus = false
}

var poolTmallIndustryBabyAuthprofileBackflowAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallIndustryBabyAuthprofileBackflowAPIResponse)
	},
}

// GetTmallIndustryBabyAuthprofileBackflowAPIResponse 从 sync.Pool 获取 TmallIndustryBabyAuthprofileBackflowAPIResponse
func GetTmallIndustryBabyAuthprofileBackflowAPIResponse() *TmallIndustryBabyAuthprofileBackflowAPIResponse {
	return poolTmallIndustryBabyAuthprofileBackflowAPIResponse.Get().(*TmallIndustryBabyAuthprofileBackflowAPIResponse)
}

// ReleaseTmallIndustryBabyAuthprofileBackflowAPIResponse 将 TmallIndustryBabyAuthprofileBackflowAPIResponse 保存到 sync.Pool
func ReleaseTmallIndustryBabyAuthprofileBackflowAPIResponse(v *TmallIndustryBabyAuthprofileBackflowAPIResponse) {
	v.Reset()
	poolTmallIndustryBabyAuthprofileBackflowAPIResponse.Put(v)
}
