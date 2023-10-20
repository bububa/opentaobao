package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Alibabajymitemexternalgoodsstatusbatchquery 交易猫外部商家商品状态批量查询接口
// alibaba.jym.item.external.goods.status.batch.query
//
// 供外部B端商家接入，请求查询商品状态，返回商品状态查询结果
func Alibabajymitemexternalgoodsstatusbatchquery(clt *core.SDKClient, req *product.AlibabajymitemexternalgoodsstatusbatchqueryAPIRequest, session string) (*product.AlibabajymitemexternalgoodsstatusbatchqueryAPIResponse, error) {
	var resp product.AlibabajymitemexternalgoodsstatusbatchqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
