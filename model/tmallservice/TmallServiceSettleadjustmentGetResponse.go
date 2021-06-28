package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询结算调整单单条记录 APIResponse
tmall.service.settleadjustment.get

提供给服务商通过结算调整单id获取结算调整单信息
*/
type TmallServiceSettleadjustmentGetAPIResponse struct {
    model.CommonResponse
    TmallServiceSettleadjustmentGetResponse
}

type TmallServiceSettleadjustmentGetResponse struct {
    XMLName xml.Name `xml:"tmall_service_settleadjustment_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *TmallServiceSettleadjustmentGetResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
