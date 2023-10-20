package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusspaceunitgetbyid 根据ID查询指定空间单元信息
// alibaba.campus.space.unit.getbyid
//
// 根据ID查询指定空间单元信息
// HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceUnitApiTopService
// HSF方法名称：getById
func Alibabacampusspaceunitgetbyid(clt *core.SDKClient, req *campus.AlibabacampusspaceunitgetbyidAPIRequest, session string) (*campus.AlibabacampusspaceunitgetbyidAPIResponse, error) {
	var resp campus.AlibabacampusspaceunitgetbyidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
