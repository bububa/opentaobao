package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMiniappMesssageNormalSendAPIRequest
轻店铺下行普通消息给用户 API请求
taobao.miniapp.messsage.normal.send

小程序下行单个普通消息 */
type TaobaoMiniappMesssageNormalSendAPIRequest struct {
	model.Params
	// 普通消息结构
	_param *DownNormalMessageDto
}

// New
