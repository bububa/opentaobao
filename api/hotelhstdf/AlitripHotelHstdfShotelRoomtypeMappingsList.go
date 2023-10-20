package hotelhstdf

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/hotelhstdf"
)

// Alitriphotelhstdfshotelroomtypemappingslist 根据HID获取所有卖家房型匹配关系
// alitrip.hotel.hstdf.shotel.roomtype.mappings.list
//
// 根据HID获取所有卖家房型匹配关系
func Alitriphotelhstdfshotelroomtypemappingslist(clt *core.SDKClient, req *hotelhstdf.AlitriphotelhstdfshotelroomtypemappingslistAPIRequest, session string) (*hotelhstdf.AlitriphotelhstdfshotelroomtypemappingslistAPIResponse, error) {
	var resp hotelhstdf.AlitriphotelhstdfshotelroomtypemappingslistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
