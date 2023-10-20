package fenxiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoDealerRequisitionorderQueryAPIResponse 按编号查询采购申请/经销采购单 API返回值
// taobao.fenxiao.dealer.requisitionorder.query
//
// 按编号查询采购申请/经销采购单，目前支持供应商和分销商查询。
type TaobaoFenxiaoDealerRequisitionorderQueryAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoDealerRequisitionorderQueryAPIResponseModel
}

// TaobaoFenxiaoDealerRequisitionorderQueryAPIResponseModel is 按编号查询采购申请/经销采购单 成功返回结果
type TaobaoFenxiaoDealerRequisitionorderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_dealer_requisitionorder_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 经销采购单结果列表
	DealerOrders []DealerOrder `json:"dealer_orders,omitempty" xml:"dealer_orders>dealer_order,omitempty"`
}
