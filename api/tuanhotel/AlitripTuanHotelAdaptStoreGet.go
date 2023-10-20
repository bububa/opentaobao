package tuanhotel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tuanhotel"
)

// Alitriptuanhoteladaptstoreget 酒店团购套餐关联适用门店
// alitrip.tuan.hotel.adapt.store.get
//
// 输入shid，返回关联门店详情信息
func Alitriptuanhoteladaptstoreget(clt *core.SDKClient, req *tuanhotel.AlitriptuanhoteladaptstoregetAPIRequest, session string) (*tuanhotel.AlitriptuanhoteladaptstoregetAPIResponse, error) {
	var resp tuanhotel.AlitriptuanhoteladaptstoregetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
