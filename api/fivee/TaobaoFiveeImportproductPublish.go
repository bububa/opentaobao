package fivee

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fivee"
)

// Taobaofiveeimportproductpublish 进口商品发布
// taobao.fivee.importproduct.publish
//
// 直营业务商家入住发布商品时，上传商品及商家证照信息
func Taobaofiveeimportproductpublish(clt *core.SDKClient, req *fivee.TaobaofiveeimportproductpublishAPIRequest, session string) (*fivee.TaobaofiveeimportproductpublishAPIResponse, error) {
	var resp fivee.TaobaofiveeimportproductpublishAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
