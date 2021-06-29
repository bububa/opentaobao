package tmallnr

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
同城配物流信息回传 APIResponse
tmall.nr.fulfill.logistics.sync

同城配业务物流信息回传，通过接口将物流信息同步给天猫
*/
type TmallNrFulfillLogisticsSyncAPIResponse struct {
    model.CommonResponse
    TmallNrFulfillLogisticsSyncResponse
}

type TmallNrFulfillLogisticsSyncResponse struct {
    XMLName xml.Name `xml:"tmall_nr_fulfill_logistics_sync_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *NrResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
