package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaouniversalbpreportasynccreatedownloadtask 创建异步下载任务
// taobao.universalbp.report.async.create.download.task
//
// 入参报表查询信息，出参下载任务id
func Taobaouniversalbpreportasynccreatedownloadtask(clt *core.SDKClient, req *simba.TaobaouniversalbpreportasynccreatedownloadtaskAPIRequest, session string) (*simba.TaobaouniversalbpreportasynccreatedownloadtaskAPIResponse, error) {
	var resp simba.TaobaouniversalbpreportasynccreatedownloadtaskAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
