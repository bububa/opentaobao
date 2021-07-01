package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelBnbownerAddAPIRequest
民宿房东信息添加 API请求
taobao.xhotel.bnbowner.add

添加和更新民宿房东的信息 */
type TaobaoXhotelBnbownerAddAPIRequest struct {
	model.Params
	// 添加房东信息的对象
	_addOwnerParam *AddOwnerParam
}

// New
