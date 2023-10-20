package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// Alibabamjocbigposbanksaleadjustmentapply 大pos银行卡调账申请
// alibaba.mj.oc.bigpos.banksale.adjustment.apply
//
// 大pos银行卡调账申请
func Alibabamjocbigposbanksaleadjustmentapply(clt *core.SDKClient, req *mos.AlibabamjocbigposbanksaleadjustmentapplyAPIRequest, session string) (*mos.AlibabamjocbigposbanksaleadjustmentapplyAPIResponse, error) {
	var resp mos.AlibabamjocbigposbanksaleadjustmentapplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
