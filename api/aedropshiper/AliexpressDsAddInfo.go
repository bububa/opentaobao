package aedropshiper

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aedropshiper"
)

// AliexpressDsAddInfo 上报DS信息
// aliexpress.ds.add.info
//
// ISV用户上报下游DS信息
func AliexpressDsAddInfo(ctx context.Context, clt *core.SDKClient, req *aedropshiper.AliexpressDsAddInfoAPIRequest, resp *aedropshiper.AliexpressDsAddInfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
