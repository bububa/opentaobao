package nlife

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询采购单详情信息 APIResponse
alibaba.nlife.store.tradedetail.get

根据集团id和采购单号，查询采购单的详情信息
*/
type AlibabaNlifeStoreTradedetailGetAPIResponse struct {
    model.CommonResponse
    AlibabaNlifeStoreTradedetailGetResponse
}

type AlibabaNlifeStoreTradedetailGetResponse struct {
    XMLName xml.Name `xml:"alibaba_nlife_store_tradedetail_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // data
    
    Data   *ProcurementDetailResponseDo `json:"data,omitempty" xml:"data,omitempty"`

    
}
