package feedflow

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowitemadgroupmodifyAPIResponse 信息流单元修改 API返回值
// taobao.feedflow.item.adgroup.modify
//
// 信息流单元修改
type TaobaofeedflowitemadgroupmodifyAPIResponse struct {
	model.CommonResponse
	TaobaofeedflowitemadgroupmodifyAPIResponseModel
}

// TaobaofeedflowitemadgroupmodifyAPIResponseModel is 信息流单元修改 成功返回结果
type TaobaofeedflowitemadgroupmodifyAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_adgroup_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *TaobaofeedflowitemadgroupmodifyResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
