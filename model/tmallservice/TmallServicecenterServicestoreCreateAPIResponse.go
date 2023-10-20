package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterServicestoreCreateAPIResponse 创建门店 API返回值
// tmall.servicecenter.servicestore.create
//
// 用于创建门店/网点。多个业务共用
type TmallServicecenterServicestoreCreateAPIResponse struct {
	model.CommonResponse
	TmallServicecenterServicestoreCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterServicestoreCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterServicestoreCreateAPIResponseModel).Reset()
}

// TmallServicecenterServicestoreCreateAPIResponseModel is 创建门店 成功返回结果
type TmallServicecenterServicestoreCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_servicestore_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 方法调用结果
	Result *TmallServicecenterServicestoreCreateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterServicestoreCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterServicestoreCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterServicestoreCreateAPIResponse)
	},
}

// GetTmallServicecenterServicestoreCreateAPIResponse 从 sync.Pool 获取 TmallServicecenterServicestoreCreateAPIResponse
func GetTmallServicecenterServicestoreCreateAPIResponse() *TmallServicecenterServicestoreCreateAPIResponse {
	return poolTmallServicecenterServicestoreCreateAPIResponse.Get().(*TmallServicecenterServicestoreCreateAPIResponse)
}

// ReleaseTmallServicecenterServicestoreCreateAPIResponse 将 TmallServicecenterServicestoreCreateAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterServicestoreCreateAPIResponse(v *TmallServicecenterServicestoreCreateAPIResponse) {
	v.Reset()
	poolTmallServicecenterServicestoreCreateAPIResponse.Put(v)
}
