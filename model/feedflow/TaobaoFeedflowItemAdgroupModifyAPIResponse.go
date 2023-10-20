package feedflow

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemAdgroupModifyAPIResponse 信息流单元修改 API返回值
// taobao.feedflow.item.adgroup.modify
//
// 信息流单元修改
type TaobaoFeedflowItemAdgroupModifyAPIResponse struct {
	model.CommonResponse
	TaobaoFeedflowItemAdgroupModifyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemAdgroupModifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFeedflowItemAdgroupModifyAPIResponseModel).Reset()
}

// TaobaoFeedflowItemAdgroupModifyAPIResponseModel is 信息流单元修改 成功返回结果
type TaobaoFeedflowItemAdgroupModifyAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_adgroup_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *TaobaoFeedflowItemAdgroupModifyResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemAdgroupModifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFeedflowItemAdgroupModifyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemAdgroupModifyAPIResponse)
	},
}

// GetTaobaoFeedflowItemAdgroupModifyAPIResponse 从 sync.Pool 获取 TaobaoFeedflowItemAdgroupModifyAPIResponse
func GetTaobaoFeedflowItemAdgroupModifyAPIResponse() *TaobaoFeedflowItemAdgroupModifyAPIResponse {
	return poolTaobaoFeedflowItemAdgroupModifyAPIResponse.Get().(*TaobaoFeedflowItemAdgroupModifyAPIResponse)
}

// ReleaseTaobaoFeedflowItemAdgroupModifyAPIResponse 将 TaobaoFeedflowItemAdgroupModifyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFeedflowItemAdgroupModifyAPIResponse(v *TaobaoFeedflowItemAdgroupModifyAPIResponse) {
	v.Reset()
	poolTaobaoFeedflowItemAdgroupModifyAPIResponse.Put(v)
}
