package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoOcApContractsignedGet 用户是否签署支付宝代扣协议
// taobao.oc.ap.contractsigned.get
//
// 用户是否签署支付宝代扣协议
func TaobaoOcApContractsignedGet(clt *core.SDKClient, req *jst.TaobaoOcApContractsignedGetAPIRequest, resp *jst.TaobaoOcApContractsignedGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
