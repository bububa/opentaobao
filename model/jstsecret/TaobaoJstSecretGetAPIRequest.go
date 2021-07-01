package jstsecret

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJstSecretGetAPIRequest
获取订单消费者的隐私号码 API请求
taobao.jst.secret.get

根据订单号获取消费者的隐私号 */
type TaobaoJstSecretGetAPIRequest struct {
	model.Params
	// 订单号
	_tid int64
	// 隐私号类型，1=手机号，2=分机号，默认为1；分机号使用时拨打手机号转分机号
	_type int64
	// 隐私号码过期天数，默认30天，取值范围[1,30]
	_expireDays int64
}

// New
