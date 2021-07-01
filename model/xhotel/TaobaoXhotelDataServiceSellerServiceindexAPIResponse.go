package xhotel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelDataServiceSellerServiceindexAPIResponse
卖家服务指数查询 API返回值
taobao.xhotel.data.service.seller.serviceindex

卖家服务指数查询 */
type TaobaoXhotelDataServiceSellerServiceindexAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelDataServiceSellerServiceindexAPIResponseModel
}

// TaobaoXhotelDataServiceSellerServiceindexAPIResponseModel is 卖家服务指数查询 成功返回结果
type TaobaoXhotelDataServiceSellerServiceindexAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_data_service_seller_serviceindex_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoXhotelDataServiceSellerServiceindexResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
