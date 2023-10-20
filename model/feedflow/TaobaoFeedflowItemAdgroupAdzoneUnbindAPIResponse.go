package feedflow

import (
	"encoding/xml"

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

// TaobaoFeedflowItemAdgroupAdzoneUnbindAPIResponseModel is 信息流单元内解绑资源位 成功返回结果
type TaobaoFeedflowItemAdgroupAdzoneUnbindAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_adgroup_adzone_unbind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果对象
	Result *TaobaoFeedflowItemAdgroupAdzoneUnbindResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
