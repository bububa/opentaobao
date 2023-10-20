package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// Taobaoocapcontractsignedget 用户是否签署支付宝代扣协议
// taobao.oc.ap.contractsigned.get
//
// 用户是否签署支付宝代扣协议
func Taobaoocapcontractsignedget(clt *core.SDKClient, req *jst.TaobaoocapcontractsignedgetAPIRequest, session string) (*jst.TaobaoocapcontractsignedgetAPIResponse, error) {
	var resp jst.TaobaoocapcontractsignedgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
