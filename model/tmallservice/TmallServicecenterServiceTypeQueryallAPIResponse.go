package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterServiceTypeQueryallAPIResponse 服务供应链服务类型 API返回值
// tmall.servicecenter.service.type.queryall
//
// 查询天猫服务类型列表
type TmallServicecenterServiceTypeQueryallAPIResponse struct {
	model.CommonResponse
	TmallServicecenterServiceTypeQueryallAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterServiceTypeQueryallAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterServiceTypeQueryallAPIResponseModel).Reset()
}

// TmallServicecenterServiceTypeQueryallAPIResponseModel is 服务供应链服务类型 成功返回结果
type TmallServicecenterServiceTypeQueryallAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_service_type_queryall_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ResultBase `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterServiceTypeQueryallAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterServiceTypeQueryallAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterServiceTypeQueryallAPIResponse)
	},
}

// GetTmallServicecenterServiceTypeQueryallAPIResponse 从 sync.Pool 获取 TmallServicecenterServiceTypeQueryallAPIResponse
func GetTmallServicecenterServiceTypeQueryallAPIResponse() *TmallServicecenterServiceTypeQueryallAPIResponse {
	return poolTmallServicecenterServiceTypeQueryallAPIResponse.Get().(*TmallServicecenterServiceTypeQueryallAPIResponse)
}

// ReleaseTmallServicecenterServiceTypeQueryallAPIResponse 将 TmallServicecenterServiceTypeQueryallAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterServiceTypeQueryallAPIResponse(v *TmallServicecenterServiceTypeQueryallAPIResponse) {
	v.Reset()
	poolTmallServicecenterServiceTypeQueryallAPIResponse.Put(v)
}
