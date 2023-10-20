package uscesl

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUsceslBizLightUpAPIResponse 价签LED等点亮 API返回值
// taobao.uscesl.biz.light.up
//
// 价签LED等点亮
type TaobaoUsceslBizLightUpAPIResponse struct {
	model.CommonResponse
	TaobaoUsceslBizLightUpAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUsceslBizLightUpAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUsceslBizLightUpAPIResponseModel).Reset()
}

// TaobaoUsceslBizLightUpAPIResponseModel is 价签LED等点亮 成功返回结果
type TaobaoUsceslBizLightUpAPIResponseModel struct {
	XMLName xml.Name `xml:"uscesl_biz_light_up_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaoUsceslBizLightUpResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUsceslBizLightUpAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoUsceslBizLightUpAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUsceslBizLightUpAPIResponse)
	},
}

// GetTaobaoUsceslBizLightUpAPIResponse 从 sync.Pool 获取 TaobaoUsceslBizLightUpAPIResponse
func GetTaobaoUsceslBizLightUpAPIResponse() *TaobaoUsceslBizLightUpAPIResponse {
	return poolTaobaoUsceslBizLightUpAPIResponse.Get().(*TaobaoUsceslBizLightUpAPIResponse)
}

// ReleaseTaobaoUsceslBizLightUpAPIResponse 将 TaobaoUsceslBizLightUpAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUsceslBizLightUpAPIResponse(v *TaobaoUsceslBizLightUpAPIResponse) {
	v.Reset()
	poolTaobaoUsceslBizLightUpAPIResponse.Put(v)
}
