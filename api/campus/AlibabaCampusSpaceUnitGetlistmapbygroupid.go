package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusspaceunitgetlistmapbygroupid 新增查询多个分组ID各自相关的空间单元信息
// alibaba.campus.space.unit.getlistmapbygroupid
//
// 新增查询多个分组ID各自相关的空间单元信息
// HSF接口名称：	com.alibaba.campus.api.space.service.top.SpaceUnitApiTopService
// HSF方法名称：	getListMapByGroupIds
func Alibabacampusspaceunitgetlistmapbygroupid(clt *core.SDKClient, req *campus.AlibabacampusspaceunitgetlistmapbygroupidAPIRequest, session string) (*campus.AlibabacampusspaceunitgetlistmapbygroupidAPIResponse, error) {
	var resp campus.AlibabacampusspaceunitgetlistmapbygroupidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
