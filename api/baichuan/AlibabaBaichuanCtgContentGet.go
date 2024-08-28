package baichuan

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// AlibabaBaichuanCtgContentGet 百川内容平台内容获取
// alibaba.baichuan.ctg.content.get
//
// 百川内容平台内容获取
func AlibabaBaichuanCtgContentGet(ctx context.Context, clt *core.SDKClient, req *baichuan.AlibabaBaichuanCtgContentGetAPIRequest, resp *baichuan.AlibabaBaichuanCtgContentGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
