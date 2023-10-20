package hotelhstdf

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/hotelhstdf"
)

// AlitripHotelHstdfShotelExportshotel 商家自主导出相似度高的标准酒店
// alitrip.hotel.hstdf.shotel.exportshotel
//
// 商家通过给出自己的卖家酒店信息，通过接口可以返回相似度高的标准酒店信息
func AlitripHotelHstdfShotelExportshotel(clt *core.SDKClient, req *hotelhstdf.AlitripHotelHstdfShotelExportshotelAPIRequest, session string) (*hotelhstdf.AlitripHotelHstdfShotelExportshotelAPIResponse, error) {
	var resp hotelhstdf.AlitripHotelHstdfShotelExportshotelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
