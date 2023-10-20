package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// Tmallitemdapeitemplatequery 搭配查询接口
// tmall.item.dapei.template.query
//
// 根据条件获取搭配内容
func Tmallitemdapeitemplatequery(clt *core.SDKClient, req *product.TmallitemdapeitemplatequeryAPIRequest, session string) (*product.TmallitemdapeitemplatequeryAPIResponse, error) {
	var resp product.TmallitemdapeitemplatequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
