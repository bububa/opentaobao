package baichuan

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// AlibabaBaichuanTaopasswordCheck 淘口令检查
// alibaba.baichuan.taopassword.check
//
// 检查当前文本是否为淘口令
func AlibabaBaichuanTaopasswordCheck(ctx context.Context, clt *core.SDKClient, req *baichuan.AlibabaBaichuanTaopasswordCheckAPIRequest, resp *baichuan.AlibabaBaichuanTaopasswordCheckAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
