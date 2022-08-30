package scs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scs"
)

// TaobaoOnebpDkxReportReportMaterialOffline 查询某计划分商品离线报表
// taobao.onebp.dkx.report.report.material.offline
//
// 查询某计划分商品离线报表
// 入参1示例：{"biz_code":"adStrategyDkx"}
// 入参2示例：{"start_time":"2021-09-23","campaign_id_list":[2853805001],"end_time":"2021-09-24","effect": 15   }
func TaobaoOnebpDkxReportReportMaterialOffline(clt *core.SDKClient, req *scs.TaobaoOnebpDkxReportReportMaterialOfflineAPIRequest, session string) (*scs.TaobaoOnebpDkxReportReportMaterialOfflineAPIResponse, error) {
	var resp scs.TaobaoOnebpDkxReportReportMaterialOfflineAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
