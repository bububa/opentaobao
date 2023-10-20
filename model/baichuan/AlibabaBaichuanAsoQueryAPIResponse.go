package baichuan

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaBaichuanAsoQueryAPIResponse 查询app在设备上的安装信息 API返回值
// alibaba.baichuan.aso.query
//
// 查询app在设备上的安装信息
type AlibabaBaichuanAsoQueryAPIResponse struct {
	model.CommonResponse
	AlibabaBaichuanAsoQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaBaichuanAsoQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaBaichuanAsoQueryAPIResponseModel).Reset()
}

// AlibabaBaichuanAsoQueryAPIResponseModel is 查询app在设备上的安装信息 成功返回结果
type AlibabaBaichuanAsoQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_baichuan_aso_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AsoQueryDeviceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaBaichuanAsoQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaBaichuanAsoQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaBaichuanAsoQueryAPIResponse)
	},
}

// GetAlibabaBaichuanAsoQueryAPIResponse 从 sync.Pool 获取 AlibabaBaichuanAsoQueryAPIResponse
func GetAlibabaBaichuanAsoQueryAPIResponse() *AlibabaBaichuanAsoQueryAPIResponse {
	return poolAlibabaBaichuanAsoQueryAPIResponse.Get().(*AlibabaBaichuanAsoQueryAPIResponse)
}

// ReleaseAlibabaBaichuanAsoQueryAPIResponse 将 AlibabaBaichuanAsoQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaBaichuanAsoQueryAPIResponse(v *AlibabaBaichuanAsoQueryAPIResponse) {
	v.Reset()
	poolAlibabaBaichuanAsoQueryAPIResponse.Put(v)
}
