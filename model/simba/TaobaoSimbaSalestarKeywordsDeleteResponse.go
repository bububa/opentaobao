package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
销量明星关键词删除 APIResponse
taobao.simba.salestar.keywords.delete

（新）关键词删除相关接口
*/
type TaobaoSimbaSalestarKeywordsDeleteAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"simba_salestar_keywords_delete_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 删除成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"