package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusspaceunitgetlistbygroupid 根据分组ID查询相应的空间单元
// alibaba.campus.space.unit.getlistbygroupid
//
// 根据分组ID查询相应的空间单元
// HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceUnitApiTopService
// HSF方法名称：getListByGroupId
func Alibabacampusspaceunitgetlistbygroupid(clt *core.SDKClient, req *campus.AlibabacampusspaceunitgetlistbygroupidAPIRequest, session string) (*campus.AlibabacampusspaceunitgetlistbygroupidAPIResponse, error) {
	var resp campus.AlibabacampusspaceunitgetlistbygroupidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
