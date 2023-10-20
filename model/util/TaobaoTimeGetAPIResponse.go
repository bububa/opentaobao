package util

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTimeGetAPIResponse 获取淘宝系统当前时间 API返回值
// taobao.time.get
//
// 获取淘宝系统当前时间
type TaobaoTimeGetAPIResponse struct {
	model.CommonResponse
	TaobaoTimeGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTimeGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTimeGetAPIResponseModel).Reset()
}

// TaobaoTimeGetAPIResponseModel is 获取淘宝系统当前时间 成功返回结果
type TaobaoTimeGetAPIResponseModel struct {
	XMLName xml.Name `xml:"time_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 淘宝系统当前时间。格式:yyyy-MM-dd HH:mm:ss
	Time string `json:"time,omitempty" xml:"time,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTimeGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Time = ""
}

var poolTaobaoTimeGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTimeGetAPIResponse)
	},
}

// GetTaobaoTimeGetAPIResponse 从 sync.Pool 获取 TaobaoTimeGetAPIResponse
func GetTaobaoTimeGetAPIResponse() *TaobaoTimeGetAPIResponse {
	return poolTaobaoTimeGetAPIResponse.Get().(*TaobaoTimeGetAPIResponse)
}

// ReleaseTaobaoTimeGetAPIResponse 将 TaobaoTimeGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTimeGetAPIResponse(v *TaobaoTimeGetAPIResponse) {
	v.Reset()
	poolTaobaoTimeGetAPIResponse.Put(v)
}
