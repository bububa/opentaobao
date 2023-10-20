package uscesl

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUsceslIteminfoPutAPIResponse 电子价签显示用商品信息写入 API返回值
// taobao.uscesl.iteminfo.put
//
// 用于电子价签上显示的商品信息的写入，包含价格及促销信息
type TaobaoUsceslIteminfoPutAPIResponse struct {
	model.CommonResponse
	TaobaoUsceslIteminfoPutAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUsceslIteminfoPutAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUsceslIteminfoPutAPIResponseModel).Reset()
}

// TaobaoUsceslIteminfoPutAPIResponseModel is 电子价签显示用商品信息写入 成功返回结果
type TaobaoUsceslIteminfoPutAPIResponseModel struct {
	XMLName xml.Name `xml:"uscesl_iteminfo_put_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// returnCode
	ReturnCode int64 `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// target
	Target bool `json:"target,omitempty" xml:"target,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUsceslIteminfoPutAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.ReturnCode = 0
	m.Target = false
}

var poolTaobaoUsceslIteminfoPutAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUsceslIteminfoPutAPIResponse)
	},
}

// GetTaobaoUsceslIteminfoPutAPIResponse 从 sync.Pool 获取 TaobaoUsceslIteminfoPutAPIResponse
func GetTaobaoUsceslIteminfoPutAPIResponse() *TaobaoUsceslIteminfoPutAPIResponse {
	return poolTaobaoUsceslIteminfoPutAPIResponse.Get().(*TaobaoUsceslIteminfoPutAPIResponse)
}

// ReleaseTaobaoUsceslIteminfoPutAPIResponse 将 TaobaoUsceslIteminfoPutAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUsceslIteminfoPutAPIResponse(v *TaobaoUsceslIteminfoPutAPIResponse) {
	v.Reset()
	poolTaobaoUsceslIteminfoPutAPIResponse.Put(v)
}
