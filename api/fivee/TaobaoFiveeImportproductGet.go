package fivee

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fivee"
)

// TaobaoFiveeImportproductGet 进口商品查询
// taobao.fivee.importproduct.get
//
// 资质共享平台查询进口商品信息
func TaobaoFiveeImportproductGet(clt *core.SDKClient, req *fivee.TaobaoFiveeImportproductGetAPIRequest, session string) (*fivee.TaobaoFiveeImportproductGetAPIResponse, error) {
	var resp fivee.TaobaoFiveeImportproductGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
