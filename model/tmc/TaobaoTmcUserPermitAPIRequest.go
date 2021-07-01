package tmc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTmcUserPermitAPIRequest
为已授权的用户开通消息服务 API请求
taobao.tmc.user.permit

为已授权的用户开通消息服务，授权消息使用。<br/><span style="color:red">注意：topic覆盖更新,务必传入全量topic，或者不传topics，使用appkey订阅的所有topic</span> */
type TaobaoTmcUserPermitAPIRequest struct {
	model.Params
	// 消息主题列表，用半角逗号分隔。当用户订阅的topic是应用订阅的子集时才需要设置，不设置表示继承应用所订阅的所有topic，一般情况建议不要设置。
	_topics []string
}

// New
