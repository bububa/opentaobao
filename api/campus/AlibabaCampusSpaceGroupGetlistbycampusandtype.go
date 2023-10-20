package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// AlibabaCampusSpaceGroupGetlistbycampusandtype 根据园区id及TypeId获取空间分组
// alibaba.campus.space.group.getlistbycampusandtype
//
// 根据园区id及TypeId获取空间分组
// HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceGroupApiTopService
// HSF方法名称：getListByCampusAndType
func AlibabaCampusSpaceGroupGetlistbycampusandtype(clt *core.SDKClient, req *campus.AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIRequest, resp *campus.AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
