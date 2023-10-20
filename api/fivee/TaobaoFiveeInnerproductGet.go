package fivee

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fivee"
)

// Taobaofiveeinnerproductget 国产商品资质查询
// taobao.fivee.innerproduct.get
//
// 资质共享平台，国产商品查询
func Taobaofiveeinnerproductget(clt *core.SDKClient, req *fivee.TaobaofiveeinnerproductgetAPIRequest, session string) (*fivee.TaobaofiveeinnerproductgetAPIResponse, error) {
	var resp fivee.TaobaofiveeinnerproductgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
