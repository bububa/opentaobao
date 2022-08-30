package scs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scs"
)

// TaobaoOnebpDkxReportReportCrowdListExpand 获取拓展人群数据报表
// taobao.onebp.dkx.report.report.crowd.list.expand
//
// 获取拓展人群数据报表
// 入参1示例：{"biz_code":"adStrategyDkx"}
// 入参2示例：{"effect":15,"start_time":"2021-09-08","end_time":"2021-09-10","campaign_id_list":[2821811613],"white_crowd_id_List":[12297883,12298696,12297989]}
func TaobaoOnebpDkxReportReportCrowdListExpand(clt *core.SDKClient, req *scs.TaobaoOnebpDkxReportReportCrowdListExpandAPIRequest, session string) (*scs.TaobaoOnebpDkxReportReportCrowdListExpandAPIResponse, error) {
	var resp scs.TaobaoOnebpDkxReportReportCrowdListExpandAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
