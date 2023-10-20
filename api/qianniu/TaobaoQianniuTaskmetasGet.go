package qianniu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qianniu"
)

// TaobaoQianniuTaskmetasGet 任务元查询接口
// taobao.qianniu.taskmetas.get
//
// 任务元查询接口
func TaobaoQianniuTaskmetasGet(clt *core.SDKClient, req *qianniu.TaobaoQianniuTaskmetasGetAPIRequest, resp *qianniu.TaobaoQianniuTaskmetasGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
