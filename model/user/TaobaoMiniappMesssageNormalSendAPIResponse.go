package user

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaominiappmesssagenormalsendAPIResponse 轻店铺下行普通消息给用户 API返回值
// taobao.miniapp.messsage.normal.send
//
// 小程序下行单个普通消息
type TaobaominiappmesssagenormalsendAPIResponse struct {
	model.CommonResponse
	TaobaominiappmesssagenormalsendAPIResponseModel
}

// TaobaominiappmesssagenormalsendAPIResponseModel is 轻店铺下行普通消息给用户 成功返回结果
type TaobaominiappmesssagenormalsendAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_messsage_normal_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回消息Id
	Model string `json:"model,omitempty" xml:"model,omitempty"`
}
