package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIRequest
回收商同步前置补贴红包的代扣成功事件 API请求
taobao.recycle.ofnpreredpacket.tpdeductsuccess

回收商->天猫后端，同步前置补贴红包的代扣成功事件 */
type TaobaoRecycleOfnpreredpacketTpdeductsuccessAPIRequest struct {
	model.Params
	// 变化的金额
	_deductAmount int64
	// 旧机单id
	_oldOrderId int64
}

// New
