package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelFscHotelModifyIncrement 酒店价库变更列表查询-供销平台
// taobao.alitrip.travel.fsc.hotel.modify.increment
//
// 按照时间纬度查询酒店变更列表
func TaobaoAlitripTravelFscHotelModifyIncrement(clt *core.SDKClient, req *axindata.TaobaoAlitripTravelFscHotelModifyIncrementAPIRequest, session string) (*axindata.TaobaoAlitripTravelFscHotelModifyIncrementAPIResponse, error) {
	var resp axindata.TaobaoAlitripTravelFscHotelModifyIncrementAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
