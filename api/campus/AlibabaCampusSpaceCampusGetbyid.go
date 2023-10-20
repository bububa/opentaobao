package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusspacecampusgetbyid 根据园区id获取园区信息
// alibaba.campus.space.campus.getbyid
//
// 根据园区id获取园区信息
// HSF接口名称：com.alibaba.campus.api.space.service.top.CampusApiTopService
// HSF方法名称：getCampusById
func Alibabacampusspacecampusgetbyid(clt *core.SDKClient, req *campus.AlibabacampusspacecampusgetbyidAPIRequest, session string) (*campus.AlibabacampusspacecampusgetbyidAPIResponse, error) {
	var resp campus.AlibabacampusspacecampusgetbyidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
