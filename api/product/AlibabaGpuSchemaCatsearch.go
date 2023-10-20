package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Alibabagpuschemacatsearch 按类目查询spu接口
// alibaba.gpu.schema.catsearch
//
// 按类目查询spu的schema接口
func Alibabagpuschemacatsearch(clt *core.SDKClient, req *product.AlibabagpuschemacatsearchAPIRequest, session string) (*product.AlibabagpuschemacatsearchAPIResponse, error) {
	var resp product.AlibabagpuschemacatsearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
