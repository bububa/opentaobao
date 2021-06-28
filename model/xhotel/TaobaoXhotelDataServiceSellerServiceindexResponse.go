package xhotel

import (
    "github.com/bububa/opentaobao/model"
)

/* 
卖家服务指数查询 APIResponse
taobao.xhotel.data.service.seller.serviceindex

卖家服务指数查询
*/
type TaobaoXhotelDataServiceSellerServiceindexAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoXhotelDataServiceSellerServiceindexResponse `json:"xhotel_data_service_seller_serviceindex_response,omitempty"` 
    TaobaoXhotelDataServiceSellerServiceindexResponse
}

/* model for simplify = false
type TaobaoXhotelDataServiceSellerServiceindexResponse struct {

    // result
    
    Result  *struct {
        TaobaoXhotelDataServiceSellerServiceindexResultSet  *TaobaoXhotelDataServiceSellerServiceindexResultSet `json:"taobao_xhotel_data_service_seller_serviceindex_result_set,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoXhotelDataServiceSellerServiceindexResponse struct {

    // result
    Result   *TaobaoXhotelDataServiceSellerServiceindexResultSet `json:"result,omitempty"`

}
