package axindata

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axindata"
)

// TaobaoAlitripTravelAxinHotelShidListQuery 阿信酒店分销-标准酒店id列表查询
// taobao.alitrip.travel.axin.hotel.shid.list.query
//
// 标准酒店id列表查询
func TaobaoAlitripTravelAxinHotelShidListQuery(clt *core.SDKClient, req *axindata.TaobaoAlitripTravelAxinHotelShidListQueryAPIRequest, session string) (*axindata.TaobaoAlitripTravelAxinHotelShidListQueryAPIResponse, error) {
	var resp axindata.TaobaoAlitripTravelAxinHotelShidListQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
