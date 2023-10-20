package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkbmpaiyangstockquery 派样商品门店库存查询接口
// alibaba.wdk.bm.paiyang.stock.query
//
// 淘鲜达接入第三方进行派样，第三方查询派样商品的门店库存信息。
func Alibabawdkbmpaiyangstockquery(clt *core.SDKClient, req *wdk.AlibabawdkbmpaiyangstockqueryAPIRequest, session string) (*wdk.AlibabawdkbmpaiyangstockqueryAPIResponse, error) {
	var resp wdk.AlibabawdkbmpaiyangstockqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
