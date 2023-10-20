package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Alibabajymitemexternalgoodsbatchoffsale 交易猫外部商家批量下架商品接口
// alibaba.jym.item.external.goods.batch.offsale
//
// 供外部B端商家接入，提交批量下架商品请求，返回批量下架任务结果
func Alibabajymitemexternalgoodsbatchoffsale(clt *core.SDKClient, req *product.AlibabajymitemexternalgoodsbatchoffsaleAPIRequest, session string) (*product.AlibabajymitemexternalgoodsbatchoffsaleAPIResponse, error) {
	var resp product.AlibabajymitemexternalgoodsbatchoffsaleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
