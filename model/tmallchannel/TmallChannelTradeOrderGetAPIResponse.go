package tmallchannel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
通过主采购单号查询采购单 API返回值 
tmall.channel.trade.order.get

通过主采购单号查询采购单
*/
type TmallChannelTradeOrderGetAPIResponse struct {
    model.CommonResponse
    TmallChannelTradeOrderGetAPIResponseModel
}

// 通过主采购单号查询采购单 成功返回结果
type TmallChannelTradeOrderGetAPIResponseModel struct {
    XMLName xml.Name `xml:"tmall_channel_trade_order_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 主采购单号
    MainPurchaseOrderNo   int64 `json:"main_purchase_order_no,omitempty" xml:"main_purchase_order_no,omitempty"`
    // 申请单单号
    ChannelPurchaseApplyOrderNo   string `json:"channel_purchase_apply_order_no,omitempty" xml:"channel_purchase_apply_order_no,omitempty"`
    // 分销商淘宝nick
    DistributorNick   string `json:"distributor_nick,omitempty" xml:"distributor_nick,omitempty"`
    // 买家淘宝nick
    BuyerTaobaoNick   string `json:"buyer_taobao_nick,omitempty" xml:"buyer_taobao_nick,omitempty"`
    // 交易类型
    TradeType   int64 `json:"trade_type,omitempty" xml:"trade_type,omitempty"`
    // 付款类型
    PayType   int64 `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
    // 渠道编码
    Channel   int64 `json:"channel,omitempty" xml:"channel,omitempty"`
    // 改价后的最新邮费
    PostFee   int64 `json:"post_fee,omitempty" xml:"post_fee,omitempty"`
    // 主采购单支付状态
    PayStatus   int64 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
    // 主采购单物流状态
    LogisticsStatus   int64 `json:"logistics_status,omitempty" xml:"logistics_status,omitempty"`
    // 交易状态
    OrderStatus   int64 `json:"order_status,omitempty" xml:"order_status,omitempty"`
    // 记录创建时间
    CreateTime   string `json:"create_time,omitempty" xml:"create_time,omitempty"`
    // 记录修改时间
    ModifiedTime   string `json:"modified_time,omitempty" xml:"modified_time,omitempty"`
    // 付款时间
    PayTime   string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
    // 子采购单列表
    SubOrderList   []TopChannelSubPurchaseOrderDto `json:"sub_order_list,omitempty" xml:"sub_order_list>top_channel_sub_purchase_order_dto,omitempty"`
    // 物流单信息
    ChannelLogisticsOrder   *TopChannelLogisticsOrderDto `json:"channel_logistics_order,omitempty" xml:"channel_logistics_order,omitempty"`
    // 解析描述
    Schema   string `json:"schema,omitempty" xml:"schema,omitempty"`
    // 支付信息
    TopPurchasePayOrder   *TopPurchasePayOrderDto `json:"top_purchase_pay_order,omitempty" xml:"top_purchase_pay_order,omitempty"`
}
