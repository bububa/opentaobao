package hotelhstdf

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/hotelhstdf"
)

// AlitripHotelHstdfShotelExportsroomtype 导出一个卖家房型下的所有标准房型
// alitrip.hotel.hstdf.shotel.exportsroomtype
//
// 导出一个卖家酒店下的所有标准房型
func AlitripHotelHstdfShotelExportsroomtype(clt *core.SDKClient, req *hotelhstdf.AlitripHotelHstdfShotelExportsroomtypeAPIRequest, resp *hotelhstdf.AlitripHotelHstdfShotelExportsroomtypeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
