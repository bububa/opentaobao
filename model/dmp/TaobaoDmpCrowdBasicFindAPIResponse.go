package dmp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDmpCrowdBasicFindAPIResponse DMP_BP版人群列表查询 API返回值
// taobao.dmp.crowd.basic.find
//
// DMP_BP版人群列表查询
type TaobaoDmpCrowdBasicFindAPIResponse struct {
	model.CommonResponse
	TaobaoDmpCrowdBasicFindAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoDmpCrowdBasicFindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoDmpCrowdBasicFindAPIResponseModel).Reset()
}

// TaobaoDmpCrowdBasicFindAPIResponseModel is DMP_BP版人群列表查询 成功返回结果
type TaobaoDmpCrowdBasicFindAPIResponseModel struct {
	XMLName xml.Name `xml:"dmp_crowd_basic_find_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 人群对象
	Result []CrowdDto `json:"result,omitempty" xml:"result>crowd_dto,omitempty"`
	// 错误码
	ResultErrorCode string `json:"result_error_code,omitempty" xml:"result_error_code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 分页对象
	Pager *Pager `json:"pager,omitempty" xml:"pager,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoDmpCrowdBasicFindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = m.Result[:0]
	m.ResultErrorCode = ""
	m.Message = ""
	m.Pager = nil
	m.IsSuccess = false
}

var poolTaobaoDmpCrowdBasicFindAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoDmpCrowdBasicFindAPIResponse)
	},
}

// GetTaobaoDmpCrowdBasicFindAPIResponse 从 sync.Pool 获取 TaobaoDmpCrowdBasicFindAPIResponse
func GetTaobaoDmpCrowdBasicFindAPIResponse() *TaobaoDmpCrowdBasicFindAPIResponse {
	return poolTaobaoDmpCrowdBasicFindAPIResponse.Get().(*TaobaoDmpCrowdBasicFindAPIResponse)
}

// ReleaseTaobaoDmpCrowdBasicFindAPIResponse 将 TaobaoDmpCrowdBasicFindAPIResponse 保存到 sync.Pool
func ReleaseTaobaoDmpCrowdBasicFindAPIResponse(v *TaobaoDmpCrowdBasicFindAPIResponse) {
	v.Reset()
	poolTaobaoDmpCrowdBasicFindAPIResponse.Put(v)
}
