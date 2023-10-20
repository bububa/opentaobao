package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmOpenRechargeOperate 储值操作接口
// alibaba.alsc.crm.open.recharge.operate
//
// 储值操作接口
func AlibabaAlscCrmOpenRechargeOperate(clt *core.SDKClient, req *alsc.AlibabaAlscCrmOpenRechargeOperateAPIRequest, resp *alsc.AlibabaAlscCrmOpenRechargeOperateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
