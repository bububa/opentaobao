package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
销量明星关键词删除 API返回值 
taobao.simba.salestar.keywords.delete

（新）关键词删除相关接口
*/
type TaobaoSimbaSalestarKeywordsDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaSalestarKeywordsDeleteAPIResponseModel
}

// 销量明星关键词删除 成功返回结果
type TaobaoSimbaSalestarKeywordsDeleteAPIResponseModel struct {
    XMLName xml.Name `xml:"simba_salestar_keywords_delete_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 删除成功
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
    // 成功删除条数
    Results   int64 `json:"results,omitempty" xml:"results,omitempty"`
}
