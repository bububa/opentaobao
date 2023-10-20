package wlb

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlborderitempagegetAPIResponse 分页查询物流宝订单商品详情 API返回值
// taobao.wlb.orderitem.page.get
//
// 分页查询物流宝订单商品详情
type TaobaowlborderitempagegetAPIResponse struct {
	model.CommonResponse
	TaobaowlborderitempagegetAPIResponseModel
}

// TaobaowlborderitempagegetAPIResponseModel is 分页查询物流宝订单商品详情 成功返回结果
type TaobaowlborderitempagegetAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_orderitem_page_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	OrderItemList []WlbOrderItem `json:"order_item_list,omitempty" xml:"order_item_list>wlb_order_item,omitempty"`
	// 总数量
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
