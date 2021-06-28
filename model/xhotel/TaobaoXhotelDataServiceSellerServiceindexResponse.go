package xhotel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家服务指数查询 APIResponse
taobao.xhotel.data.service.seller.serviceindex

卖家服务指数查询
*/
type TaobaoXhotelDataServiceSellerServiceindexAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"xhotel_data_service_seller_serviceindex_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *TaobaoXhotelDataServiceSellerServiceindexResultSet `json:"result,omitempty" xml:"