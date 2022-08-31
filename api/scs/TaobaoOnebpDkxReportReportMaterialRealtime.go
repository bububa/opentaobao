package scs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scs"
)

// TaobaoOnebpDkxReportReportMaterialRealtime 查询某计划分商品实时报表
// taobao.onebp.dkx.report.report.material.realtime
//
// 查询某计划分商品实时报表
// 入参1示例：{&#34;biz_code&#34;:&#34;adStrategyDkx&#34;}
// 入参2示例：{&#34;start_time&#34;:&#34;2021-09-24&#34;,&#34;campaign_id_list&#34;:[2853805001],&#34;end_time&#34;:&#34;2021-09-24&#34;,&#34;launch_product_id_list&#34;:[101011001,101001005,101001013,101001014,101016001]}
func TaobaoOnebpDkxReportReportMaterialRealtime(clt *core.SDKClient, req *scs.TaobaoOnebpDkxReportReportMaterialRealtimeAPIRequest, session string) (*scs.TaobaoOnebpDkxReportReportMaterialRealtimeAPIResponse, error) {
	var resp scs.TaobaoOnebpDkxReportReportMaterialRealtimeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
