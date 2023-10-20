package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterServicestoreUpdateAPIResponse 修改门店信息 API返回值
// tmall.servicecenter.servicestore.update
//
// 用于修改门店/网点信息。多个业务共用
type TmallServicecenterServicestoreUpdateAPIResponse struct {
	model.CommonResponse
	TmallServicecenterServicestoreUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterServicestoreUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterServicestoreUpdateAPIResponseModel).Reset()
}

// TmallServicecenterServicestoreUpdateAPIResponseModel is 修改门店信息 成功返回结果
type TmallServicecenterServicestoreUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_servicestore_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 方法调用结果
	Result *TmallServicecenterServicestoreUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterServicestoreUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterServicestoreUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterServicestoreUpdateAPIResponse)
	},
}

// GetTmallServicecenterServicestoreUpdateAPIResponse 从 sync.Pool 获取 TmallServicecenterServicestoreUpdateAPIResponse
func GetTmallServicecenterServicestoreUpdateAPIResponse() *TmallServicecenterServicestoreUpdateAPIResponse {
	return poolTmallServicecenterServicestoreUpdateAPIResponse.Get().(*TmallServicecenterServicestoreUpdateAPIResponse)
}

// ReleaseTmallServicecenterServicestoreUpdateAPIResponse 将 TmallServicecenterServicestoreUpdateAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterServicestoreUpdateAPIResponse(v *TmallServicecenterServicestoreUpdateAPIResponse) {
	v.Reset()
	poolTmallServicecenterServicestoreUpdateAPIResponse.Put(v)
}
