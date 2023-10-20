package tmc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTmcGroupDeleteAPIResponse 删除指定的分组或分组下的用户 API返回值
// taobao.tmc.group.delete
//
// 删除指定的分组或分组下的用户，授权消息使用
type TaobaoTmcGroupDeleteAPIResponse struct {
	model.CommonResponse
	TaobaoTmcGroupDeleteAPIResponseModel
}

// TaobaoTmcGroupDeleteAPIResponseModel is 删除指定的分组或分组下的用户 成功返回结果
type TaobaoTmcGroupDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"tmc_group_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
