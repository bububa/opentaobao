package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpAdzoneHorizontalFindpageAPIResponse 查看资源包列表 API返回值
// taobao.universalbp.adzone.horizontal.findpage
//
// 查看已存在的计划上设置的资源包列表
type TaobaoUniversalbpAdzoneHorizontalFindpageAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpAdzoneHorizontalFindpageAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUniversalbpAdzoneHorizontalFindpageAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUniversalbpAdzoneHorizontalFindpageAPIResponseModel).Reset()
}

// TaobaoUniversalbpAdzoneHorizontalFindpageAPIResponseModel is 查看资源包列表 成功返回结果
type TaobaoUniversalbpAdzoneHorizontalFindpageAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_adzone_horizontal_findpage_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpAdzoneHorizontalFindpageTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUniversalbpAdzoneHorizontalFindpageAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoUniversalbpAdzoneHorizontalFindpageAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpAdzoneHorizontalFindpageAPIResponse)
	},
}

// GetTaobaoUniversalbpAdzoneHorizontalFindpageAPIResponse 从 sync.Pool 获取 TaobaoUniversalbpAdzoneHorizontalFindpageAPIResponse
func GetTaobaoUniversalbpAdzoneHorizontalFindpageAPIResponse() *TaobaoUniversalbpAdzoneHorizontalFindpageAPIResponse {
	return poolTaobaoUniversalbpAdzoneHorizontalFindpageAPIResponse.Get().(*TaobaoUniversalbpAdzoneHorizontalFindpageAPIResponse)
}

// ReleaseTaobaoUniversalbpAdzoneHorizontalFindpageAPIResponse 将 TaobaoUniversalbpAdzoneHorizontalFindpageAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUniversalbpAdzoneHorizontalFindpageAPIResponse(v *TaobaoUniversalbpAdzoneHorizontalFindpageAPIResponse) {
	v.Reset()
	poolTaobaoUniversalbpAdzoneHorizontalFindpageAPIResponse.Put(v)
}
