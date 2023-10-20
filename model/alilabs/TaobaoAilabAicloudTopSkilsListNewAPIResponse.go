package alilabs

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopSkilsListNewAPIResponse 获取产品下挂载的技能列表 API返回值
// taobao.ailab.aicloud.top.skils.list.new
//
// 星空平台提供的获取产品下挂载的技能列表新接口
type TaobaoAilabAicloudTopSkilsListNewAPIResponse struct {
	model.CommonResponse
	TaobaoAilabAicloudTopSkilsListNewAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopSkilsListNewAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAilabAicloudTopSkilsListNewAPIResponseModel).Reset()
}

// TaobaoAilabAicloudTopSkilsListNewAPIResponseModel is 获取产品下挂载的技能列表 成功返回结果
type TaobaoAilabAicloudTopSkilsListNewAPIResponseModel struct {
	XMLName xml.Name `xml:"ailab_aicloud_top_skils_list_new_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回包装类
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAilabAicloudTopSkilsListNewAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoAilabAicloudTopSkilsListNewAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopSkilsListNewAPIResponse)
	},
}

// GetTaobaoAilabAicloudTopSkilsListNewAPIResponse 从 sync.Pool 获取 TaobaoAilabAicloudTopSkilsListNewAPIResponse
func GetTaobaoAilabAicloudTopSkilsListNewAPIResponse() *TaobaoAilabAicloudTopSkilsListNewAPIResponse {
	return poolTaobaoAilabAicloudTopSkilsListNewAPIResponse.Get().(*TaobaoAilabAicloudTopSkilsListNewAPIResponse)
}

// ReleaseTaobaoAilabAicloudTopSkilsListNewAPIResponse 将 TaobaoAilabAicloudTopSkilsListNewAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAilabAicloudTopSkilsListNewAPIResponse(v *TaobaoAilabAicloudTopSkilsListNewAPIResponse) {
	v.Reset()
	poolTaobaoAilabAicloudTopSkilsListNewAPIResponse.Put(v)
}
