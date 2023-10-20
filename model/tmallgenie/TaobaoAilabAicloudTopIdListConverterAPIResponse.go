package tmallgenie

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopIdListConverterAPIResponse 将淘宝openId或者设备id/用户id转换为openId API返回值
// taobao.ailab.aicloud.top.id.list.converter
//
// 将淘宝openId或者设备id/用户id转换为openId
type TaobaoAilabAicloudTopIdListConverterAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopIdListConverterAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopIdListConverterAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAilabAicloudTopIdListConverterAPIResponseModel).Reset()
}

// TaobaoAilabAicloudTopIdListConverterAPIResponseModel is 将淘宝openId或者设备id/用户id转换为openId 成功返回结果
type TaobaoAilabAicloudTopIdListConverterAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_id_list_converter_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回体
	Result *TaobaoAilabAicloudTopIdListConverterResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopIdListConverterAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAilabAicloudTopIdListConverterAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopIdListConverterAPIResponse)
	},
}

// GetTaobaoAilabAicloudTopIdListConverterAPIResponse 从 sync.Pool 获取 TaobaoAilabAicloudTopIdListConverterAPIResponse
func GetTaobaoAilabAicloudTopIdListConverterAPIResponse() *TaobaoAilabAicloudTopIdListConverterAPIResponse {
	return poolTaobaoAilabAicloudTopIdListConverterAPIResponse.Get().(*TaobaoAilabAicloudTopIdListConverterAPIResponse)
}

// ReleaseTaobaoAilabAicloudTopIdListConverterAPIResponse 将 TaobaoAilabAicloudTopIdListConverterAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAilabAicloudTopIdListConverterAPIResponse(v *TaobaoAilabAicloudTopIdListConverterAPIResponse) {
	v.Reset()
	poolTaobaoAilabAicloudTopIdListConverterAPIResponse.Put(v)
}
