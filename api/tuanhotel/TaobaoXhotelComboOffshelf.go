package tuanhotel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tuanhotel"
)

// Taobaoxhotelcombooffshelf 酒店套餐下架
// taobao.xhotel.combo.offshelf
//
// 酒店套餐下架
func Taobaoxhotelcombooffshelf(clt *core.SDKClient, req *tuanhotel.TaobaoxhotelcombooffshelfAPIRequest, session string) (*tuanhotel.TaobaoxhotelcombooffshelfAPIResponse, error) {
	var resp tuanhotel.TaobaoxhotelcombooffshelfAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
