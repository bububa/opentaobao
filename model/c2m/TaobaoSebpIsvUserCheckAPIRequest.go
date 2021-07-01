package c2m

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSebpIsvUserCheckAPIRequest
淘小铺账户实名校验接口 API请求
taobao.sebp.isv.user.check

校验淘小铺账户和身份信息匹配成功 */
type TaobaoSebpIsvUserCheckAPIRequest struct {
	model.Params
	// 淘宝账号
	_userName string
	// 姓名
	_name string
	// 证件号
	_identity string
	// 支付宝账号
	_alipay string
}

// New
