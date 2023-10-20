package feedflow

import (
	"encoding/xml"

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

// TaobaoFeedflowItemAdgroupAddAPIResponseModel is 信息流增加单元 成功返回结果
type TaobaoFeedflowItemAdgroupAddAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_adgroup_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *TaobaoFeedflowItemAdgroupAddResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
