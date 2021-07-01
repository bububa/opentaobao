package hotelalliance

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/hotelalliance"
)

/* AlitripHotelAllianceHidGet
获取联盟hid
alitrip.hotel.alliance.hid.get

获取符合条件的菲住联盟hid，目前支持指定日期上线的菲住联盟hid查询 */
func AlitripHotelAllianceHidGet(clt *core.SDKClient, req *hotelalliance.AlitripHotelAllianceHidGetAPIRequest, session string) (*hotelalliance.AlitripHotelAllianceHidGetAPIResponse, error) {
	var resp hotelalliance.AlitripHotelAllianceHidGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
