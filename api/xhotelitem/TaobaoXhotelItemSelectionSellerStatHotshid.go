package xhotelitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelitem"
)

// Taobaoxhotelitemselectionsellerstathotshid 供应链选品热销标准酒店覆盖情况
// taobao.xhotel.item.selection.seller.stat.hotshid
//
// 供应链选品热销标准酒店覆盖情况
func Taobaoxhotelitemselectionsellerstathotshid(clt *core.SDKClient, req *xhotelitem.TaobaoxhotelitemselectionsellerstathotshidAPIRequest, session string) (*xhotelitem.TaobaoxhotelitemselectionsellerstathotshidAPIResponse, error) {
	var resp xhotelitem.TaobaoxhotelitemselectionsellerstathotshidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
