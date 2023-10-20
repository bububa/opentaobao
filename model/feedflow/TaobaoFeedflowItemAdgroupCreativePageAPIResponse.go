package feedflow

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemAdgroupCreativePageAPIResponse 信息流单元下查看创意 API返回值
// taobao.feedflow.item.adgroup.creative.page
//
// 信息流单元下查看创意
type TaobaoFeedflowItemAdgroupCreativePageAPIResponse struct {
	model.CommonResponse
	TaobaoFeedflowItemAdgroupCreativePageAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemAdgroupCreativePageAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFeedflowItemAdgroupCreativePageAPIResponseModel).Reset()
}

// TaobaoFeedflowItemAdgroupCreativePageAPIResponseModel is 信息流单元下查看创意 成功返回结果
type TaobaoFeedflowItemAdgroupCreativePageAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_adgroup_creative_page_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果返回消息
	Result *TaobaoFeedflowItemAdgroupCreativePageResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemAdgroupCreativePageAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFeedflowItemAdgroupCreativePageAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemAdgroupCreativePageAPIResponse)
	},
}

// GetTaobaoFeedflowItemAdgroupCreativePageAPIResponse 从 sync.Pool 获取 TaobaoFeedflowItemAdgroupCreativePageAPIResponse
func GetTaobaoFeedflowItemAdgroupCreativePageAPIResponse() *TaobaoFeedflowItemAdgroupCreativePageAPIResponse {
	return poolTaobaoFeedflowItemAdgroupCreativePageAPIResponse.Get().(*TaobaoFeedflowItemAdgroupCreativePageAPIResponse)
}

// ReleaseTaobaoFeedflowItemAdgroupCreativePageAPIResponse 将 TaobaoFeedflowItemAdgroupCreativePageAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFeedflowItemAdgroupCreativePageAPIResponse(v *TaobaoFeedflowItemAdgroupCreativePageAPIResponse) {
	v.Reset()
	poolTaobaoFeedflowItemAdgroupCreativePageAPIResponse.Put(v)
}
