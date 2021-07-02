package jms

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJushitaJmsGroupDeleteAPIResponse 删除ONS分组 API返回值
// taobao.jushita.jms.group.delete
//
// 删除ONS分组
type TaobaoJushitaJmsGroupDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoJushitaJmsGroupDeleteAPIResponseModel
}

// TaobaoJushitaJmsGroupDeleteAPIResponseModel is 删除ONS分组 成功返回结果
type TaobaoJushitaJmsGroupDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"jushita_jms_group_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作结果
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
