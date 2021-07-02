package jms

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJushitaJmsGroupGetAPIResponse 查询ONS分组 API返回值
// taobao.jushita.jms.group.get
//
// 查询当前appkey在ONS中已有的分组
type TaobaoJushitaJmsGroupGetAPIResponse struct {
	model.CommonResponse
	TaobaoJushitaJmsGroupGetAPIResponseModel
}

// TaobaoJushitaJmsGroupGetAPIResponseModel is 查询ONS分组 成功返回结果
type TaobaoJushitaJmsGroupGetAPIResponseModel struct {
	XMLName xml.Name `xml:"jushita_jms_group_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的总数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
	// 分组信息
	Groups []MsgGroupDo `json:"groups,omitempty" xml:"groups>msg_group_do,omitempty"`
}
