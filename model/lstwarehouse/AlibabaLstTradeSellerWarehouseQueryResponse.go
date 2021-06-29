package lstwarehouse

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商-本云商家-仓库查询接口 APIResponse
alibaba.lst.trade.seller.warehouse.query

查询本地云仓商家的仓库
*/
type AlibabaLstTradeSellerWarehouseQueryAPIResponse struct {
    model.CommonResponse
    AlibabaLstTradeSellerWarehouseQueryResponse
}

type AlibabaLstTradeSellerWarehouseQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_lst_trade_seller_warehouse_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaLstTradeSellerWarehouseQueryResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
