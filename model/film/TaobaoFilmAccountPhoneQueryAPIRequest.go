package film

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFilmAccountPhoneQueryAPIRequest
根据手机查询匹配账号列表 API请求
taobao.film.account.phone.query

根据手机号查询匹配的账号列表 */
type TaobaoFilmAccountPhoneQueryAPIRequest struct {
	model.Params
	// 11位手机号码
	_phone string
}

// New
