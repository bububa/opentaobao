package baichuan

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaBaichuanTaopasswordConfigAPIResponse 淘口令配置数据 API返回值
// alibaba.baichuan.taopassword.config
//
// 百川淘口令规则配置接口
type AlibabaBaichuanTaopasswordConfigAPIResponse struct {
	model.CommonResponse
	AlibabaBaichuanTaopasswordConfigAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaBaichuanTaopasswordConfigAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaBaichuanTaopasswordConfigAPIResponseModel).Reset()
}

// AlibabaBaichuanTaopasswordConfigAPIResponseModel is 淘口令配置数据 成功返回结果
type AlibabaBaichuanTaopasswordConfigAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_baichuan_taopassword_config_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ShareResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaBaichuanTaopasswordConfigAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaBaichuanTaopasswordConfigAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaBaichuanTaopasswordConfigAPIResponse)
	},
}

// GetAlibabaBaichuanTaopasswordConfigAPIResponse 从 sync.Pool 获取 AlibabaBaichuanTaopasswordConfigAPIResponse
func GetAlibabaBaichuanTaopasswordConfigAPIResponse() *AlibabaBaichuanTaopasswordConfigAPIResponse {
	return poolAlibabaBaichuanTaopasswordConfigAPIResponse.Get().(*AlibabaBaichuanTaopasswordConfigAPIResponse)
}

// ReleaseAlibabaBaichuanTaopasswordConfigAPIResponse 将 AlibabaBaichuanTaopasswordConfigAPIResponse 保存到 sync.Pool
func ReleaseAlibabaBaichuanTaopasswordConfigAPIResponse(v *AlibabaBaichuanTaopasswordConfigAPIResponse) {
	v.Reset()
	poolAlibabaBaichuanTaopasswordConfigAPIResponse.Put(v)
}
