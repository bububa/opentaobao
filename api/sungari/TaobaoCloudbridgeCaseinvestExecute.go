package sungari

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/sungari"
)

// Taobaocloudbridgecaseinvestexecute 红盾云桥案件协查服务
// taobao.cloudbridge.caseinvest.execute
//
// 通过API接口直接提供政府部门录入及查询函件服务
func Taobaocloudbridgecaseinvestexecute(clt *core.SDKClient, req *sungari.TaobaocloudbridgecaseinvestexecuteAPIRequest, session string) (*sungari.TaobaocloudbridgecaseinvestexecuteAPIResponse, error) {
	var resp sungari.TaobaocloudbridgecaseinvestexecuteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
