package traderate

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFliggyWrateGetmixratelistAPIResponse 飞猪通用评价接口 API返回值
// taobao.fliggy.wrate.getmixratelist
//
// 飞猪评价通用接口
type TaobaoFliggyWrateGetmixratelistAPIResponse struct {
	model.CommonResponse
	TaobaoFliggyWrateGetmixratelistAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFliggyWrateGetmixratelistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFliggyWrateGetmixratelistAPIResponseModel).Reset()
}

// TaobaoFliggyWrateGetmixratelistAPIResponseModel is 飞猪通用评价接口 成功返回结果
type TaobaoFliggyWrateGetmixratelistAPIResponseModel struct {
	XMLName xml.Name `xml:"fliggy_wrate_getmixratelist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoFliggyWrateGetmixratelistResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFliggyWrateGetmixratelistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFliggyWrateGetmixratelistAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFliggyWrateGetmixratelistAPIResponse)
	},
}

// GetTaobaoFliggyWrateGetmixratelistAPIResponse 从 sync.Pool 获取 TaobaoFliggyWrateGetmixratelistAPIResponse
func GetTaobaoFliggyWrateGetmixratelistAPIResponse() *TaobaoFliggyWrateGetmixratelistAPIResponse {
	return poolTaobaoFliggyWrateGetmixratelistAPIResponse.Get().(*TaobaoFliggyWrateGetmixratelistAPIResponse)
}

// ReleaseTaobaoFliggyWrateGetmixratelistAPIResponse 将 TaobaoFliggyWrateGetmixratelistAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFliggyWrateGetmixratelistAPIResponse(v *TaobaoFliggyWrateGetmixratelistAPIResponse) {
	v.Reset()
	poolTaobaoFliggyWrateGetmixratelistAPIResponse.Put(v)
}
