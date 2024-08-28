package aliexpresssumaitong

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliexpresssumaitong"
)

// AliexpressTaxationPlatformOpenGet 平台固定参数获取
// aliexpress.taxation.platform.open.get
//
// Aliexpress开放平台固定参数获取
func AliexpressTaxationPlatformOpenGet(ctx context.Context, clt *core.SDKClient, req *aliexpresssumaitong.AliexpressTaxationPlatformOpenGetAPIRequest, resp *aliexpresssumaitong.AliexpressTaxationPlatformOpenGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
