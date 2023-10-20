package jst

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstMiniappCrowdCreateAPIResponse 小程序活动创建 API返回值
// taobao.jst.miniapp.crowd.create
//
// 小程序活动创建
type TaobaoJstMiniappCrowdCreateAPIResponse struct {
	model.CommonResponse
	TaobaoJstMiniappCrowdCreateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoJstMiniappCrowdCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoJstMiniappCrowdCreateAPIResponseModel).Reset()
}

// TaobaoJstMiniappCrowdCreateAPIResponseModel is 小程序活动创建 成功返回结果
type TaobaoJstMiniappCrowdCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"jst_miniapp_crowd_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 活动Code，活动的唯一标识
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// 状态码
	RCode int64 `json:"r_code,omitempty" xml:"r_code,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoJstMiniappCrowdCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
	m.RCode = 0
}

var poolTaobaoJstMiniappCrowdCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoJstMiniappCrowdCreateAPIResponse)
	},
}

// GetTaobaoJstMiniappCrowdCreateAPIResponse 从 sync.Pool 获取 TaobaoJstMiniappCrowdCreateAPIResponse
func GetTaobaoJstMiniappCrowdCreateAPIResponse() *TaobaoJstMiniappCrowdCreateAPIResponse {
	return poolTaobaoJstMiniappCrowdCreateAPIResponse.Get().(*TaobaoJstMiniappCrowdCreateAPIResponse)
}

// ReleaseTaobaoJstMiniappCrowdCreateAPIResponse 将 TaobaoJstMiniappCrowdCreateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoJstMiniappCrowdCreateAPIResponse(v *TaobaoJstMiniappCrowdCreateAPIResponse) {
	v.Reset()
	poolTaobaoJstMiniappCrowdCreateAPIResponse.Put(v)
}
