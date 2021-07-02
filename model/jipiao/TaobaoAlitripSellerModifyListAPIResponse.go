package jipiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripSellerModifyListAPIResponse 【机票代理商订单】改签订单列表 API返回值
// taobao.alitrip.seller.modify.list
//
// 提供机票代理商查询改签订单列表
type TaobaoAlitripSellerModifyListAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripSellerModifyListAPIResponseModel
}

// TaobaoAlitripSellerModifyListAPIResponseModel is 【机票代理商订单】改签订单列表 成功返回结果
type TaobaoAlitripSellerModifyListAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_seller_modify_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 改签订单列表
	OrderList []SyncOrderDo `json:"order_list,omitempty" xml:"order_list>sync_order_do,omitempty"`
	// 查出总记录数
	TotalItem int64 `json:"total_item,omitempty" xml:"total_item,omitempty"`
}
