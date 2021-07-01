package subuser

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSubuserInfoUpdateAPIRequest
修改指定账户子账号的基本信息 API请求
taobao.subuser.info.update

修改指定账户子账号的基本信息（通过主账号登陆只能修改属于该主账号的子账号基本信息） */
type TaobaoSubuserInfoUpdateAPIRequest struct {
	model.Params
	// 是否停用子账号 true:表示停用该子账号false:表示开启该子账号
	_isDisableSubaccount bool
	// 子账号是否参与分流 true:参与分流 false:不参与分流
	_isDispatch bool
	// 子账号ID
	_subId int64
}

// New
