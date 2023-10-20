package feedflow

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemAdgroupAdzoneBindAPIResponse 信息流单元内绑定资源位 API返回值
// taobao.feedflow.item.adgroup.adzone.bind
//
// 信息流单元内绑定资源位
type TaobaoFeedflowItemAdgroupAdzoneBindAPIResponse struct {
	model.CommonResponse
	TaobaoFeedflowItemAdgroupAdzoneBindAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemAdgroupAdzoneBindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFeedflowItemAdgroupAdzoneBindAPIResponseModel).Reset()
}

// TaobaoFeedflowItemAdgroupAdzoneBindAPIResponseModel is 信息流单元内绑定资源位 成功返回结果
type TaobaoFeedflowItemAdgroupAdzoneBindAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_adgroup_adzone_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果对象
	Result *TaobaoFeedflowItemAdgroupAdzoneBindResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemAdgroupAdzoneBindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFeedflowItemAdgroupAdzoneBindAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemAdgroupAdzoneBindAPIResponse)
	},
}

// GetTaobaoFeedflowItemAdgroupAdzoneBindAPIResponse 从 sync.Pool 获取 TaobaoFeedflowItemAdgroupAdzoneBindAPIResponse
func GetTaobaoFeedflowItemAdgroupAdzoneBindAPIResponse() *TaobaoFeedflowItemAdgroupAdzoneBindAPIResponse {
	return poolTaobaoFeedflowItemAdgroupAdzoneBindAPIResponse.Get().(*TaobaoFeedflowItemAdgroupAdzoneBindAPIResponse)
}

// ReleaseTaobaoFeedflowItemAdgroupAdzoneBindAPIResponse 将 TaobaoFeedflowItemAdgroupAdzoneBindAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFeedflowItemAdgroupAdzoneBindAPIResponse(v *TaobaoFeedflowItemAdgroupAdzoneBindAPIResponse) {
	v.Reset()
	poolTaobaoFeedflowItemAdgroupAdzoneBindAPIResponse.Put(v)
}
