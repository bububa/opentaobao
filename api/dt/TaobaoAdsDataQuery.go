package dt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/dt"
)

// TaobaoAdsDataQuery 导入数据查询
// taobao.ads.data.query
//
// 导入数据查询
func TaobaoAdsDataQuery(clt *core.SDKClient, req *dt.TaobaoAdsDataQueryAPIRequest, resp *dt.TaobaoAdsDataQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
