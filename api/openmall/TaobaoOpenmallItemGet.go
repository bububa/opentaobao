package openmall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// Taobaoopenmallitemget 获取商品详情物料
// taobao.openmall.item.get
//
// 获取联盟开放的openmall商品
func Taobaoopenmallitemget(clt *core.SDKClient, req *openmall.TaobaoopenmallitemgetAPIRequest, session string) (*openmall.TaobaoopenmallitemgetAPIResponse, error) {
	var resp openmall.TaobaoopenmallitemgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
