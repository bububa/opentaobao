package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusspacegroupgetlistbycampusandtype 根据园区id及TypeId获取空间分组
// alibaba.campus.space.group.getlistbycampusandtype
//
// 根据园区id及TypeId获取空间分组
// HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceGroupApiTopService
// HSF方法名称：getListByCampusAndType
func Alibabacampusspacegroupgetlistbycampusandtype(clt *core.SDKClient, req *campus.AlibabacampusspacegroupgetlistbycampusandtypeAPIRequest, session string) (*campus.AlibabacampusspacegroupgetlistbycampusandtypeAPIResponse, error) {
	var resp campus.AlibabacampusspacegroupgetlistbycampusandtypeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
