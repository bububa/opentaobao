package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoCardExpandcardQueryAPIRequest
购物金卡查询 API请求
taobao.card.expandcard.query

购物金充值信息查询接口，会返回余额等信息。 */
type TaobaoCardExpandcardQueryAPIRequest struct {
	model.Params
	// 卡使用范围，不传则会查询所有
	_usedScopeCode string
	// 支付宝accountNo
	_accountNo string
}

// New
