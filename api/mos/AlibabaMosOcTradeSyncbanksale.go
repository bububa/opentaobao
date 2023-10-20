package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// Alibabamosoctradesyncbanksale 云闪付、银行卡销售数据回传接口
// alibaba.mos.oc.trade.syncbanksale
//
// 云闪付、银行卡销售数据回传
func Alibabamosoctradesyncbanksale(clt *core.SDKClient, req *mos.AlibabamosoctradesyncbanksaleAPIRequest, session string) (*mos.AlibabamosoctradesyncbanksaleAPIResponse, error) {
	var resp mos.AlibabamosoctradesyncbanksaleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
