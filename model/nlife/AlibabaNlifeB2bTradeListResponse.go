package nlife

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取企业下的采购单列表 API返回值 
alibaba.nlife.b2b.trade.list

获取指定门店下的采购单列表
*/
type AlibabaNlifeB2bTradeListAPIResponse struct {
    model.CommonResponse
    AlibabaNlifeB2bTradeListResponse
}

// 获取企业下的采购单列表 成功返回结果
type AlibabaNlifeB2bTradeListResponse struct {
    XMLName xml.Name `xml:"alibaba_nlife_b2b_trade_list_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 查询结果数据
    Data   *ProcurementResponseDO `json:"data,omitempty" xml:"data,omitempty"`
}
