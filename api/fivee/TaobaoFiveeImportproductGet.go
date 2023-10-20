package fivee

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fivee"
)

// Taobaofiveeimportproductget 进口商品查询
// taobao.fivee.importproduct.get
//
// 资质共享平台查询进口商品信息
func Taobaofiveeimportproductget(clt *core.SDKClient, req *fivee.TaobaofiveeimportproductgetAPIRequest, session string) (*fivee.TaobaofiveeimportproductgetAPIResponse, error) {
	var resp fivee.TaobaofiveeimportproductgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
