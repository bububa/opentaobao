package fivee

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fivee"
)

// TaobaoFiveeInnerproductPublish 国产商品发布
// taobao.fivee.innerproduct.publish
//
// 资质共享平台国产商品发布
func TaobaoFiveeInnerproductPublish(clt *core.SDKClient, req *fivee.TaobaoFiveeInnerproductPublishAPIRequest, session string) (*fivee.TaobaoFiveeInnerproductPublishAPIResponse, error) {
	var resp fivee.TaobaoFiveeInnerproductPublishAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
