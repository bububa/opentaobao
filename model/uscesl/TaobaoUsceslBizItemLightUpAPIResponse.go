package uscesl

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUsceslBizItemLightUpAPIResponse 商品条码亮灯API API返回值
// taobao.uscesl.biz.item.light.up
//
// 亮灯API
type TaobaoUsceslBizItemLightUpAPIResponse struct {
	model.CommonResponse
	TaobaoUsceslBizItemLightUpAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUsceslBizItemLightUpAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUsceslBizItemLightUpAPIResponseModel).Reset()
}

// TaobaoUsceslBizItemLightUpAPIResponseModel is 商品条码亮灯API 成功返回结果
type TaobaoUsceslBizItemLightUpAPIResponseModel struct {
	XMLName xml.Name `xml:"uscesl_biz_item_light_up_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoUsceslBizItemLightUpResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUsceslBizItemLightUpAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoUsceslBizItemLightUpAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUsceslBizItemLightUpAPIResponse)
	},
}

// GetTaobaoUsceslBizItemLightUpAPIResponse 从 sync.Pool 获取 TaobaoUsceslBizItemLightUpAPIResponse
func GetTaobaoUsceslBizItemLightUpAPIResponse() *TaobaoUsceslBizItemLightUpAPIResponse {
	return poolTaobaoUsceslBizItemLightUpAPIResponse.Get().(*TaobaoUsceslBizItemLightUpAPIResponse)
}

// ReleaseTaobaoUsceslBizItemLightUpAPIResponse 将 TaobaoUsceslBizItemLightUpAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUsceslBizItemLightUpAPIResponse(v *TaobaoUsceslBizItemLightUpAPIResponse) {
	v.Reset()
	poolTaobaoUsceslBizItemLightUpAPIResponse.Put(v)
}
