package sungari

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/sungari"
)

// TaobaoCloudbridgeCaseinvestExecute 红盾云桥案件协查服务
// taobao.cloudbridge.caseinvest.execute
//
// 通过API接口直接提供政府部门录入及查询函件服务
func TaobaoCloudbridgeCaseinvestExecute(clt *core.SDKClient, req *sungari.TaobaoCloudbridgeCaseinvestExecuteAPIRequest, resp *sungari.TaobaoCloudbridgeCaseinvestExecuteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
