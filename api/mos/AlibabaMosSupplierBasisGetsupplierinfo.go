package mos

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMosSupplierBasisGetsupplierinfo 获取供应商基础信息
// alibaba.mos.supplier.basis.getsupplierinfo
//
// 基于供应商id获取供应商基础脱敏信息
func AlibabaMosSupplierBasisGetsupplierinfo(ctx context.Context, clt *core.SDKClient, req *mos.AlibabaMosSupplierBasisGetsupplierinfoAPIRequest, resp *mos.AlibabaMosSupplierBasisGetsupplierinfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
