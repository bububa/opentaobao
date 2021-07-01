package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJushitaJdpUserDeleteAPIRequest
删除数据推送用户 API请求
taobao.jushita.jdp.user.delete

删除应用的数据推送用户，用户被删除后，重新添加时会重新同步历史数据。 */
type TaobaoJushitaJdpUserDeleteAPIRequest struct {
	model.Params
	// 要删除用户的昵称
	_nick string
	// 需要删除的用户编号
	_userId int64
}

// New
