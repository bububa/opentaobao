package idle

// AppraiseOrderInfoDto 
type AppraiseOrderInfoDto struct {
    // 服务商appkey
    AppKey   string `json:"app_key,omitempty" xml:"app_key,omitempty"`
    // 订单号
    BizOrderId   string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
    // 买家收货地址。主状态>=6，买家确认购买时填写
    BuyerAddress   string `json:"buyer_address,omitempty" xml:"buyer_address,omitempty"`
    // 买家关闭原因
    BuyerCloseReason   string `json:"buyer_close_reason,omitempty" xml:"buyer_close_reason,omitempty"`
    // 渠道
    Channel   string `json:"channel,omitempty" xml:"channel,omitempty"`
    // 预留 JSON 格式渠道业务数据
    ChannelData   string `json:"channel_data,omitempty" xml:"channel_data,omitempty"`
    // 订单环境 'online'：线上环境；'pre'：测试环境
    Env   string `json:"env,omitempty" xml:"env,omitempty"`
    // 订单创建时间
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    // 订单主状态。(主状态,子状态,状态说明)示例如下： ("1", "1", "买家拍下未付款") ("2", "1", "买家拍下已付款") ("3", "1", "卖家已发货") 等。详情参考对接文档
    OrderStatus   string `json:"order_status,omitempty" xml:"order_status,omitempty"`
    // 卖家关闭原因
    SellerCloseReason   string `json:"seller_close_reason,omitempty" xml:"seller_close_reason,omitempty"`
    // 卖家手机号码。逆向流程中，主状态>=102卖家已发货后发起逆向时有值
    SellerPhone   string `json:"seller_phone,omitempty" xml:"seller_phone,omitempty"`
    // 逆向流程卖家收货地址。逆向流程中，主状态>=102卖家已发货后发起逆向时有值
    SellerReceiptAddress   string `json:"seller_receipt_address,omitempty" xml:"seller_receipt_address,omitempty"`
    // 卖家发货给验货中心的单号。主状态>=3，卖家已发货时有值
    Seller2acMailNo   string `json:"seller2ac_mail_no,omitempty" xml:"seller2ac_mail_no,omitempty"`
    // 卖家收货姓名，逆向退货用。逆向流程中，主状态>=102卖家已发货后发起逆向时有值
    SellerReceiptName   string `json:"seller_receipt_name,omitempty" xml:"seller_receipt_name,omitempty"`
    // 鉴定商发货给买家的单号 主状态>=7，验货中心已发货时填写
    Ac2buyerMailNo   string `json:"ac2buyer_mail_no,omitempty" xml:"ac2buyer_mail_no,omitempty"`
    // 逆向流程鉴定商发货给卖家的单号 逆向流程中，主状态>=103卖家已发货后发起逆向时填写
    Ac2sellerMailNo   string `json:"ac2seller_mail_no,omitempty" xml:"ac2seller_mail_no,omitempty"`
    // 鉴定场景：1表示新品鉴定，2表示旧品鉴定
    IdleAppraiseScene   string `json:"idle_appraise_scene,omitempty" xml:"idle_appraise_scene,omitempty"`
    // spu信息jsonStr
    SpuInfo   string `json:"spu_info,omitempty" xml:"spu_info,omitempty"`
    // 买家支付宝ID，用于卖家责任时赔付买家 逆向流程中，主状态105鉴定为赝品时填写
    BuyerAlipayUserId   string `json:"buyer_alipay_user_id,omitempty" xml:"buyer_alipay_user_id,omitempty"`
    // 订单子状态，说明见order_status
    OrderSubStatus   string `json:"order_sub_status,omitempty" xml:"order_sub_status,omitempty"`
    // 买家手机号。主状态>=6，买家确认购买时有值
    BuyerPhone   string `json:"buyer_phone,omitempty" xml:"buyer_phone,omitempty"`
    // 买家收货姓名。主状态>=6，买家确认购买时有值
    BuyerReceiptName   string `json:"buyer_receipt_name,omitempty" xml:"buyer_receipt_name,omitempty"`
    // 需付给买家的赔付金额，单位分。状态为：已付款后发货超时（101-1 101-2）、卖家取消订单（101-3 101-4）、鉴定为赝品（主状态105）时有值。
    Compensation2buyer   int64 `json:"compensation2buyer,omitempty" xml:"compensation2buyer,omitempty"`
    // 二手卖家承诺验货项。数组，每条记录的keyId代表验货项id，比如1001可能代表内存。valueId和valueName分别代表验货项值的id和描述。keyId、valueId的取值参考对接文档
    IdleAppraiseCheckpoints   string `json:"idle_appraise_checkpoints,omitempty" xml:"idle_appraise_checkpoints,omitempty"`
    // 品牌Id
    BrandId   string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
    // spuId
    SpuId   string `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
    // 服务商应收实时分账金额，分。仅当交易成功或者卖家无责买家不买 这两个状态有效，其余为0。当保证金被罚没的状态时，不会实时分账而是线下结算，所以只有上述两个状态有值。
    SupplierChargeFeeCent   int64 `json:"supplier_charge_fee_cent,omitempty" xml:"supplier_charge_fee_cent,omitempty"`
    // 商品详情页
    ItemDetailInfo   string `json:"item_detail_info,omitempty" xml:"item_detail_info,omitempty"`
    // 买家村庄
    BuyerCountry   string `json:"buyer_country,omitempty" xml:"buyer_country,omitempty"`
    // 买家区
    BuyerArea   string `json:"buyer_area,omitempty" xml:"buyer_area,omitempty"`
    // 买家城市
    BuyerCity   string `json:"buyer_city,omitempty" xml:"buyer_city,omitempty"`
    // 买家省份
    BuyerProvince   string `json:"buyer_province,omitempty" xml:"buyer_province,omitempty"`
    // 卖家村庄
    SellerCountry   string `json:"seller_country,omitempty" xml:"seller_country,omitempty"`
    // 卖家区
    SellerArea   string `json:"seller_area,omitempty" xml:"seller_area,omitempty"`
    // 卖家城市
    SellerCity   string `json:"seller_city,omitempty" xml:"seller_city,omitempty"`
    // 卖家省份
    SellerProvince   string `json:"seller_province,omitempty" xml:"seller_province,omitempty"`
    // 是否是经主发发布/编辑的验货宝商品
    AppraiseFromNewPublisher   bool `json:"appraise_from_new_publisher,omitempty" xml:"appraise_from_new_publisher,omitempty"`
    // 类目聚合场景,如"YHB_3C"
    CateAggScene   string `json:"cate_agg_scene,omitempty" xml:"cate_agg_scene,omitempty"`
}
