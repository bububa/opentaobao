package nrt

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
同步摊位收银比例 APIResponse
tmall.nrt.stall.payratio.synchronize

ISV同步摊位收银比例到阿里
*/
type TmallNrtStallPayratioSynchronizeAPIResponse struct {
    model.CommonResponse
    TmallNrtStallPayratioSynchronizeResponse
}

type TmallNrtStallPayratioSynchronizeResponse struct {
    XMLName xml.Name `xml:"tmall_nrt_stall_payratio_synchronize_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

}
