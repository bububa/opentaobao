package baichuan

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// AlibabaBaichuanTaopasswordQuery 查询解析淘口令
// alibaba.baichuan.taopassword.query
//
// 查询，解析淘口令
func AlibabaBaichuanTaopasswordQuery(ctx context.Context, clt *core.SDKClient, req *baichuan.AlibabaBaichuanTaopasswordQueryAPIRequest, resp *baichuan.AlibabaBaichuanTaopasswordQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
