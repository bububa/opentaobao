package baichuan

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaBaichuanCtgUserRelationAPIResponse 用户 API返回值
// alibaba.baichuan.ctg.user.relation
//
// 提供给优酷查询道长和淘宝账户的绑定关系
type AlibabaBaichuanCtgUserRelationAPIResponse struct {
	model.CommonResponse
	AlibabaBaichuanCtgUserRelationAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaBaichuanCtgUserRelationAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaBaichuanCtgUserRelationAPIResponseModel).Reset()
}

// AlibabaBaichuanCtgUserRelationAPIResponseModel is 用户 成功返回结果
type AlibabaBaichuanCtgUserRelationAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_baichuan_ctg_user_relation_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的整体结果
	Result *AlibabaBaichuanCtgUserRelationResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaBaichuanCtgUserRelationAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaBaichuanCtgUserRelationAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaBaichuanCtgUserRelationAPIResponse)
	},
}

// GetAlibabaBaichuanCtgUserRelationAPIResponse 从 sync.Pool 获取 AlibabaBaichuanCtgUserRelationAPIResponse
func GetAlibabaBaichuanCtgUserRelationAPIResponse() *AlibabaBaichuanCtgUserRelationAPIResponse {
	return poolAlibabaBaichuanCtgUserRelationAPIResponse.Get().(*AlibabaBaichuanCtgUserRelationAPIResponse)
}

// ReleaseAlibabaBaichuanCtgUserRelationAPIResponse 将 AlibabaBaichuanCtgUserRelationAPIResponse 保存到 sync.Pool
func ReleaseAlibabaBaichuanCtgUserRelationAPIResponse(v *AlibabaBaichuanCtgUserRelationAPIResponse) {
	v.Reset()
	poolAlibabaBaichuanCtgUserRelationAPIResponse.Put(v)
}
