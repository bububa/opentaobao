package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Alibabajymitemexternalgoodsdetailquery 交易猫外部商家商品详情查询接口
// alibaba.jym.item.external.goods.detail.query
//
// 供外部B端商家接入，请求查询商品详情，返回商品详情查询结果
func Alibabajymitemexternalgoodsdetailquery(clt *core.SDKClient, req *product.AlibabajymitemexternalgoodsdetailqueryAPIRequest, session string) (*product.AlibabajymitemexternalgoodsdetailqueryAPIResponse, error) {
	var resp product.AlibabajymitemexternalgoodsdetailqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
