package uscesl

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUsceslBizApSearchAPIResponse AP列表查询 API返回值
// taobao.uscesl.biz.ap.search
//
// 查询当前门店下登记的AP列表
type TaobaoUsceslBizApSearchAPIResponse struct {
	model.CommonResponse
	TaobaoUsceslBizApSearchAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUsceslBizApSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUsceslBizApSearchAPIResponseModel).Reset()
}

// TaobaoUsceslBizApSearchAPIResponseModel is AP列表查询 成功返回结果
type TaobaoUsceslBizApSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"uscesl_biz_ap_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 成功与否看result.success，返回true或者false
	Result *TaobaoUsceslBizApSearchResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUsceslBizApSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoUsceslBizApSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUsceslBizApSearchAPIResponse)
	},
}

// GetTaobaoUsceslBizApSearchAPIResponse 从 sync.Pool 获取 TaobaoUsceslBizApSearchAPIResponse
func GetTaobaoUsceslBizApSearchAPIResponse() *TaobaoUsceslBizApSearchAPIResponse {
	return poolTaobaoUsceslBizApSearchAPIResponse.Get().(*TaobaoUsceslBizApSearchAPIResponse)
}

// ReleaseTaobaoUsceslBizApSearchAPIResponse 将 TaobaoUsceslBizApSearchAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUsceslBizApSearchAPIResponse(v *TaobaoUsceslBizApSearchAPIResponse) {
	v.Reset()
	poolTaobaoUsceslBizApSearchAPIResponse.Put(v)
}
