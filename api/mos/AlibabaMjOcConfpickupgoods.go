package mos

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMjOcConfpickupgoods 提货核销
// alibaba.mj.oc.confpickupgoods
//
// 此API用于在银泰商场中，消费者在提货中心提货时， 商户后台调用此接口进行提货核销操作
func AlibabaMjOcConfpickupgoods(ctx context.Context, clt *core.SDKClient, req *mos.AlibabaMjOcConfpickupgoodsAPIRequest, resp *mos.AlibabaMjOcConfpickupgoodsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
