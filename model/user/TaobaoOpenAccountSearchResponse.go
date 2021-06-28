package user

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
open account数据搜索 APIResponse
taobao.open.account.search

open account数据搜索
*/
type TaobaoOpenAccountSearchAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"open_account_search_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    Data   *OpenAccountSearchResult `json:"data,omitempty" xml:"