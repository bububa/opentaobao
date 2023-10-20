package jstinteractive

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstInteractiveAssetsConfigAPIResponse 任务素材配置接口 API返回值
// taobao.jst.interactive.assets.config
//
// 任务素材配置接口
type TaobaoJstInteractiveAssetsConfigAPIResponse struct {
	model.CommonResponse
	TaobaoJstInteractiveAssetsConfigAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoJstInteractiveAssetsConfigAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoJstInteractiveAssetsConfigAPIResponseModel).Reset()
}

// TaobaoJstInteractiveAssetsConfigAPIResponseModel is 任务素材配置接口 成功返回结果
type TaobaoJstInteractiveAssetsConfigAPIResponseModel struct {
	XMLName xml.Name `xml:"jst_interactive_assets_config_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoJstInteractiveAssetsConfigAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoJstInteractiveAssetsConfigAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoJstInteractiveAssetsConfigAPIResponse)
	},
}

// GetTaobaoJstInteractiveAssetsConfigAPIResponse 从 sync.Pool 获取 TaobaoJstInteractiveAssetsConfigAPIResponse
func GetTaobaoJstInteractiveAssetsConfigAPIResponse() *TaobaoJstInteractiveAssetsConfigAPIResponse {
	return poolTaobaoJstInteractiveAssetsConfigAPIResponse.Get().(*TaobaoJstInteractiveAssetsConfigAPIResponse)
}

// ReleaseTaobaoJstInteractiveAssetsConfigAPIResponse 将 TaobaoJstInteractiveAssetsConfigAPIResponse 保存到 sync.Pool
func ReleaseTaobaoJstInteractiveAssetsConfigAPIResponse(v *TaobaoJstInteractiveAssetsConfigAPIResponse) {
	v.Reset()
	poolTaobaoJstInteractiveAssetsConfigAPIResponse.Put(v)
}
