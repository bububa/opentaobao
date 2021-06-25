package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
推广计划下的推广组报表基础数据查询(只有汇总数据，无分类类型) APIResponse
taobao.simba.rpt.campadgroupbase.get

推广计划下的推广组报表基础数据查询(只有汇总数据，无分类类型)
*/
type TaobaoSimbaRptCampadgroupbaseGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaRptCampadgroupbaseGetResponse `json:"taobao_simba_rpt_campadgroupbase_get_response,omitempty"`
}

type TaobaoSimbaRptCampadgroupbaseGetResponse struct {

    // 推广计划下推广组的基础数据列表
    RptCampadgroupBaseList   string `json:"rpt_campadgroup_base_list,omitempty"`

}
