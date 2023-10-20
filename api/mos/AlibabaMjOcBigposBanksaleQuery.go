package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// Alibabamjocbigposbanksalequery 大pos银行卡查账接口
// alibaba.mj.oc.bigpos.banksale.query
//
// 大pos银行卡查账接口，给收银员查询银行卡销售记录，便于调账
func Alibabamjocbigposbanksalequery(clt *core.SDKClient, req *mos.AlibabamjocbigposbanksalequeryAPIRequest, session string) (*mos.AlibabamjocbigposbanksalequeryAPIResponse, error) {
	var resp mos.AlibabamjocbigposbanksalequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
