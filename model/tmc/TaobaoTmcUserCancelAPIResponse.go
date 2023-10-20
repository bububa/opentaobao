package tmc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotmcusercancelAPIResponse 取消用户的消息服务 API返回值
// taobao.tmc.user.cancel
//
// 取消用户的消息服务
type TaobaotmcusercancelAPIResponse struct {
	model.CommonResponse
	TaobaotmcusercancelAPIResponseModel
}

// TaobaotmcusercancelAPIResponseModel is 取消用户的消息服务 成功返回结果
type TaobaotmcusercancelAPIResponseModel struct {
	XMLName xml.Name `xml:"tmc_user_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功,如果为false并且没有错误码，表示删除的用户不存在。
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
