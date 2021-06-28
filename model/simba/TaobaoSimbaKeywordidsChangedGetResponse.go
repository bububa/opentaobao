package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取修改的词ID APIResponse
taobao.simba.keywordids.changed.get

获取修改的词ID
*/
type TaobaoSimbaKeywordidsChangedGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"simba_keywordids_changed_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 词的ID列表
    
    ChangedKeywordIds   []int64 `json:"changed_keyword_ids,omitempty" xml:"