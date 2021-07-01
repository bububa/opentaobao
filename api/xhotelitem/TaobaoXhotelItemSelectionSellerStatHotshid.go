package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

/* TaobaoXhotelItemSelectionSellerStatHotshid
供应链选品热销标准酒店覆盖情况
taobao.xhotel.item.selection.seller.stat.hotshid

供应链选品热销标准酒店覆盖情况 */
func TaobaoXhotelItemSelectionSellerStatHotshid(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelItemSelectionSellerStatHotshidAPIRequest, session string) (*xhotelitem.TaobaoXhotelItemSelectionSellerStatHotshidAPIResponse, error) {
	var resp xhotelitem.TaobaoXhotelItemSelectionSellerStatHotshidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
