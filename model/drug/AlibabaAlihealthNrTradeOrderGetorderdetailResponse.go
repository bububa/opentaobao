package drug

import (
    "github.com/bububa/opentaobao/model"
)

/* 
根据订单id获取单条订单详情 APIResponse
alibaba.alihealth.nr.trade.order.getorderdetail

阿里健康O2O，获取订单详情，修复组合商品价格精度问题
*/
type AlibabaAlihealthNrTradeOrderGetorderdetailAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlihealthNrTradeOrderGetorderdetailResponse `json:"alibaba_alihealth_nr_trade_order_getorderdetail_response,omitempty"` 
    AlibabaAlihealthNrTradeOrderGetorderdetailResponse
}

/* model for simplify = false
type AlibabaAlihealthNrTradeOrderGetorderdetailResponse struct {

    // 高德，本地商户
    
    OutStoreId   string `json:"out_store_id,omitempty"`
    

    // 1单轨、2双规、3医保支付、4医保自费
    
    RxType   string `json:"rx_type,omitempty"`
    

    // 预计送达时间
    
    ExpectedDeliveryTime   string `json:"expected_delivery_time,omitempty"`
    

    // 收货地址-区
    
    BuyerAddressDistrict   string `json:"buyer_address_district,omitempty"`
    

    // 收货地址-市
    
    BuyerAddressCity   string `json:"buyer_address_city,omitempty"`
    

    // 收货地址-省
    
    BuyerAddressProvince   string `json:"buyer_address_province,omitempty"`
    

    // 自提标识
    
    DrugTake   int64 `json:"drug_take,omitempty"`
    

    // 处方图片
    
    RxPicList  struct {
        String  []string `json:"string,omitempty"`
    } `json:"rx_pic_list,omitempty"`
    

    // 处方药回拨电话
    
    RxPhone   string `json:"rx_phone,omitempty"`
    

    // 处方药标示
    
    Rx   int64 `json:"rx,omitempty"`
    

    // 子订单
    
    SubOrderList  struct {
        SubOrderDto  []SubOrderDto `json:"sub_order_dto,omitempty"`
    } `json:"sub_order_list,omitempty"`
    

    // 商家设置的运费
    
    Carriage   int64 `json:"carriage,omitempty"`
    

    // 该店铺的老顾客
    
    IsFrequenter   int64 `json:"is_frequenter,omitempty"`
    

    // 买家收货地址经纬度
    
    BuyerAddressPoi   string `json:"buyer_address_poi,omitempty"`
    

    // 买家地址
    
    BuyerAddress   string `json:"buyer_address,omitempty"`
    

    // 订单优惠列表
    
    PromotionList  struct {
        OrderPromotionDto  []OrderPromotionDto `json:"order_promotion_dto,omitempty"`
    } `json:"promotion_list,omitempty"`
    

    // 发票抬头
    
    InvoiceTitle   string `json:"invoice_title,omitempty"`
    

    // 买家留言
    
    BuyerNote   string `json:"buyer_note,omitempty"`
    

    // 买家收货地址上的电话
    
    BuyerPhone   string `json:"buyer_phone,omitempty"`
    

    // 买家收货地址上的名字
    
    BuyerName   string `json:"buyer_name,omitempty"`
    

    // 总优惠信息
    
    TotalPromotion   int64 `json:"total_promotion,omitempty"`
    

    // 实付金额
    
    ActualPayFee   int64 `json:"actual_pay_fee,omitempty"`
    

    // 实收金额
    
    ActualReceiveFee   int64 `json:"actual_receive_fee,omitempty"`
    

    // 订单总金额
    
    TotalFee   int64 `json:"total_fee,omitempty"`
    

    // 店铺经纬度
    
    StoreAddressPoi   string `json:"store_address_poi,omitempty"`
    

    // 店铺名
    
    StoreName   string `json:"store_name,omitempty"`
    

    // abc2123sdfc
    
    StoreId   string `json:"store_id,omitempty"`
    

    // 店铺shopId
    
    ShopId   int64 `json:"shop_id,omitempty"`
    

    // 订单支付时间
    
    PayTime   int64 `json:"pay_time,omitempty"`
    

    // 订单创建时间
    
    CreateTime   int64 `json:"create_time,omitempty"`
    

    // 单店单日订单唯一序号
    
    SerialNo   string `json:"serial_no,omitempty"`
    

    // 订单状态：1-等待买家付款；2-等待商家接单；4-退款中；12-商家配送中；20-订单关闭；22-订单关闭；21-交易完成；
    
    OrderStatus   int64 `json:"order_status,omitempty"`
    

    // 淘宝订单id
    
    OrderId   int64 `json:"order_id,omitempty"`
    

}
*/

type AlibabaAlihealthNrTradeOrderGetorderdetailResponse struct {

    // 高德，本地商户
    OutStoreId   string `json:"out_store_id,omitempty"`

    // 1单轨、2双规、3医保支付、4医保自费
    RxType   string `json:"rx_type,omitempty"`

    // 预计送达时间
    ExpectedDeliveryTime   string `json:"expected_delivery_time,omitempty"`

    // 收货地址-区
    BuyerAddressDistrict   string `json:"buyer_address_district,omitempty"`

    // 收货地址-市
    BuyerAddressCity   string `json:"buyer_address_city,omitempty"`

    // 收货地址-省
    BuyerAddressProvince   string `json:"buyer_address_province,omitempty"`

    // 自提标识
    DrugTake   int64 `json:"drug_take,omitempty"`

    // 处方图片
    RxPicList   []string `json:"rx_pic_list,omitempty"`

    // 处方药回拨电话
    RxPhone   string `json:"rx_phone,omitempty"`

    // 处方药标示
    Rx   int64 `json:"rx,omitempty"`

    // 子订单
    SubOrderList   []SubOrderDto `json:"sub_order_list,omitempty"`

    // 商家设置的运费
    Carriage   int64 `json:"carriage,omitempty"`

    // 该店铺的老顾客
    IsFrequenter   int64 `json:"is_frequenter,omitempty"`

    // 买家收货地址经纬度
    BuyerAddressPoi   string `json:"buyer_address_poi,omitempty"`

    // 买家地址
    BuyerAddress   string `json:"buyer_address,omitempty"`

    // 订单优惠列表
    PromotionList   []OrderPromotionDto `json:"promotion_list,omitempty"`

    // 发票抬头
    InvoiceTitle   string `json:"invoice_title,omitempty"`

    // 买家留言
    BuyerNote   string `json:"buyer_note,omitempty"`

    // 买家收货地址上的电话
    BuyerPhone   string `json:"buyer_phone,omitempty"`

    // 买家收货地址上的名字
    BuyerName   string `json:"buyer_name,omitempty"`

    // 总优惠信息
    TotalPromotion   int64 `json:"total_promotion,omitempty"`

    // 实付金额
    ActualPayFee   int64 `json:"actual_pay_fee,omitempty"`

    // 实收金额
    ActualReceiveFee   int64 `json:"actual_receive_fee,omitempty"`

    // 订单总金额
    TotalFee   int64 `json:"total_fee,omitempty"`

    // 店铺经纬度
    StoreAddressPoi   string `json:"store_address_poi,omitempty"`

    // 店铺名
    StoreName   string `json:"store_name,omitempty"`

    // abc2123sdfc
    StoreId   string `json:"store_id,omitempty"`

    // 店铺shopId
    ShopId   int64 `json:"shop_id,omitempty"`

    // 订单支付时间
    PayTime   int64 `json:"pay_time,omitempty"`

    // 订单创建时间
    CreateTime   int64 `json:"create_time,omitempty"`

    // 单店单日订单唯一序号
    SerialNo   string `json:"serial_no,omitempty"`

    // 订单状态：1-等待买家付款；2-等待商家接单；4-退款中；12-商家配送中；20-订单关闭；22-订单关闭；21-交易完成；
    OrderStatus   int64 `json:"order_status,omitempty"`

    // 淘宝订单id
    OrderId   int64 `json:"order_id,omitempty"`

}
