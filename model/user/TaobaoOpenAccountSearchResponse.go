package user

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
open account数据搜索 API返回值 
taobao.open.account.search

open account数据搜索
*/
type TaobaoOpenAccountSearchAPIResponse struct {
    model.CommonResponse
    TaobaoOpenAccountSearchResponse
}

// open account数据搜索 成功返回结果
type TaobaoOpenAccountSearchResponse struct {
    XMLName xml.Name `xml:"open_account_search_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Data   *OpenAccountSearchResult `json:"data,omitempty" xml:"data,omitempty"`
}
