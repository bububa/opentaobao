package baichuan

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaBaichuanCtgContentGetAPIResponse 百川内容平台内容获取 API返回值
// alibaba.baichuan.ctg.content.get
//
// 百川内容平台内容获取
type AlibabaBaichuanCtgContentGetAPIResponse struct {
	model.CommonResponse
	AlibabaBaichuanCtgContentGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaBaichuanCtgContentGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaBaichuanCtgContentGetAPIResponseModel).Reset()
}

// AlibabaBaichuanCtgContentGetAPIResponseModel is 百川内容平台内容获取 成功返回结果
type AlibabaBaichuanCtgContentGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_baichuan_ctg_content_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	DataList []AlibabaBaichuanCtgContentGetData `json:"data_list,omitempty" xml:"data_list>alibaba_baichuan_ctg_content_get_data,omitempty"`
	// errorMessage
	ErrMessage string `json:"err_message,omitempty" xml:"err_message,omitempty"`
	// errorCode
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// hasNext
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaBaichuanCtgContentGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.DataList = m.DataList[:0]
	m.ErrMessage = ""
	m.ErrCode = ""
	m.HasNext = false
	m.IsSuccess = false
}

var poolAlibabaBaichuanCtgContentGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaBaichuanCtgContentGetAPIResponse)
	},
}

// GetAlibabaBaichuanCtgContentGetAPIResponse 从 sync.Pool 获取 AlibabaBaichuanCtgContentGetAPIResponse
func GetAlibabaBaichuanCtgContentGetAPIResponse() *AlibabaBaichuanCtgContentGetAPIResponse {
	return poolAlibabaBaichuanCtgContentGetAPIResponse.Get().(*AlibabaBaichuanCtgContentGetAPIResponse)
}

// ReleaseAlibabaBaichuanCtgContentGetAPIResponse 将 AlibabaBaichuanCtgContentGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaBaichuanCtgContentGetAPIResponse(v *AlibabaBaichuanCtgContentGetAPIResponse) {
	v.Reset()
	poolAlibabaBaichuanCtgContentGetAPIResponse.Put(v)
}
