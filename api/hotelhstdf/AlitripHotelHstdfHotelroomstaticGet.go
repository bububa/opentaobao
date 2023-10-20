package hotelhstdf

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/hotelhstdf"
)

// Alitriphotelhstdfhotelroomstaticget 根据类型查询静态字段
// alitrip.hotel.hstdf.hotelroomstatic.get
//
// 根据类型查询分页静态字段
func Alitriphotelhstdfhotelroomstaticget(clt *core.SDKClient, req *hotelhstdf.AlitriphotelhstdfhotelroomstaticgetAPIRequest, session string) (*hotelhstdf.AlitriphotelhstdfhotelroomstaticgetAPIResponse, error) {
	var resp hotelhstdf.AlitriphotelhstdfhotelroomstaticgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
