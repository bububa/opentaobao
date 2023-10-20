package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpCreativeHorizontalFindpageAPIResponse 横向管理创意分页查询 API返回值
// taobao.universalbp.creative.horizontal.findpage
//
// 横向管理创意分页查询
type TaobaoUniversalbpCreativeHorizontalFindpageAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpCreativeHorizontalFindpageAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUniversalbpCreativeHorizontalFindpageAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUniversalbpCreativeHorizontalFindpageAPIResponseModel).Reset()
}

// TaobaoUniversalbpCreativeHorizontalFindpageAPIResponseModel is 横向管理创意分页查询 成功返回结果
type TaobaoUniversalbpCreativeHorizontalFindpageAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_creative_horizontal_findpage_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpCreativeHorizontalFindpageTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUniversalbpCreativeHorizontalFindpageAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoUniversalbpCreativeHorizontalFindpageAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpCreativeHorizontalFindpageAPIResponse)
	},
}

// GetTaobaoUniversalbpCreativeHorizontalFindpageAPIResponse 从 sync.Pool 获取 TaobaoUniversalbpCreativeHorizontalFindpageAPIResponse
func GetTaobaoUniversalbpCreativeHorizontalFindpageAPIResponse() *TaobaoUniversalbpCreativeHorizontalFindpageAPIResponse {
	return poolTaobaoUniversalbpCreativeHorizontalFindpageAPIResponse.Get().(*TaobaoUniversalbpCreativeHorizontalFindpageAPIResponse)
}

// ReleaseTaobaoUniversalbpCreativeHorizontalFindpageAPIResponse 将 TaobaoUniversalbpCreativeHorizontalFindpageAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUniversalbpCreativeHorizontalFindpageAPIResponse(v *TaobaoUniversalbpCreativeHorizontalFindpageAPIResponse) {
	v.Reset()
	poolTaobaoUniversalbpCreativeHorizontalFindpageAPIResponse.Put(v)
}
