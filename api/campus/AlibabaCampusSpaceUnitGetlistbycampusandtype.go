package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusspaceunitgetlistbycampusandtype 根据园区id及TypeId获取空间单元
// alibaba.campus.space.unit.getlistbycampusandtype
//
// 根据园区id及TypeId获取空间单元
// HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceUnitApiTopService
// HSF方法名称：getListByCampusAndType
func Alibabacampusspaceunitgetlistbycampusandtype(clt *core.SDKClient, req *campus.AlibabacampusspaceunitgetlistbycampusandtypeAPIRequest, session string) (*campus.AlibabacampusspaceunitgetlistbycampusandtypeAPIResponse, error) {
	var resp campus.AlibabacampusspaceunitgetlistbycampusandtypeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
