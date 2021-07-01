package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenAccountTokenApplyAPIRequest
申请免登Open Account Token API请求
taobao.open.account.token.apply

申请免登Open Account Token */
type TaobaoOpenAccountTokenApplyAPIRequest struct {
	model.Params
	// 时间戳单位是毫秒
	_tokenTimestamp int64
	// open account id
	_openAccountId int64
	// isv自己账号的唯一id
	_isvAccountId string
	// 用于防重放的唯一id
	_uuid string
	// ISV APP的登录态时长单位是秒
	_loginStateExpireIn int64
	// 用于透传一些业务附加参数
	_ext string
}

// New
