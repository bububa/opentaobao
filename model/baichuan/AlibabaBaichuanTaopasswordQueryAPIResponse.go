package baichuan

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaBaichuanTaopasswordQueryAPIResponse 查询解析淘口令 API返回值
// alibaba.baichuan.taopassword.query
//
// 查询，解析淘口令
type AlibabaBaichuanTaopasswordQueryAPIResponse struct {
	model.CommonResponse
	AlibabaBaichuanTaopasswordQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaBaichuanTaopasswordQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaBaichuanTaopasswordQueryAPIResponseModel).Reset()
}

// AlibabaBaichuanTaopasswordQueryAPIResponseModel is 查询解析淘口令 成功返回结果
type AlibabaBaichuanTaopasswordQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_baichuan_taopassword_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BcTaoPasswordResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaBaichuanTaopasswordQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaBaichuanTaopasswordQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaBaichuanTaopasswordQueryAPIResponse)
	},
}

// GetAlibabaBaichuanTaopasswordQueryAPIResponse 从 sync.Pool 获取 AlibabaBaichuanTaopasswordQueryAPIResponse
func GetAlibabaBaichuanTaopasswordQueryAPIResponse() *AlibabaBaichuanTaopasswordQueryAPIResponse {
	return poolAlibabaBaichuanTaopasswordQueryAPIResponse.Get().(*AlibabaBaichuanTaopasswordQueryAPIResponse)
}

// ReleaseAlibabaBaichuanTaopasswordQueryAPIResponse 将 AlibabaBaichuanTaopasswordQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaBaichuanTaopasswordQueryAPIResponse(v *AlibabaBaichuanTaopasswordQueryAPIResponse) {
	v.Reset()
	poolAlibabaBaichuanTaopasswordQueryAPIResponse.Put(v)
}
