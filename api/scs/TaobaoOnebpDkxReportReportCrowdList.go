package scs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scs"
)

// TaobaoOnebpDkxReportReportCrowdList 获取人群离线报表
// taobao.onebp.dkx.report.report.crowd.list
//
// 获取人群离线报表
// 入参1示例：{"biz_code":"adStrategyDkx"}
// 入参2示例：{"start_time":"2021-09-08","campaign_id_list":[2821811613],"effect":15,"end_time":"2021-09-10","crowd_id":12297883}
func TaobaoOnebpDkxReportReportCrowdList(clt *core.SDKClient, req *scs.TaobaoOnebpDkxReportReportCrowdListAPIRequest, session string) (*scs.TaobaoOnebpDkxReportReportCrowdListAPIResponse, error) {
	var resp scs.TaobaoOnebpDkxReportReportCrowdListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
