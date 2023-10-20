package tmc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTmcGroupAddAPIResponse 为已开通用户添加用户分组 API返回值
// taobao.tmc.group.add
//
// 为已开通用户添加用户分组，授权消息使用
type TaobaoTmcGroupAddAPIResponse struct {
	model.CommonResponse
	TaobaoTmcGroupAddAPIResponseModel
}

// TaobaoTmcGroupAddAPIResponseModel is 为已开通用户添加用户分组 成功返回结果
type TaobaoTmcGroupAddAPIResponseModel struct {
	XMLName xml.Name `xml:"tmc_group_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 创建时间
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	// 分组名称
	GroupName string `json:"group_name,omitempty" xml:"group_name,omitempty"`
}
