package btrip

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripBtripOpenplatformAddressGetAPIResponse 【商旅】开放平台对外页面跳转 API返回值
// alitrip.btrip.openplatform.address.get
//
// 获取类目预定页跳转地址
type AlitripBtripOpenplatformAddressGetAPIResponse struct {
	model.CommonResponse
	AlitripBtripOpenplatformAddressGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripBtripOpenplatformAddressGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripBtripOpenplatformAddressGetAPIResponseModel).Reset()
}

// AlitripBtripOpenplatformAddressGetAPIResponseModel is 【商旅】开放平台对外页面跳转 成功返回结果
type AlitripBtripOpenplatformAddressGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_openplatform_address_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripBtripOpenplatformAddressGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripBtripOpenplatformAddressGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripBtripOpenplatformAddressGetAPIResponse)
	},
}

// GetAlitripBtripOpenplatformAddressGetAPIResponse 从 sync.Pool 获取 AlitripBtripOpenplatformAddressGetAPIResponse
func GetAlitripBtripOpenplatformAddressGetAPIResponse() *AlitripBtripOpenplatformAddressGetAPIResponse {
	return poolAlitripBtripOpenplatformAddressGetAPIResponse.Get().(*AlitripBtripOpenplatformAddressGetAPIResponse)
}

// ReleaseAlitripBtripOpenplatformAddressGetAPIResponse 将 AlitripBtripOpenplatformAddressGetAPIResponse 保存到 sync.Pool
func ReleaseAlitripBtripOpenplatformAddressGetAPIResponse(v *AlitripBtripOpenplatformAddressGetAPIResponse) {
	v.Reset()
	poolAlitripBtripOpenplatformAddressGetAPIResponse.Put(v)
}
