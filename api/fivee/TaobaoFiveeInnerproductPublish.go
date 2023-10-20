package fivee

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fivee"
)

// Taobaofiveeinnerproductpublish 国产商品发布
// taobao.fivee.innerproduct.publish
//
// 资质共享平台国产商品发布
func Taobaofiveeinnerproductpublish(clt *core.SDKClient, req *fivee.TaobaofiveeinnerproductpublishAPIRequest, session string) (*fivee.TaobaofiveeinnerproductpublishAPIResponse, error) {
	var resp fivee.TaobaofiveeinnerproductpublishAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
