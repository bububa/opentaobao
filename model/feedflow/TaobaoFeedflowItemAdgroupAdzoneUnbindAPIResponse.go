package feedflow

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowitemadgroupadzoneunbindAPIResponse 信息流单元内解绑资源位 API返回值
// taobao.feedflow.item.adgroup.adzone.unbind
//
// 信息流单元内解绑资源位
type TaobaofeedflowitemadgroupadzoneunbindAPIResponse struct {
	model.CommonResponse
	TaobaofeedflowitemadgroupadzoneunbindAPIResponseModel
}

// TaobaofeedflowitemadgroupadzoneunbindAPIResponseModel is 信息流单元内解绑资源位 成功返回结果
type TaobaofeedflowitemadgroupadzoneunbindAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_adgroup_adzone_unbind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果对象
	Result *TaobaofeedflowitemadgroupadzoneunbindResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
