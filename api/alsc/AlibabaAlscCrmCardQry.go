package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmCardQry 查询卡实例
// alibaba.alsc.crm.card.qry
//
// 查询卡实例（优先使用卡实例ID查询，没有则用物理卡号查询）
func AlibabaAlscCrmCardQry(clt *core.SDKClient, req *alsc.AlibabaAlscCrmCardQryAPIRequest, resp *alsc.AlibabaAlscCrmCardQryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
