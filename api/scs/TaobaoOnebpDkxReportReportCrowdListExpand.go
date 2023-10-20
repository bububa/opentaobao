package scs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scs"
)

// TaobaoOnebpDkxReportReportCrowdListExpand 获取拓展人群数据报表
// taobao.onebp.dkx.report.report.crowd.list.expand
//
// 获取拓展人群数据报表
// 入参1示例：{&#34;biz_code&#34;:&#34;adStrategyDkx&#34;}
// 入参2示例：{&#34;effect&#34;:15,&#34;start_time&#34;:&#34;2021-09-08&#34;,&#34;end_time&#34;:&#34;2021-09-10&#34;,&#34;campaign_id_list&#34;:[2821811613],&#34;white_crowd_id_List&#34;:[12297883,12298696,12297989]}
func TaobaoOnebpDkxReportReportCrowdListExpand(clt *core.SDKClient, req *scs.TaobaoOnebpDkxReportReportCrowdListExpandAPIRequest, resp *scs.TaobaoOnebpDkxReportReportCrowdListExpandAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
