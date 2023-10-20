package openmall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// Taobaoopenmallitemsaleareaget 查询商品可售区域
// taobao.openmall.item.salearea.get
//
// 获取openmall商品的可售区域
func Taobaoopenmallitemsaleareaget(clt *core.SDKClient, req *openmall.TaobaoopenmallitemsaleareagetAPIRequest, session string) (*openmall.TaobaoopenmallitemsaleareagetAPIResponse, error) {
	var resp openmall.TaobaoopenmallitemsaleareagetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
