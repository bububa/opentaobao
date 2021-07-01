package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenAccountTokenValidateAPIRequest
open account token验证 API请求
taobao.open.account.token.validate

open account token验证 */
type TaobaoOpenAccountTokenValidateAPIRequest struct {
	model.Params
	// token
	_paramToken string
}

// New
