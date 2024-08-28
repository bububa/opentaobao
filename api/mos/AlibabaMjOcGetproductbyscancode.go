package mos

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMjOcGetproductbyscancode POS商品查询接口
// alibaba.mj.oc.getproductbyscancode
//
// 此API用于在银泰商场中，POS端扫码获取商品信息
func AlibabaMjOcGetproductbyscancode(ctx context.Context, clt *core.SDKClient, req *mos.AlibabaMjOcGetproductbyscancodeAPIRequest, resp *mos.AlibabaMjOcGetproductbyscancodeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
