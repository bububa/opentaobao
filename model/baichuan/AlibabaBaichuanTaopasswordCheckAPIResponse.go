package baichuan

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaBaichuanTaopasswordCheckAPIResponse 淘口令检查 API返回值
// alibaba.baichuan.taopassword.check
//
// 检查当前文本是否为淘口令
type AlibabaBaichuanTaopasswordCheckAPIResponse struct {
	model.CommonResponse
	AlibabaBaichuanTaopasswordCheckAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaBaichuanTaopasswordCheckAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaBaichuanTaopasswordCheckAPIResponseModel).Reset()
}

// AlibabaBaichuanTaopasswordCheckAPIResponseModel is 淘口令检查 成功返回结果
type AlibabaBaichuanTaopasswordCheckAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_baichuan_taopassword_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaBaichuanTaopasswordCheckResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaBaichuanTaopasswordCheckAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaBaichuanTaopasswordCheckAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaBaichuanTaopasswordCheckAPIResponse)
	},
}

// GetAlibabaBaichuanTaopasswordCheckAPIResponse 从 sync.Pool 获取 AlibabaBaichuanTaopasswordCheckAPIResponse
func GetAlibabaBaichuanTaopasswordCheckAPIResponse() *AlibabaBaichuanTaopasswordCheckAPIResponse {
	return poolAlibabaBaichuanTaopasswordCheckAPIResponse.Get().(*AlibabaBaichuanTaopasswordCheckAPIResponse)
}

// ReleaseAlibabaBaichuanTaopasswordCheckAPIResponse 将 AlibabaBaichuanTaopasswordCheckAPIResponse 保存到 sync.Pool
func ReleaseAlibabaBaichuanTaopasswordCheckAPIResponse(v *AlibabaBaichuanTaopasswordCheckAPIResponse) {
	v.Reset()
	poolAlibabaBaichuanTaopasswordCheckAPIResponse.Put(v)
}
