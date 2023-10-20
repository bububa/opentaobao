package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCfdaXtptAppAcceptInfoAPIResponse 协同平台数据下行接口 API返回值
// alibaba.cfda.xtpt.app.accept.info
//
// 协同平台数据下行接口
type AlibabaCfdaXtptAppAcceptInfoAPIResponse struct {
	model.CommonResponse
	AlibabaCfdaXtptAppAcceptInfoAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCfdaXtptAppAcceptInfoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCfdaXtptAppAcceptInfoAPIResponseModel).Reset()
}

// AlibabaCfdaXtptAppAcceptInfoAPIResponseModel is 协同平台数据下行接口 成功返回结果
type AlibabaCfdaXtptAppAcceptInfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_cfda_xtpt_app_accept_info_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaCfdaXtptAppAcceptInfoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCfdaXtptAppAcceptInfoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCfdaXtptAppAcceptInfoAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCfdaXtptAppAcceptInfoAPIResponse)
	},
}

// GetAlibabaCfdaXtptAppAcceptInfoAPIResponse 从 sync.Pool 获取 AlibabaCfdaXtptAppAcceptInfoAPIResponse
func GetAlibabaCfdaXtptAppAcceptInfoAPIResponse() *AlibabaCfdaXtptAppAcceptInfoAPIResponse {
	return poolAlibabaCfdaXtptAppAcceptInfoAPIResponse.Get().(*AlibabaCfdaXtptAppAcceptInfoAPIResponse)
}

// ReleaseAlibabaCfdaXtptAppAcceptInfoAPIResponse 将 AlibabaCfdaXtptAppAcceptInfoAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCfdaXtptAppAcceptInfoAPIResponse(v *AlibabaCfdaXtptAppAcceptInfoAPIResponse) {
	v.Reset()
	poolAlibabaCfdaXtptAppAcceptInfoAPIResponse.Put(v)
}
