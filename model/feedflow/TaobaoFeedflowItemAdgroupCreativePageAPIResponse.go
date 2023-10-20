package feedflow

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowitemadgroupcreativepageAPIResponse 信息流单元下查看创意 API返回值
// taobao.feedflow.item.adgroup.creative.page
//
// 信息流单元下查看创意
type TaobaofeedflowitemadgroupcreativepageAPIResponse struct {
	model.CommonResponse
	TaobaofeedflowitemadgroupcreativepageAPIResponseModel
}

// TaobaofeedflowitemadgroupcreativepageAPIResponseModel is 信息流单元下查看创意 成功返回结果
type TaobaofeedflowitemadgroupcreativepageAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_adgroup_creative_page_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果返回消息
	Result *TaobaofeedflowitemadgroupcreativepageResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
