package util

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// AlibabaRetailShorturlGet 短链接获取
// alibaba.retail.shorturl.get
//
// 短链接获取
func AlibabaRetailShorturlGet(ctx context.Context, clt *core.SDKClient, req *util.AlibabaRetailShorturlGetAPIRequest, resp *util.AlibabaRetailShorturlGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
