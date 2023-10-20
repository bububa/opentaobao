package tmallchannel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallchanneltradeordergetsAPIResponse 分页查询采购单 API返回值
// tmall.channel.trade.order.gets
//
// 分页查询采购单
type TmallchanneltradeordergetsAPIResponse struct {
	model.CommonResponse
	TmallchanneltradeordergetsAPIResponseModel
}

// TmallchanneltradeordergetsAPIResponseModel is 分页查询采购单 成功返回结果
type TmallchanneltradeordergetsAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_channel_trade_order_gets_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 此页中包含的元素
	PageElements []TopChannelPurchaseOrderDto `json:"page_elements,omitempty" xml:"page_elements>top_channel_purchase_order_dto,omitempty"`
	// 每页显示数量
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 查询第几页
	PageNumber int64 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// 所有元素个数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
