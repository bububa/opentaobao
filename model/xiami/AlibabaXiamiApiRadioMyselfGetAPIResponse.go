package xiami

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaXiamiApiRadioMyselfGetAPIResponse 我的电台 API返回值
// alibaba.xiami.api.radio.myself.get
//
// 我的电台
type AlibabaXiamiApiRadioMyselfGetAPIResponse struct {
	model.CommonResponse
	AlibabaXiamiApiRadioMyselfGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaXiamiApiRadioMyselfGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaXiamiApiRadioMyselfGetAPIResponseModel).Reset()
}

// AlibabaXiamiApiRadioMyselfGetAPIResponseModel is 我的电台 成功返回结果
type AlibabaXiamiApiRadioMyselfGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_xiami_api_radio_myself_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 歌曲列表
	Data []Song `json:"data,omitempty" xml:"data>song,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaXiamiApiRadioMyselfGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = m.Data[:0]
}

var poolAlibabaXiamiApiRadioMyselfGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaXiamiApiRadioMyselfGetAPIResponse)
	},
}

// GetAlibabaXiamiApiRadioMyselfGetAPIResponse 从 sync.Pool 获取 AlibabaXiamiApiRadioMyselfGetAPIResponse
func GetAlibabaXiamiApiRadioMyselfGetAPIResponse() *AlibabaXiamiApiRadioMyselfGetAPIResponse {
	return poolAlibabaXiamiApiRadioMyselfGetAPIResponse.Get().(*AlibabaXiamiApiRadioMyselfGetAPIResponse)
}

// ReleaseAlibabaXiamiApiRadioMyselfGetAPIResponse 将 AlibabaXiamiApiRadioMyselfGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaXiamiApiRadioMyselfGetAPIResponse(v *AlibabaXiamiApiRadioMyselfGetAPIResponse) {
	v.Reset()
	poolAlibabaXiamiApiRadioMyselfGetAPIResponse.Put(v)
}
