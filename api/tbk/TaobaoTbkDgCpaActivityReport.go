package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// Taobaotbkdgcpaactivityreport 淘宝客-推广者-任务奖励效果报表
// taobao.tbk.dg.cpa.activity.report
//
// 提供给媒体使用的cpa活动报表查询api，当前仅试运行媒体可使用
func Taobaotbkdgcpaactivityreport(clt *core.SDKClient, req *tbk.TaobaotbkdgcpaactivityreportAPIRequest, session string) (*tbk.TaobaotbkdgcpaactivityreportAPIResponse, error) {
	var resp tbk.TaobaotbkdgcpaactivityreportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
