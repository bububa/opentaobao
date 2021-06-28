package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
取得一个推广组的推荐关键词列表 APIResponse
taobao.simba.keywords.recommend.get

取得一个推广组的推荐关键词列表
*/
type TaobaoSimbaKeywordsRecommendGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"simba_keywords_recommend_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 推荐词分页对象，当输入的页码大于最大数值时，将返回最大的page_no值，并且结果中的数据列表为空值
    
    RecommendWords   *RecommendWordPage `json:"recommend_words,omitempty" xml:"