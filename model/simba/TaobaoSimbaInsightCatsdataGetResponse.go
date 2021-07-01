package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取类目的大盘数据 API返回值 
taobao.simba.insight.catsdata.get

根据类目id获取类目的大盘数据，包括展现指数，点击指数，点击率，本次提供的insight相关的其它接口的都是这种情况。
*/
type TaobaoSimbaInsightCatsdataGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaInsightCatsdataGetResponse
}

// 获取类目的大盘数据 成功返回结果
type TaobaoSimbaInsightCatsdataGetResponse struct {
    XMLName xml.Name `xml:"simba_insight_catsdata_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 类目详细数据列表
    CatDataList   []InsightCategoryDataDto `json:"cat_data_list,omitempty" xml:"cat_data_list>insight_category_data_dto,omitempty"`
}
