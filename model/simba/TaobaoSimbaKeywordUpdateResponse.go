package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
（新）关键词更新相关接口 APIResponse
taobao.simba.keyword.update

（新）关键词更新相关接口
*/
type TaobaoSimbaKeywordUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"simba_keyword_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 整体的返回值
    
    Results   []SiriusBidwordDto `json:"results,omitempty" xml:"