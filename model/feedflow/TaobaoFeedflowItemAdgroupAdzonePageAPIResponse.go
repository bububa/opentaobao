package feedflow

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemAdgroupAdzonePageAPIResponse 信息流单元下查看绑定资源位 API返回值
// taobao.feedflow.item.adgroup.adzone.page
//
// 信息流单元下查看绑定资源位
type TaobaoFeedflowItemAdgroupAdzonePageAPIResponse struct {
	model.CommonResponse
	TaobaoFeedflowItemAdgroupAdzonePageAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemAdgroupAdzonePageAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFeedflowItemAdgroupAdzonePageAPIResponseModel).Reset()
}

// TaobaoFeedflowItemAdgroupAdzonePageAPIResponseModel is 信息流单元下查看绑定资源位 成功返回结果
type TaobaoFeedflowItemAdgroupAdzonePageAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_adgroup_adzone_page_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果对象
	Result *TaobaoFeedflowItemAdgroupAdzonePageResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemAdgroupAdzonePageAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFeedflowItemAdgroupAdzonePageAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemAdgroupAdzonePageAPIResponse)
	},
}

// GetTaobaoFeedflowItemAdgroupAdzonePageAPIResponse 从 sync.Pool 获取 TaobaoFeedflowItemAdgroupAdzonePageAPIResponse
func GetTaobaoFeedflowItemAdgroupAdzonePageAPIResponse() *TaobaoFeedflowItemAdgroupAdzonePageAPIResponse {
	return poolTaobaoFeedflowItemAdgroupAdzonePageAPIResponse.Get().(*TaobaoFeedflowItemAdgroupAdzonePageAPIResponse)
}

// ReleaseTaobaoFeedflowItemAdgroupAdzonePageAPIResponse 将 TaobaoFeedflowItemAdgroupAdzonePageAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFeedflowItemAdgroupAdzonePageAPIResponse(v *TaobaoFeedflowItemAdgroupAdzonePageAPIResponse) {
	v.Reset()
	poolTaobaoFeedflowItemAdgroupAdzonePageAPIResponse.Put(v)
}
