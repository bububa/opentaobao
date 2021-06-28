package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
单品搜索人群修改状态 APIResponse
taobao.simba.serchcrowd.state.batch.update

暂停或启用单品推广搜索人群溢价
*/
type TaobaoSimbaSerchcrowdStateBatchUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"simba_serchcrowd_state_batch_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 部分失败时返回错误List
    
    ErrorList   []string `json:"error_list,omitempty" xml:"