package drug

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取订单详情 APIResponse
alibaba.alihealth.nr.trade.order.get

阿里健康O2O，获取订单详情
*/
type AlibabaAlihealthNrTradeOrderGetAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthNrTradeOrderGetResponse
}

type AlibabaAlihealthNrTradeOrderGetResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_nr_trade_order_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 淘宝订单id
    
    OrderId   int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`

    
    // 订单状态：1-等待买家付款；2-等待商家接单；4-退款中；12-商家配送中；20-订单关闭；22-订单关闭；21-交易完成；
    
    OrderStatus   int64 `json:"order_status,omitempty" xml:"order_status,omitempty"`

    
    // 单店单日订单唯一序号
    
    SerialNo   string `json:"serial_no,omitempty" xml:"serial_no,omitempty"`

    
    // 订单创建时间
    
    CreateTime   int64 `json:"create_time,omitempty" xml:"create_time,omitempty"`

    
    // 订单支付时间
    
    PayTime   int64 `json:"pay_time,omitempty" xml:"pay_time,omitempty"`

    
    // 店铺shopId
    
    ShopId   int64 `json:"shop_id,omitempty" xml:"shop_id,omitempty"`

    
    // 店铺storeId
    
    StoreId   string `json:"store_id,omitempty" xml:"store_id,omitempty"`

    
    // 店铺名
    
    StoreName   string `json:"store_name,omitempty" xml:"store_name,omitempty"`

    
    // 店铺经纬度
    
    StoreAddressPoi   string `json:"store_address_poi,omitempty" xml:"store_address_poi,omitempty"`

    
    // 订单总金额
    
    TotalFee   int64 `json:"total_fee,omitempty" xml:"total_fee,omitempty"`

    
    // 实收金额
    
    ActualReceiveFee   int64 `json:"actual_receive_fee,omitempty" xml:"actual_receive_fee,omitempty"`

    
    // 实付金额
    
    ActualPayFee   int64 `json:"actual_pay_fee,omitempty" xml:"actual_pay_fee,omitempty"`

    
    // 订单优惠列表
    
    PromotionList   []OrderPromotionDto `json:"promotion_list,omitempty" xml:"promotion_list>order_promotion_dto,omitempty"`
    
    
    // 总优惠信息
    
    TotalPromotion   int64 `json:"total_promotion,omitempty" xml:"total_promotion,omitempty"`

    
    // 买家收货地址上的名字
    
    BuyerName   string `json:"buyer_name,omitempty" xml:"buyer_name,omitempty"`

    
    // 买家收货地址上的电话
    
    BuyerPhone   string `json:"buyer_phone,omitempty" xml:"buyer_phone,omitempty"`

    
    // 买家留言
    
    BuyerNote   string `json:"buyer_note,omitempty" xml:"buyer_note,omitempty"`

    
    // 发票抬头
    
    InvoiceTitle   string `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`

    
    // 买家地址
    
    BuyerAddress   string `json:"buyer_address,omitempty" xml:"buyer_address,omitempty"`

    
    // 买家收货地址经纬度
    
    BuyerAddressPoi   string `json:"buyer_address_poi,omitempty" xml:"buyer_address_poi,omitempty"`

    
    // 该店铺的老顾客
    
    IsFrequenter   int64 `json:"is_frequenter,omitempty" xml:"is_frequenter,omitempty"`

    
    // 商家设置的运费
    
    Carriage   int64 `json:"carriage,omitempty" xml:"carriage,omitempty"`

    
    // 子订单
    
    SubOrderList   []SubOrderDto `json:"sub_order_list,omitempty" xml:"sub_order_list>sub_order_dto,omitempty"`
    
    
    // 处方药标示
    
    Rx   int64 `json:"rx,omitempty" xml:"rx,omitempty"`

    
    // 处方药回拨电话
    
    RxPhone   string `json:"rx_phone,omitempty" xml:"rx_phone,omitempty"`

    
    // 处方图片
    
    RxPicList   []string `json:"rx_pic_list,omitempty" xml:"rx_pic_list>string,omitempty"`
    
    
    // 自提标识
    
    DrugTake   int64 `json:"drug_take,omitempty" xml:"drug_take,omitempty"`

    
    // 收货地址-省
    
    BuyerAddressProvince   string `json:"buyer_address_province,omitempty" xml:"buyer_address_province,omitempty"`

    
    // 收货地址-市
    
    BuyerAddressCity   string `json:"buyer_address_city,omitempty" xml:"buyer_address_city,omitempty"`

    
    // 收货地址-区
    
    BuyerAddressDistrict   string `json:"buyer_address_district,omitempty" xml:"buyer_address_district,omitempty"`

    
    // 预计送达时间
    
    ExpectedDeliveryTime   string `json:"expected_delivery_time,omitempty" xml:"expected_delivery_time,omitempty"`

    
    // 1单轨、2双规、3医保支付、4医保自费
    
    RxType   string `json:"rx_type,omitempty" xml:"rx_type,omitempty"`

    
    // 高德、本地商户
    
    OutStoreId   string `json:"out_store_id,omitempty" xml:"out_store_id,omitempty"`

    
}
