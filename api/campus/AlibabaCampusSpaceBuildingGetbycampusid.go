package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusspacebuildinggetbycampusid 根据园区ID获取楼宇
// alibaba.campus.space.building.getbycampusid
//
// 根据园区ID获取楼宇
// HSF接口名称：com.alibaba.campus.api.space.service.top.BuildingApiTopService
// HSF方法名称：getBuildingList
func Alibabacampusspacebuildinggetbycampusid(clt *core.SDKClient, req *campus.AlibabacampusspacebuildinggetbycampusidAPIRequest, session string) (*campus.AlibabacampusspacebuildinggetbycampusidAPIResponse, error) {
	var resp campus.AlibabacampusspacebuildinggetbycampusidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
