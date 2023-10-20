package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Alibabajymitemgameseverquery 查询商品发布客户端下可用服务器列表
// alibaba.jym.item.game.sever.query
//
// 查询商品发布客户端下可用服务器列表
func Alibabajymitemgameseverquery(clt *core.SDKClient, req *product.AlibabajymitemgameseverqueryAPIRequest, session string) (*product.AlibabajymitemgameseverqueryAPIResponse, error) {
	var resp product.AlibabajymitemgameseverqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
