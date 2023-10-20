package feedflow

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemAdgroupAddAPIResponse 信息流增加单元 API返回值
// taobao.feedflow.item.adgroup.add
//
// 信息流增加单元
type TaobaoFeedflowItemAdgroupAddAPIResponse struct {
	model.CommonResponse
	TaobaoFeedflowItemAdgroupAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemAdgroupAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFeedflowItemAdgroupAddAPIResponseModel).Reset()
}

// TaobaoFeedflowItemAdgroupAddAPIResponseModel is 信息流增加单元 成功返回结果
type TaobaoFeedflowItemAdgroupAddAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_adgroup_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *TaobaoFeedflowItemAdgroupAddResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemAdgroupAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFeedflowItemAdgroupAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemAdgroupAddAPIResponse)
	},
}

// GetTaobaoFeedflowItemAdgroupAddAPIResponse 从 sync.Pool 获取 TaobaoFeedflowItemAdgroupAddAPIResponse
func GetTaobaoFeedflowItemAdgroupAddAPIResponse() *TaobaoFeedflowItemAdgroupAddAPIResponse {
	return poolTaobaoFeedflowItemAdgroupAddAPIResponse.Get().(*TaobaoFeedflowItemAdgroupAddAPIResponse)
}

// ReleaseTaobaoFeedflowItemAdgroupAddAPIResponse 将 TaobaoFeedflowItemAdgroupAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFeedflowItemAdgroupAddAPIResponse(v *TaobaoFeedflowItemAdgroupAddAPIResponse) {
	v.Reset()
	poolTaobaoFeedflowItemAdgroupAddAPIResponse.Put(v)
}
