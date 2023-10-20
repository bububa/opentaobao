package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusspaceunitgetlist 多条件查询空间单元信息
// alibaba.campus.space.unit.getlist
//
// 多条件查询空间单元信息
// HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceUnitApiTopService
// HSF方法名称：getList
func Alibabacampusspaceunitgetlist(clt *core.SDKClient, req *campus.AlibabacampusspaceunitgetlistAPIRequest, session string) (*campus.AlibabacampusspaceunitgetlistAPIResponse, error) {
	var resp campus.AlibabacampusspaceunitgetlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
