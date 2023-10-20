package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Alibabajymitemexternalgoodsbatchpublish 交易猫外部商家批量发布商品接口
// alibaba.jym.item.external.goods.batch.publish
//
// 供外部B端商家接入，提交批量发布商品请求，返回批量发布任务结果
func Alibabajymitemexternalgoodsbatchpublish(clt *core.SDKClient, req *product.AlibabajymitemexternalgoodsbatchpublishAPIRequest, session string) (*product.AlibabajymitemexternalgoodsbatchpublishAPIResponse, error) {
	var resp product.AlibabajymitemexternalgoodsbatchpublishAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
