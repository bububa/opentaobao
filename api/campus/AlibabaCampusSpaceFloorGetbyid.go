package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusSpaceFloorGetbyid 根据id获取楼层
// alibaba.campus.space.floor.getbyid
//
// 根据id获取楼层
func AlibabaCampusSpaceFloorGetbyid(clt *core.SDKClient, req *campus.AlibabaCampusSpaceFloorGetbyidAPIRequest, session string) (*campus.AlibabaCampusSpaceFloorGetbyidAPIResponse, error) {
	var resp campus.AlibabaCampusSpaceFloorGetbyidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
