package servicecenter

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商15分钟获取数据 APIResponse
tmall.service.settleadjustment.search

天猫服务平台，按修改时间，时间间隔在15中内（包含15分钟），获取调整单数据
*/
type TmallServiceSettleadjustmentSearchAPIResponse struct {
    model.CommonResponse
    TmallServiceSettleadjustmentSearchResponse
}

type TmallServiceSettleadjustmentSearchResponse struct {
    XMLName xml.Name `xml:"tmall_service_settleadjustment_search_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *TmallServiceSettleadjustmentSearchResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
