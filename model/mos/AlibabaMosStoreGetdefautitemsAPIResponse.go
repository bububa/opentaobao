package mos

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosStoreGetdefautitemsAPIResponse 获取默认状态下商品列表 API返回值
// alibaba.mos.store.getdefautitems
//
// 获取默认状态下商品列表
type AlibabaMosStoreGetdefautitemsAPIResponse struct {
	model.CommonResponse
	AlibabaMosStoreGetdefautitemsAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMosStoreGetdefautitemsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMosStoreGetdefautitemsAPIResponseModel).Reset()
}

// AlibabaMosStoreGetdefautitemsAPIResponseModel is 获取默认状态下商品列表 成功返回结果
type AlibabaMosStoreGetdefautitemsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_store_getdefautitems_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaMosStoreGetdefautitemsResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMosStoreGetdefautitemsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMosStoreGetdefautitemsAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMosStoreGetdefautitemsAPIResponse)
	},
}

// GetAlibabaMosStoreGetdefautitemsAPIResponse 从 sync.Pool 获取 AlibabaMosStoreGetdefautitemsAPIResponse
func GetAlibabaMosStoreGetdefautitemsAPIResponse() *AlibabaMosStoreGetdefautitemsAPIResponse {
	return poolAlibabaMosStoreGetdefautitemsAPIResponse.Get().(*AlibabaMosStoreGetdefautitemsAPIResponse)
}

// ReleaseAlibabaMosStoreGetdefautitemsAPIResponse 将 AlibabaMosStoreGetdefautitemsAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMosStoreGetdefautitemsAPIResponse(v *AlibabaMosStoreGetdefautitemsAPIResponse) {
	v.Reset()
	poolAlibabaMosStoreGetdefautitemsAPIResponse.Put(v)
}
