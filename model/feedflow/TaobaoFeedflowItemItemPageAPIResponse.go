package feedflow

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemItemPageAPIResponse 信息流查看商品列表 API返回值
// taobao.feedflow.item.item.page
//
// 信息流查看商品列表
type TaobaoFeedflowItemItemPageAPIResponse struct {
	model.CommonResponse
	TaobaoFeedflowItemItemPageAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemItemPageAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFeedflowItemItemPageAPIResponseModel).Reset()
}

// TaobaoFeedflowItemItemPageAPIResponseModel is 信息流查看商品列表 成功返回结果
type TaobaoFeedflowItemItemPageAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_item_page_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果对象
	Result *TaobaoFeedflowItemItemPageResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemItemPageAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFeedflowItemItemPageAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemItemPageAPIResponse)
	},
}

// GetTaobaoFeedflowItemItemPageAPIResponse 从 sync.Pool 获取 TaobaoFeedflowItemItemPageAPIResponse
func GetTaobaoFeedflowItemItemPageAPIResponse() *TaobaoFeedflowItemItemPageAPIResponse {
	return poolTaobaoFeedflowItemItemPageAPIResponse.Get().(*TaobaoFeedflowItemItemPageAPIResponse)
}

// ReleaseTaobaoFeedflowItemItemPageAPIResponse 将 TaobaoFeedflowItemItemPageAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFeedflowItemItemPageAPIResponse(v *TaobaoFeedflowItemItemPageAPIResponse) {
	v.Reset()
	poolTaobaoFeedflowItemItemPageAPIResponse.Put(v)
}
