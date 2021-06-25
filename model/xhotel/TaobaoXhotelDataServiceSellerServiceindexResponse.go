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
    Response *TaobaoXhotelDataServiceSellerServiceindexResponse `json:"taobao_xhotel_data_service_seller_serviceindex_response,omitempty"`
}

type TaobaoXhotelDataServiceSellerServiceindexResponse struct {

    // result
    Result   *TaobaoXhotelDataServiceSellerServiceindexResultSet `json:"result,omitempty"`

}
