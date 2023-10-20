package feedflow

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemOptionPageAPIResponse 分页查询定向标签列表 API返回值
// taobao.feedflow.item.option.page
//
// 分页查询定向标签列表
type TaobaoFeedflowItemOptionPageAPIResponse struct {
	model.CommonResponse
	TaobaoFeedflowItemOptionPageAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemOptionPageAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFeedflowItemOptionPageAPIResponseModel).Reset()
}

// TaobaoFeedflowItemOptionPageAPIResponseModel is 分页查询定向标签列表 成功返回结果
type TaobaoFeedflowItemOptionPageAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_option_page_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果对象
	Result *TaobaoFeedflowItemOptionPageResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemOptionPageAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFeedflowItemOptionPageAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemOptionPageAPIResponse)
	},
}

// GetTaobaoFeedflowItemOptionPageAPIResponse 从 sync.Pool 获取 TaobaoFeedflowItemOptionPageAPIResponse
func GetTaobaoFeedflowItemOptionPageAPIResponse() *TaobaoFeedflowItemOptionPageAPIResponse {
	return poolTaobaoFeedflowItemOptionPageAPIResponse.Get().(*TaobaoFeedflowItemOptionPageAPIResponse)
}

// ReleaseTaobaoFeedflowItemOptionPageAPIResponse 将 TaobaoFeedflowItemOptionPageAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFeedflowItemOptionPageAPIResponse(v *TaobaoFeedflowItemOptionPageAPIResponse) {
	v.Reset()
	poolTaobaoFeedflowItemOptionPageAPIResponse.Put(v)
}
