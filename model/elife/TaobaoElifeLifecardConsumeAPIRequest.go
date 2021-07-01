package elife

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoElifeLifecardConsumeAPIRequest
品牌惠卡券核销 API请求
taobao.elife.lifecard.consume

用户线上购买生活汇品牌惠虚拟消费卡，线下购物时，商家码枪核销，涉及用户虚拟卡余额扣减操作 */
type TaobaoElifeLifecardConsumeAPIRequest struct {
	model.Params
	// 交易请求参数
	_consumeRequest *ConsumeRequest
}

// New
