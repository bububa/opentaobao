package jst

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJushitaJdpUserAddAPIResponse 添加数据推送用户 API返回值
// taobao.jushita.jdp.user.add
//
// 提供给接入数据推送的应用添加数据推送服务的用户
type TaobaoJushitaJdpUserAddAPIResponse struct {
	model.CommonResponse
	TaobaoJushitaJdpUserAddAPIResponseModel
}

// TaobaoJushitaJdpUserAddAPIResponseModel is 添加数据推送用户 成功返回结果
type TaobaoJushitaJdpUserAddAPIResponseModel struct {
	XMLName xml.Name `xml:"jushita_jdp_user_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否添加成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
