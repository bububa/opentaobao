package baichuan

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// AlibabaBaichuanTaopasswordConfig 淘口令配置数据
// alibaba.baichuan.taopassword.config
//
// 百川淘口令规则配置接口
func AlibabaBaichuanTaopasswordConfig(ctx context.Context, clt *core.SDKClient, req *baichuan.AlibabaBaichuanTaopasswordConfigAPIRequest, resp *baichuan.AlibabaBaichuanTaopasswordConfigAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
