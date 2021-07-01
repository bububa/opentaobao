package uscesl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUsceslBizEslUnbindAPIRequest
电子价签解绑接口 API请求
taobao.uscesl.biz.esl.unbind

电子价签解绑接口 */
type TaobaoUsceslBizEslUnbindAPIRequest struct {
	model.Params
	// 价签条码
	_eslBarCode string
	// 价签系统注册的门店storeId
	_storeId int64
	// 商家标识key
	_bizBrandKey string
}

// New
