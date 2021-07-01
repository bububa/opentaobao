package openim

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenimTribeSetmembernickAPIResponse
设置群成员昵称 API返回值
taobao.openim.tribe.setmembernick

设置群成员昵称，存在如下两种场景
1 群主或管理员设置群成员昵称，该操作有权限控制。只针对普通群的群主和管理员开发此功能；讨论组群主不支持此设置操作
2 群成员设置自己的昵称，该功能对群所有成员开放 */
type TaobaoOpenimTribeSetmembernickAPIResponse struct {
	model.CommonResponse
	TaobaoOpenimTribeSetmembernickAPIResponseModel
}

// TaobaoOpenimTribeSetmembernickAPIResponseModel is 设置群成员昵称 成功返回结果
type TaobaoOpenimTribeSetmembernickAPIResponseModel struct {
	XMLName xml.Name `xml:"openim_tribe_setmembernick_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	TribeCode int64 `json:"tribe_code,omitempty" xml:"tribe_code,omitempty"`
}
