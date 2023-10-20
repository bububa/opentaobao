package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusspacefloorgetbybuildingid 根据楼宇ID获取楼层
// alibaba.campus.space.floor.getbybuildingid
//
// 根据楼宇ID获取楼层
// HSF接口名称：com.alibaba.campus.api.space.service.top.FloorApiTopService
// HSF方法名称：getFloorList
func Alibabacampusspacefloorgetbybuildingid(clt *core.SDKClient, req *campus.AlibabacampusspacefloorgetbybuildingidAPIRequest, session string) (*campus.AlibabacampusspacefloorgetbybuildingidAPIResponse, error) {
	var resp campus.AlibabacampusspacefloorgetbybuildingidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
