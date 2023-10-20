package openim

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenimrelationsgetAPIResponse 获取openim账号的聊天关系 API返回值
// taobao.openim.relations.get
//
// 获取openim账号的聊天关系
type TaobaoopenimrelationsgetAPIResponse struct {
	model.CommonResponse
	TaobaoopenimrelationsgetAPIResponseModel
}

// TaobaoopenimrelationsgetAPIResponseModel is 获取openim账号的聊天关系 成功返回结果
type TaobaoopenimrelationsgetAPIResponseModel struct {
	XMLName xml.Name `xml:"openim_relations_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 用户信息列表
	Users []OpenImUser `json:"users,omitempty" xml:"users>open_im_user,omitempty"`
}
