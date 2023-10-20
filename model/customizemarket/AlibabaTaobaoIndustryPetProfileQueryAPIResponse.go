package customizemarket

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTaobaoIndustryPetProfileQueryAPIResponse 用户宠物列表查询 API返回值
// alibaba.taobao.industry.pet.profile.query
//
// 用户宠物列表查询
type AlibabaTaobaoIndustryPetProfileQueryAPIResponse struct {
	model.CommonResponse
	AlibabaTaobaoIndustryPetProfileQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTaobaoIndustryPetProfileQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTaobaoIndustryPetProfileQueryAPIResponseModel).Reset()
}

// AlibabaTaobaoIndustryPetProfileQueryAPIResponseModel is 用户宠物列表查询 成功返回结果
type AlibabaTaobaoIndustryPetProfileQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_taobao_industry_pet_profile_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 参数错误
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 501
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 数据
	Object *BasePageBean `json:"object,omitempty" xml:"object,omitempty"`
	// 是否成功
	ResultStatus bool `json:"result_status,omitempty" xml:"result_status,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTaobaoIndustryPetProfileQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgInfo = ""
	m.MsgCode = ""
	m.Object = nil
	m.ResultStatus = false
}

var poolAlibabaTaobaoIndustryPetProfileQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTaobaoIndustryPetProfileQueryAPIResponse)
	},
}

// GetAlibabaTaobaoIndustryPetProfileQueryAPIResponse 从 sync.Pool 获取 AlibabaTaobaoIndustryPetProfileQueryAPIResponse
func GetAlibabaTaobaoIndustryPetProfileQueryAPIResponse() *AlibabaTaobaoIndustryPetProfileQueryAPIResponse {
	return poolAlibabaTaobaoIndustryPetProfileQueryAPIResponse.Get().(*AlibabaTaobaoIndustryPetProfileQueryAPIResponse)
}

// ReleaseAlibabaTaobaoIndustryPetProfileQueryAPIResponse 将 AlibabaTaobaoIndustryPetProfileQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTaobaoIndustryPetProfileQueryAPIResponse(v *AlibabaTaobaoIndustryPetProfileQueryAPIResponse) {
	v.Reset()
	poolAlibabaTaobaoIndustryPetProfileQueryAPIResponse.Put(v)
}
