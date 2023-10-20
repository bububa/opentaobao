package dt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/dt"
)

// TaobaoAdsDataImport 数据导入
// taobao.ads.data.import
//
// 数据导入
func TaobaoAdsDataImport(clt *core.SDKClient, req *dt.TaobaoAdsDataImportAPIRequest, session string) (*dt.TaobaoAdsDataImportAPIResponse, error) {
	var resp dt.TaobaoAdsDataImportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
