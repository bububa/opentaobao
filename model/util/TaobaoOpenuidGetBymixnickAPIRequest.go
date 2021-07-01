package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenuidGetBymixnickAPIRequest
通过mixnick转换openuid API请求
taobao.openuid.get.bymixnick

通过mixnick转换openuid */
type TaobaoOpenuidGetBymixnickAPIRequest struct {
	model.Params
	// 无线类应用获取到的混淆的nick
	_mixNick string
}

// New
