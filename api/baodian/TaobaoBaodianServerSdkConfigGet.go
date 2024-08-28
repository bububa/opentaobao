package baodian

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baodian"
)

// TaobaoBaodianServerSdkConfigGet 获取宝点SDK的配置项（已迁移）
// taobao.baodian.server.sdk.config.get
//
// 获取SDK各种配置项（已迁移）
func TaobaoBaodianServerSdkConfigGet(ctx context.Context, clt *core.SDKClient, req *baodian.TaobaoBaodianServerSdkConfigGetAPIRequest, resp *baodian.TaobaoBaodianServerSdkConfigGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
