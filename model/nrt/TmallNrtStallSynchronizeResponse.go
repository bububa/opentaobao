package nrt

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
摊位信息同步 APIResponse
tmall.nrt.stall.synchronize

摊位信息同步
*/
type TmallNrtStallSynchronizeAPIResponse struct {
    model.CommonResponse
    TmallNrtStallSynchronizeResponse
}

type TmallNrtStallSynchronizeResponse struct {
    XMLName xml.Name `xml:"tmall_nrt_stall_synchronize_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    TmallNrtStallSynchronize   *ResultDO `json:"tmall_nrt_stall_synchronize,omitempty" xml:"tmall_nrt_stall_synchronize,omitempty"`

    
}
