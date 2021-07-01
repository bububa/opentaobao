package tmallnr

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
同城配物流信息回传 API返回值 
tmall.nr.fulfill.logistics.sync

同城配业务物流信息回传，通过接口将物流信息同步给天猫
*/
type TmallNrFulfillLogisticsSyncAPIResponse struct {
    model.CommonResponse
    TmallNrFulfillLogisticsSyncAPIResponseModel
}

// 同城配物流信息回传 成功返回结果
type TmallNrFulfillLogisticsSyncAPIResponseModel struct {
    XMLName xml.Name `xml:"tmall_nr_fulfill_logistics_sync_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *NrResult `json:"result,omitempty" xml:"result,omitempty"`
}
