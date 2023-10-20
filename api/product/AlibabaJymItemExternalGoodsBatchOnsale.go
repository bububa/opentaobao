package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Alibabajymitemexternalgoodsbatchonsale 交易猫外部商家批量上架商品接口
// alibaba.jym.item.external.goods.batch.onsale
//
// 供外部B端商家接入，提交批量上架商品请求，返回批量上架任务结果
func Alibabajymitemexternalgoodsbatchonsale(clt *core.SDKClient, req *product.AlibabajymitemexternalgoodsbatchonsaleAPIRequest, session string) (*product.AlibabajymitemexternalgoodsbatchonsaleAPIResponse, error) {
	var resp product.AlibabajymitemexternalgoodsbatchonsaleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
