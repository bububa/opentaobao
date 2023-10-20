package tmallgenie

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopSkilsListAPIResponse 获取硬件平台设备下挂载的技能列表 API返回值
// taobao.ailab.aicloud.top.skils.list
//
// 提供给在硬件平台接入设备的技能列表
type TaobaoAilabAicloudTopSkilsListAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopSkilsListAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopSkilsListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAilabAicloudTopSkilsListAPIResponseModel).Reset()
}

// TaobaoAilabAicloudTopSkilsListAPIResponseModel is 获取硬件平台设备下挂载的技能列表 成功返回结果
type TaobaoAilabAicloudTopSkilsListAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_skils_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopSkilsListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAilabAicloudTopSkilsListAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopSkilsListAPIResponse)
	},
}

// GetTaobaoAilabAicloudTopSkilsListAPIResponse 从 sync.Pool 获取 TaobaoAilabAicloudTopSkilsListAPIResponse
func GetTaobaoAilabAicloudTopSkilsListAPIResponse() *TaobaoAilabAicloudTopSkilsListAPIResponse {
	return poolTaobaoAilabAicloudTopSkilsListAPIResponse.Get().(*TaobaoAilabAicloudTopSkilsListAPIResponse)
}

// ReleaseTaobaoAilabAicloudTopSkilsListAPIResponse 将 TaobaoAilabAicloudTopSkilsListAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAilabAicloudTopSkilsListAPIResponse(v *TaobaoAilabAicloudTopSkilsListAPIResponse) {
	v.Reset()
	poolTaobaoAilabAicloudTopSkilsListAPIResponse.Put(v)
}
