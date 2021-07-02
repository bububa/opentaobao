package tmallnr

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallNrSoldOrderlistQueryJstAPIResponse 已入塔商家查询订单列表 API返回值
// tmall.nr.sold.orderlist.query.jst
//
// 该服务用于已入聚石塔的商家，获取订单列表信息；
type TmallNrSoldOrderlistQueryJstAPIResponse struct {
	model.CommonResponse
	TmallNrSoldOrderlistQueryJstAPIResponseModel
}

// TmallNrSoldOrderlistQueryJstAPIResponseModel is 已入塔商家查询订单列表 成功返回结果
type TmallNrSoldOrderlistQueryJstAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nr_sold_orderlist_query_jst_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *NewRetailResult `json:"result,omitempty" xml:"result,omitempty"`
}
