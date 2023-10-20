package tmc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotmcusergetAPIResponse 获取用户已开通消息 API返回值
// taobao.tmc.user.get
//
// 查询指定用户开通的消息通道和组
type TaobaotmcusergetAPIResponse struct {
	model.CommonResponse
	TaobaotmcusergetAPIResponseModel
}

// TaobaotmcusergetAPIResponseModel is 获取用户已开通消息 成功返回结果
type TaobaotmcusergetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmc_user_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 开通的用户数据
	TmcUser *TmcUser `json:"tmc_user,omitempty" xml:"tmc_user,omitempty"`
}
