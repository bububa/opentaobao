package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
类目信息获取 APIResponse
taobao.simba.insight.catsinfo.get

获取类目信息，此接口既提供所有顶级类目的查询，又提供给定类目id自身信息和子类目信息的查询，所以可以根据此接口逐层获取所有的类目信息
*/
type TaobaoSimbaInsightCatsinfoGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaInsightCatsinfoGetResponse `json:"taobao_simba_insight_catsinfo_get_response,omitempty"`
}

type TaobaoSimbaInsightCatsinfoGetResponse struct {

    // 类目详细信息
    CategoryInfoList   []InsightCategoryInfoDTO `json:"category_info_list,omitempty"`

}
