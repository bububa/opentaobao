package alicom

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWttUserRegioninfoByipGetAPIResponse 根据ip获取省市信息 API返回值
// alibaba.wtt.user.regioninfo.byip.get
//
// 通过ip获取省市信息
type AlibabaWttUserRegioninfoByipGetAPIResponse struct {
	model.CommonResponse
	AlibabaWttUserRegioninfoByipGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWttUserRegioninfoByipGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWttUserRegioninfoByipGetAPIResponseModel).Reset()
}

// AlibabaWttUserRegioninfoByipGetAPIResponseModel is 根据ip获取省市信息 成功返回结果
type AlibabaWttUserRegioninfoByipGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wtt_user_regioninfo_byip_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 地址信息
	Model *RegionInfo `json:"model,omitempty" xml:"model,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWttUserRegioninfoByipGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Model = nil
}

var poolAlibabaWttUserRegioninfoByipGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWttUserRegioninfoByipGetAPIResponse)
	},
}

// GetAlibabaWttUserRegioninfoByipGetAPIResponse 从 sync.Pool 获取 AlibabaWttUserRegioninfoByipGetAPIResponse
func GetAlibabaWttUserRegioninfoByipGetAPIResponse() *AlibabaWttUserRegioninfoByipGetAPIResponse {
	return poolAlibabaWttUserRegioninfoByipGetAPIResponse.Get().(*AlibabaWttUserRegioninfoByipGetAPIResponse)
}

// ReleaseAlibabaWttUserRegioninfoByipGetAPIResponse 将 AlibabaWttUserRegioninfoByipGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWttUserRegioninfoByipGetAPIResponse(v *AlibabaWttUserRegioninfoByipGetAPIResponse) {
	v.Reset()
	poolAlibabaWttUserRegioninfoByipGetAPIResponse.Put(v)
}
