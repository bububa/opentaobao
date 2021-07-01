package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAppstoreSubscribeGetAPIRequest
查询appstore应用订购关系 API请求
taobao.appstore.subscribe.get

查询appstore应用订购关系(对于新上架的多版本应用，建议使用taobao.vas.subscribe.get) */
type TaobaoAppstoreSubscribeGetAPIRequest struct {
	model.Params
	// 用户昵称
	_nick string
}

// New
