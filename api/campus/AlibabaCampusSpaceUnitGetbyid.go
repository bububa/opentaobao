package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusSpaceUnitGetbyid 根据ID查询指定空间单元信息
// alibaba.campus.space.unit.getbyid
//
// 根据ID查询指定空间单元信息
// HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceUnitApiTopService
// HSF方法名称：getById
func AlibabaCampusSpaceUnitGetbyid(clt *core.SDKClient, req *campus.AlibabaCampusSpaceUnitGetbyidAPIRequest, resp *campus.AlibabaCampusSpaceUnitGetbyidAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
