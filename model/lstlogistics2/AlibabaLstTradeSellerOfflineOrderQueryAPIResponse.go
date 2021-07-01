package lstlogistics2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商-线下订单-查询接口 API返回值 
alibaba.lst.trade.seller.offline.order.query

供应商线下订单数据上传后查询物流状态
*/
type AlibabaLstTradeSellerOfflineOrderQueryAPIResponse struct {
    model.CommonResponse
    AlibabaLstTradeSellerOfflineOrderQueryAPIResponseModel
}

// 供应商-线下订单-查询接口 成功返回结果
type AlibabaLstTradeSellerOfflineOrderQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_lst_trade_seller_offline_order_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *AlibabaLstTradeSellerOfflineOrderQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
