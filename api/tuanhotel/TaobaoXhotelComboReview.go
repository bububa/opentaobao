package tuanhotel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tuanhotel"
)

// Taobaoxhotelcomboreview 套餐审核接口
// taobao.xhotel.combo.review
//
// 套餐审核接口
func Taobaoxhotelcomboreview(clt *core.SDKClient, req *tuanhotel.TaobaoxhotelcomboreviewAPIRequest, session string) (*tuanhotel.TaobaoxhotelcomboreviewAPIResponse, error) {
	var resp tuanhotel.TaobaoxhotelcomboreviewAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
