package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

/* TmallItemCalculateHscodeGet
算法获取hscode
tmall.item.calculate.hscode.get

算法获取hscode */
func TmallItemCalculateHscodeGet(clt *core.SDKClient, req *product.TmallItemCalculateHscodeGetAPIRequest, session string) (*product.TmallItemCalculateHscodeGetAPIResponse, error) {
	var resp product.TmallItemCalculateHscodeGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
