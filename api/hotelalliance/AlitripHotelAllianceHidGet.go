package hotelalliance

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/hotelalliance"
)

// Alitriphotelalliancehidget 获取联盟hid
// alitrip.hotel.alliance.hid.get
//
// 获取符合条件的菲住联盟hid，目前支持指定日期上线的菲住联盟hid查询
func Alitriphotelalliancehidget(clt *core.SDKClient, req *hotelalliance.AlitriphotelalliancehidgetAPIRequest, session string) (*hotelalliance.AlitriphotelalliancehidgetAPIResponse, error) {
	var resp hotelalliance.AlitriphotelalliancehidgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
