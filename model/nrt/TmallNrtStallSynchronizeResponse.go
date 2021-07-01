package nrt

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
摊位信息同步 API返回值 
tmall.nrt.stall.synchronize

摊位信息同步
*/
type TmallNrtStallSynchronizeAPIResponse struct {
    model.CommonResponse
    TmallNrtStallSynchronizeResponse
}

// 摊位信息同步 成功返回结果
type TmallNrtStallSynchronizeResponse struct {
    XMLName xml.Name `xml:"tmall_nrt_stall_synchronize_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    TmallNrtStallSynchronize   *ResultDo `json:"tmall_nrt_stall_synchronize,omitempty" xml:"tmall_nrt_stall_synchronize,omitempty"`
}
