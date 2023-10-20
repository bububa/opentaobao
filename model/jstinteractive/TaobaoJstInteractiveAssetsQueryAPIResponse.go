package jstinteractive

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstInteractiveAssetsQueryAPIResponse 查询可配置任务素材接口 API返回值
// taobao.jst.interactive.assets.query
//
// 查询可配置任务素材列表，用以配置任务素材
type TaobaoJstInteractiveAssetsQueryAPIResponse struct {
	model.CommonResponse
	TaobaoJstInteractiveAssetsQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoJstInteractiveAssetsQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoJstInteractiveAssetsQueryAPIResponseModel).Reset()
}

// TaobaoJstInteractiveAssetsQueryAPIResponseModel is 查询可配置任务素材接口 成功返回结果
type TaobaoJstInteractiveAssetsQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"jst_interactive_assets_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 可配置任务素材列表
	AssetsList []AssetsConfig `json:"assets_list,omitempty" xml:"assets_list>assets_config,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoJstInteractiveAssetsQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.AssetsList = m.AssetsList[:0]
}

var poolTaobaoJstInteractiveAssetsQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoJstInteractiveAssetsQueryAPIResponse)
	},
}

// GetTaobaoJstInteractiveAssetsQueryAPIResponse 从 sync.Pool 获取 TaobaoJstInteractiveAssetsQueryAPIResponse
func GetTaobaoJstInteractiveAssetsQueryAPIResponse() *TaobaoJstInteractiveAssetsQueryAPIResponse {
	return poolTaobaoJstInteractiveAssetsQueryAPIResponse.Get().(*TaobaoJstInteractiveAssetsQueryAPIResponse)
}

// ReleaseTaobaoJstInteractiveAssetsQueryAPIResponse 将 TaobaoJstInteractiveAssetsQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoJstInteractiveAssetsQueryAPIResponse(v *TaobaoJstInteractiveAssetsQueryAPIResponse) {
	v.Reset()
	poolTaobaoJstInteractiveAssetsQueryAPIResponse.Put(v)
}
