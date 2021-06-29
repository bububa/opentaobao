package traderate

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
评价大家印象印象短语接口 APIResponse
taobao.traderate.impr.imprwords.get

根据淘宝后台类目的一级类目和叶子类目
*/
type TaobaoTraderateImprImprwordsGetAPIResponse struct {
    model.CommonResponse
    TaobaoTraderateImprImprwordsGetResponse
}

type TaobaoTraderateImprImprwordsGetResponse struct {
    XMLName xml.Name `xml:"traderate_impr_imprwords_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回类目下所有大家印象的标签
    
    ImprWords   []string `json:"impr_words,omitempty" xml:"impr_words>string,omitempty"`
    
    
}
