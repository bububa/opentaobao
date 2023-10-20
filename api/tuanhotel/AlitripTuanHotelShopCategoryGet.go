package tuanhotel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tuanhotel"
)

// Alitriptuanhotelshopcategoryget 商家店铺类目查询
// alitrip.tuan.hotel.shop.category.get
//
// 查询商家店铺类目信息
func Alitriptuanhotelshopcategoryget(clt *core.SDKClient, req *tuanhotel.AlitriptuanhotelshopcategorygetAPIRequest, session string) (*tuanhotel.AlitriptuanhotelshopcategorygetAPIResponse, error) {
	var resp tuanhotel.AlitriptuanhotelshopcategorygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
