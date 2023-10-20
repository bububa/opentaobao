package xhotelonlineorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelonlineorder"
)

// Taobaoxhotelintlrateupdate 不落库商家推送更新酒店rate
// taobao.xhotel.intl.rate.update
//
// 商家主动推送不落库商品的酒店信息
func Taobaoxhotelintlrateupdate(clt *core.SDKClient, req *xhotelonlineorder.TaobaoxhotelintlrateupdateAPIRequest, session string) (*xhotelonlineorder.TaobaoxhotelintlrateupdateAPIResponse, error) {
	var resp xhotelonlineorder.TaobaoxhotelintlrateupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
