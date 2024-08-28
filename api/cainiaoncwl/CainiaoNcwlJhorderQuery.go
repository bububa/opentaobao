package cainiaoncwl

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaoncwl"
)

// CainiaoNcwlJhorderQuery 农村物流集货单查询接口
// cainiao.ncwl.jhorder.query
//
// 提供给接入商家，查询农村物流集货单
func CainiaoNcwlJhorderQuery(ctx context.Context, clt *core.SDKClient, req *cainiaoncwl.CainiaoNcwlJhorderQueryAPIRequest, resp *cainiaoncwl.CainiaoNcwlJhorderQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
