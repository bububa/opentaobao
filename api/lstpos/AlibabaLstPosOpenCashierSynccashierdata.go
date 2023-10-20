package lstpos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstpos"
)

// AlibabaLstPosOpenCashierSynccashierdata 收银快照同步接口(最多10条订单信息)
// alibaba.lst.pos.open.cashier.synccashierdata
//
// 收银快照同步接口(最多10条订单信息)
func AlibabaLstPosOpenCashierSynccashierdata(clt *core.SDKClient, req *lstpos.AlibabaLstPosOpenCashierSynccashierdataAPIRequest, resp *lstpos.AlibabaLstPosOpenCashierSynccashierdataAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
