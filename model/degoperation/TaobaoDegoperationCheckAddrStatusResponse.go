package degoperation

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
地址 APIResponse
taobao.degoperation.check.addr.status

激励
*/
type TaobaoDegoperationCheckAddrStatusAPIResponse struct {
    model.CommonResponse
    TaobaoDegoperationCheckAddrStatusResponse
}

type TaobaoDegoperationCheckAddrStatusResponse struct {
    XMLName xml.Name `xml:"degoperation_check_addr_status_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *BonusResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
