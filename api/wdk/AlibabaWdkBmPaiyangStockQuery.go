package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

/* AlibabaWdkBmPaiyangStockQuery
派样商品门店库存查询接口
alibaba.wdk.bm.paiyang.stock.query

淘鲜达接入第三方进行派样，第三方查询派样商品的门店库存信息。 */
func AlibabaWdkBmPaiyangStockQuery(clt *core.SDKClient, req *wdk.AlibabaWdkBmPaiyangStockQueryAPIRequest, session string) (*wdk.AlibabaWdkBmPaiyangStockQueryAPIResponse, error) {
	var resp wdk.AlibabaWdkBmPaiyangStockQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
