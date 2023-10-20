package feedflow

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemAdgroupCreativeAddBindAPIResponse 信息流新增并且绑定创意 API返回值
// taobao.feedflow.item.adgroup.creative.add.bind
//
// 信息流新增并且绑定创意
type TaobaoFeedflowItemAdgroupCreativeAddBindAPIResponse struct {
	model.CommonResponse
	TaobaoFeedflowItemAdgroupCreativeAddBindAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemAdgroupCreativeAddBindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFeedflowItemAdgroupCreativeAddBindAPIResponseModel).Reset()
}

// TaobaoFeedflowItemAdgroupCreativeAddBindAPIResponseModel is 信息流新增并且绑定创意 成功返回结果
type TaobaoFeedflowItemAdgroupCreativeAddBindAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_adgroup_creative_add_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果对象
	Result *TaobaoFeedflowItemAdgroupCreativeAddBindResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemAdgroupCreativeAddBindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFeedflowItemAdgroupCreativeAddBindAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemAdgroupCreativeAddBindAPIResponse)
	},
}

// GetTaobaoFeedflowItemAdgroupCreativeAddBindAPIResponse 从 sync.Pool 获取 TaobaoFeedflowItemAdgroupCreativeAddBindAPIResponse
func GetTaobaoFeedflowItemAdgroupCreativeAddBindAPIResponse() *TaobaoFeedflowItemAdgroupCreativeAddBindAPIResponse {
	return poolTaobaoFeedflowItemAdgroupCreativeAddBindAPIResponse.Get().(*TaobaoFeedflowItemAdgroupCreativeAddBindAPIResponse)
}

// ReleaseTaobaoFeedflowItemAdgroupCreativeAddBindAPIResponse 将 TaobaoFeedflowItemAdgroupCreativeAddBindAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFeedflowItemAdgroupCreativeAddBindAPIResponse(v *TaobaoFeedflowItemAdgroupCreativeAddBindAPIResponse) {
	v.Reset()
	poolTaobaoFeedflowItemAdgroupCreativeAddBindAPIResponse.Put(v)
}
