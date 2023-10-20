package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpCreativePreaddAPIResponse 创建单品创意前置信息 API返回值
// taobao.universalbp.creative.preadd
//
// 用于关键词场景创建单品创意前使用
type TaobaoUniversalbpCreativePreaddAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpCreativePreaddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUniversalbpCreativePreaddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUniversalbpCreativePreaddAPIResponseModel).Reset()
}

// TaobaoUniversalbpCreativePreaddAPIResponseModel is 创建单品创意前置信息 成功返回结果
type TaobaoUniversalbpCreativePreaddAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_creative_preadd_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpCreativePreaddTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUniversalbpCreativePreaddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoUniversalbpCreativePreaddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpCreativePreaddAPIResponse)
	},
}

// GetTaobaoUniversalbpCreativePreaddAPIResponse 从 sync.Pool 获取 TaobaoUniversalbpCreativePreaddAPIResponse
func GetTaobaoUniversalbpCreativePreaddAPIResponse() *TaobaoUniversalbpCreativePreaddAPIResponse {
	return poolTaobaoUniversalbpCreativePreaddAPIResponse.Get().(*TaobaoUniversalbpCreativePreaddAPIResponse)
}

// ReleaseTaobaoUniversalbpCreativePreaddAPIResponse 将 TaobaoUniversalbpCreativePreaddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUniversalbpCreativePreaddAPIResponse(v *TaobaoUniversalbpCreativePreaddAPIResponse) {
	v.Reset()
	poolTaobaoUniversalbpCreativePreaddAPIResponse.Put(v)
}
