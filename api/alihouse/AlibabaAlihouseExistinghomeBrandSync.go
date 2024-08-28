package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeBrandSync 二手房公司品牌数据同步
// alibaba.alihouse.existinghome.brand.sync
//
// 二手房公司品牌数据同步
func AlibabaAlihouseExistinghomeBrandSync(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeBrandSyncAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeBrandSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
