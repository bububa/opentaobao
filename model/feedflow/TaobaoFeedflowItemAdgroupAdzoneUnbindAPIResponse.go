package feedflow

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemAdgroupAdzoneUnbindAPIResponse 信息流单元内解绑资源位 API返回值
// taobao.feedflow.item.adgroup.adzone.unbind
//
// 信息流单元内解绑资源位
type TaobaoFeedflowItemAdgroupAdzoneUnbindAPIResponse struct {
	model.CommonResponse
	TaobaoFeedflowItemAdgroupAdzoneUnbindAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemAdgroupAdzoneUnbindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFeedflowItemAdgroupAdzoneUnbindAPIResponseModel).Reset()
}

// TaobaoFeedflowItemAdgroupAdzoneUnbindAPIResponseModel is 信息流单元内解绑资源位 成功返回结果
type TaobaoFeedflowItemAdgroupAdzoneUnbindAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_adgroup_adzone_unbind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果对象
	Result *TaobaoFeedflowItemAdgroupAdzoneUnbindResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemAdgroupAdzoneUnbindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFeedflowItemAdgroupAdzoneUnbindAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemAdgroupAdzoneUnbindAPIResponse)
	},
}

// GetTaobaoFeedflowItemAdgroupAdzoneUnbindAPIResponse 从 sync.Pool 获取 TaobaoFeedflowItemAdgroupAdzoneUnbindAPIResponse
func GetTaobaoFeedflowItemAdgroupAdzoneUnbindAPIResponse() *TaobaoFeedflowItemAdgroupAdzoneUnbindAPIResponse {
	return poolTaobaoFeedflowItemAdgroupAdzoneUnbindAPIResponse.Get().(*TaobaoFeedflowItemAdgroupAdzoneUnbindAPIResponse)
}

// ReleaseTaobaoFeedflowItemAdgroupAdzoneUnbindAPIResponse 将 TaobaoFeedflowItemAdgroupAdzoneUnbindAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFeedflowItemAdgroupAdzoneUnbindAPIResponse(v *TaobaoFeedflowItemAdgroupAdzoneUnbindAPIResponse) {
	v.Reset()
	poolTaobaoFeedflowItemAdgroupAdzoneUnbindAPIResponse.Put(v)
}
