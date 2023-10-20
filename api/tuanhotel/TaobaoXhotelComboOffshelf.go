package tuanhotel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tuanhotel"
)

// TaobaoXhotelComboOffshelf 酒店套餐下架
// taobao.xhotel.combo.offshelf
//
// 酒店套餐下架
func TaobaoXhotelComboOffshelf(clt *core.SDKClient, req *tuanhotel.TaobaoXhotelComboOffshelfAPIRequest, session string) (*tuanhotel.TaobaoXhotelComboOffshelfAPIResponse, error) {
	var resp tuanhotel.TaobaoXhotelComboOffshelfAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
