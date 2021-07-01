package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenAccountUpdateAPIRequest
Open Account数据更新 API请求
taobao.open.account.update

Open Account数据更新 */
type TaobaoOpenAccountUpdateAPIRequest struct {
	model.Params
	// Open Account
	_paramList []OpenAccount
}

// New
