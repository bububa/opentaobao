package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
派样商品门店库存查询接口 APIResponse
alibaba.wdk.bm.paiyang.stock.query

淘鲜达接入第三方进行派样，第三方查询派样商品的门店库存信息。
*/
type AlibabaWdkBmPaiyangStockQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_bm_paiyang_stock_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 请求出参
    
    Result   *BmResult `json:"result,omitempty" xml:"