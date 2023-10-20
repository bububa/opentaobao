package fivee

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fivee"
)

// TaobaoFiveeImportproductGet 进口商品查询
// taobao.fivee.importproduct.get
//
// 资质共享平台查询进口商品信息
func TaobaoFiveeImportproductGet(clt *core.SDKClient, req *fivee.TaobaoFiveeImportproductGetAPIRequest, resp *fivee.TaobaoFiveeImportproductGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
