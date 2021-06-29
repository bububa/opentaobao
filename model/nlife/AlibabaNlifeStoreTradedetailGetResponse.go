package nlife

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询采购单详情信息 API返回值 
alibaba.nlife.store.tradedetail.get

根据集团id和采购单号，查询采购单的详情信息
*/
type AlibabaNlifeStoreTradedetailGetAPIResponse struct {
    model.CommonResponse
    AlibabaNlifeStoreTradedetailGetResponse
}

// 查询采购单详情信息 成功返回结果
type AlibabaNlifeStoreTradedetailGetResponse struct {
    XMLName xml.Name `xml:"alibaba_nlife_store_tradedetail_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // data
    Data   *ProcurementDetailResponseDo `json:"data,omitempty" xml:"data,omitempty"`
}
