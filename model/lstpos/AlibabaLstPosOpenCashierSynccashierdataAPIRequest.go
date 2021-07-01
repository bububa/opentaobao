package lstpos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstPosOpenCashierSynccashierdataAPIRequest
收银快照同步接口(最多10条订单信息) API请求
alibaba.lst.pos.open.cashier.synccashierdata

收银快照同步接口(最多10条订单信息) */
type AlibabaLstPosOpenCashierSynccashierdataAPIRequest struct {
	model.Params
	// 订单对象列表
	_cashierFlowDTOList []CashierFlowDto
	// 门店对应的主账号id
	_userId int64
}

// New
