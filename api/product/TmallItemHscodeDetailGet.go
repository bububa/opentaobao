package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TmallItemHscodeDetailGet 通过hscode获取计量单位
// tmall.item.hscode.detail.get
//
// 通过hscode获取计量单位和销售单位
func TmallItemHscodeDetailGet(clt *core.SDKClient, req *product.TmallItemHscodeDetailGetAPIRequest, session string) (*product.TmallItemHscodeDetailGetAPIResponse, error) {
	var resp product.TmallItemHscodeDetailGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
