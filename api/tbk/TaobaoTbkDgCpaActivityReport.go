package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkDgCpaActivityReport 淘宝客-推广者-任务奖励效果报表
// taobao.tbk.dg.cpa.activity.report
//
// 提供给媒体使用的cpa活动报表查询api，当前仅试运行媒体可使用
func TaobaoTbkDgCpaActivityReport(clt *core.SDKClient, req *tbk.TaobaoTbkDgCpaActivityReportAPIRequest, session string) (*tbk.TaobaoTbkDgCpaActivityReportAPIResponse, error) {
	var resp tbk.TaobaoTbkDgCpaActivityReportAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
