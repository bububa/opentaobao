package feedflow

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowitemtargetvalidlistAPIResponse 获取有权限的定向列表 API返回值
// taobao.feedflow.item.target.validlist
//
// 获取有权限的定向列表
type TaobaofeedflowitemtargetvalidlistAPIResponse struct {
	model.CommonResponse
	TaobaofeedflowitemtargetvalidlistAPIResponseModel
}

// TaobaofeedflowitemtargetvalidlistAPIResponseModel is 获取有权限的定向列表 成功返回结果
type TaobaofeedflowitemtargetvalidlistAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_target_validlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果对象
	Result *TaobaofeedflowitemtargetvalidlistResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
