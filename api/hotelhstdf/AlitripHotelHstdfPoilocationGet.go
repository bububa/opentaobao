package hotelhstdf

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/hotelhstdf"
)

// Alitriphotelhstdfpoilocationget 根据平台城市id分页查询poi location
// alitrip.hotel.hstdf.poilocation.get
//
// 根据平台城市id分页查询poi location
func Alitriphotelhstdfpoilocationget(clt *core.SDKClient, req *hotelhstdf.AlitriphotelhstdfpoilocationgetAPIRequest, session string) (*hotelhstdf.AlitriphotelhstdfpoilocationgetAPIResponse, error) {
	var resp hotelhstdf.AlitriphotelhstdfpoilocationgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
