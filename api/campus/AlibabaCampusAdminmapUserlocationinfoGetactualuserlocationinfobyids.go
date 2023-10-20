package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusadminmapuserlocationinfogetactualuserlocationinfobyids 根据userId(支持单个或批量)获取用户实时位置信息
// alibaba.campus.adminmap.userlocationinfo.getactualuserlocationinfobyids
//
// 根据userId(支持单个或批量)获取用户实时位置信息
// HSF接口名称：com.alibaba.campus.api.adminmap.service.top.UserLocationQueryApiTopService
// HSF方法名称：getActualUserLocationInfoByIds
func Alibabacampusadminmapuserlocationinfogetactualuserlocationinfobyids(clt *core.SDKClient, req *campus.AlibabacampusadminmapuserlocationinfogetactualuserlocationinfobyidsAPIRequest, session string) (*campus.AlibabacampusadminmapuserlocationinfogetactualuserlocationinfobyidsAPIResponse, error) {
	var resp campus.AlibabacampusadminmapuserlocationinfogetactualuserlocationinfobyidsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
