package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusSpaceUnitGetlistbycampusandtype 根据园区id及TypeId获取空间单元
// alibaba.campus.space.unit.getlistbycampusandtype
//
// 根据园区id及TypeId获取空间单元
// HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceUnitApiTopService
// HSF方法名称：getListByCampusAndType
func AlibabaCampusSpaceUnitGetlistbycampusandtype(clt *core.SDKClient, req *campus.AlibabaCampusSpaceUnitGetlistbycampusandtypeAPIRequest, resp *campus.AlibabaCampusSpaceUnitGetlistbycampusandtypeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
