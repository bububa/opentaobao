package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TaobaoItemBarcodeUpdate 更新商品条形码信息
// taobao.item.barcode.update
//
// 通过该接口，将商品以及SKU上得条形码信息补全
func TaobaoItemBarcodeUpdate(clt *core.SDKClient, req *tbitem.TaobaoItemBarcodeUpdateAPIRequest, resp *tbitem.TaobaoItemBarcodeUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
