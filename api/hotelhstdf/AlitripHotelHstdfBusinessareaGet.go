package hotelhstdf

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/hotelhstdf"
)

// AlitripHotelHstdfBusinessareaGet 根据城市查询商圈
// alitrip.hotel.hstdf.businessarea.get
//
// 根据cityId分页查询商圈信息
func AlitripHotelHstdfBusinessareaGet(clt *core.SDKClient, req *hotelhstdf.AlitripHotelHstdfBusinessareaGetAPIRequest, resp *hotelhstdf.AlitripHotelHstdfBusinessareaGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
