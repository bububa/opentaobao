package mos

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosStoreGetstorelistAPIResponse 根据屏编号获取专柜集 API返回值
// alibaba.mos.store.getstorelist
//
// 根据屏编号获取专柜集
type AlibabaMosStoreGetstorelistAPIResponse struct {
	model.CommonResponse
	AlibabaMosStoreGetstorelistAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMosStoreGetstorelistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMosStoreGetstorelistAPIResponseModel).Reset()
}

// AlibabaMosStoreGetstorelistAPIResponseModel is 根据屏编号获取专柜集 成功返回结果
type AlibabaMosStoreGetstorelistAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_store_getstorelist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaMosStoreGetstorelistResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMosStoreGetstorelistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMosStoreGetstorelistAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMosStoreGetstorelistAPIResponse)
	},
}

// GetAlibabaMosStoreGetstorelistAPIResponse 从 sync.Pool 获取 AlibabaMosStoreGetstorelistAPIResponse
func GetAlibabaMosStoreGetstorelistAPIResponse() *AlibabaMosStoreGetstorelistAPIResponse {
	return poolAlibabaMosStoreGetstorelistAPIResponse.Get().(*AlibabaMosStoreGetstorelistAPIResponse)
}

// ReleaseAlibabaMosStoreGetstorelistAPIResponse 将 AlibabaMosStoreGetstorelistAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMosStoreGetstorelistAPIResponse(v *AlibabaMosStoreGetstorelistAPIResponse) {
	v.Reset()
	poolAlibabaMosStoreGetstorelistAPIResponse.Put(v)
}
