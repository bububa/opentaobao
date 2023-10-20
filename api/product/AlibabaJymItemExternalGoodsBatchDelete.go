package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Alibabajymitemexternalgoodsbatchdelete 交易猫外部商家批量删除商品接口
// alibaba.jym.item.external.goods.batch.delete
//
// 交易猫外部商家批量删除商品接口
func Alibabajymitemexternalgoodsbatchdelete(clt *core.SDKClient, req *product.AlibabajymitemexternalgoodsbatchdeleteAPIRequest, session string) (*product.AlibabajymitemexternalgoodsbatchdeleteAPIResponse, error) {
	var resp product.AlibabajymitemexternalgoodsbatchdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
