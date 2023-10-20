package feedflow

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemAdgroupPageAPIResponse 查询单元列表 API返回值
// taobao.feedflow.item.adgroup.page
//
// 通过计划id查询单元信息
type TaobaoFeedflowItemAdgroupPageAPIResponse struct {
	model.CommonResponse
	TaobaoFeedflowItemAdgroupPageAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemAdgroupPageAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFeedflowItemAdgroupPageAPIResponseModel).Reset()
}

// TaobaoFeedflowItemAdgroupPageAPIResponseModel is 查询单元列表 成功返回结果
type TaobaoFeedflowItemAdgroupPageAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_adgroup_page_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回数据
	Result *TaobaoFeedflowItemAdgroupPageResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemAdgroupPageAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFeedflowItemAdgroupPageAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemAdgroupPageAPIResponse)
	},
}

// GetTaobaoFeedflowItemAdgroupPageAPIResponse 从 sync.Pool 获取 TaobaoFeedflowItemAdgroupPageAPIResponse
func GetTaobaoFeedflowItemAdgroupPageAPIResponse() *TaobaoFeedflowItemAdgroupPageAPIResponse {
	return poolTaobaoFeedflowItemAdgroupPageAPIResponse.Get().(*TaobaoFeedflowItemAdgroupPageAPIResponse)
}

// ReleaseTaobaoFeedflowItemAdgroupPageAPIResponse 将 TaobaoFeedflowItemAdgroupPageAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFeedflowItemAdgroupPageAPIResponse(v *TaobaoFeedflowItemAdgroupPageAPIResponse) {
	v.Reset()
	poolTaobaoFeedflowItemAdgroupPageAPIResponse.Put(v)
}
