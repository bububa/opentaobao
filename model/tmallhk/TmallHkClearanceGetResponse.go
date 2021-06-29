package tmallhk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫国际-清关材料查询 APIResponse
tmall.hk.clearance.get

提供订单收货人身份信息查询功能。
*/
type TmallHkClearanceGetAPIResponse struct {
    model.CommonResponse
    TmallHkClearanceGetResponse
}

type TmallHkClearanceGetResponse struct {
    XMLName xml.Name `xml:"tmall_hk_clearance_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 查询结果对象
    
    Result   *CertifyQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
