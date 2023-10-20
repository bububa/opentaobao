package dt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/dt"
)

// TaobaoAdsDataImport 数据导入
// taobao.ads.data.import
//
// 数据导入
func TaobaoAdsDataImport(clt *core.SDKClient, req *dt.TaobaoAdsDataImportAPIRequest, resp *dt.TaobaoAdsDataImportAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
