package qianniu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qianniu"
)

// Taobaoqianniutaskmetasget 任务元查询接口
// taobao.qianniu.taskmetas.get
//
// 任务元查询接口
func Taobaoqianniutaskmetasget(clt *core.SDKClient, req *qianniu.TaobaoqianniutaskmetasgetAPIRequest, session string) (*qianniu.TaobaoqianniutaskmetasgetAPIResponse, error) {
	var resp qianniu.TaobaoqianniutaskmetasgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
