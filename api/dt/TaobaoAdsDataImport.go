package dt

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/dt"
)

// TaobaoAdsDataImport 数据导入
// taobao.ads.data.import
//
// 数据导入
func TaobaoAdsDataImport(ctx context.Context, clt *core.SDKClient, req *dt.TaobaoAdsDataImportAPIRequest, resp *dt.TaobaoAdsDataImportAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
