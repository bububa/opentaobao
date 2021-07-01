package waybill

// WaybillCloudPrintApplyNewRequest 
type WaybillCloudPrintApplyNewRequest struct {
    // <a href="http://open.taobao.com/doc2/detail.htm?spm=a219a.7629140.0.0.8cf9Nj&treeId=17&articleId=105085&docType=1#1">物流公司Code</a>，长度小于20
    CpCode   string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
    // 目前已经不推荐使用此字段，请不要使用
    ProductCode   string `json:"product_code,omitempty" xml:"product_code,omitempty"`
    // 发货人信息
    Sender   *UserInfoDto `json:"sender,omitempty" xml:"sender,omitempty"`
    // 请求面单信息，数量限制为10
    TradeOrderInfoDtos   []TradeOrderInfoDto `json:"trade_order_info_dtos,omitempty" xml:"trade_order_info_dtos>trade_order_info_dto,omitempty"`
    // 仓code， 仓库WMS系统对接落地配业务，其它场景请不要使用
    StoreCode   string `json:"store_code,omitempty" xml:"store_code,omitempty"`
    // 配送资源code， 仓库WMS系统对接落地配业务，其它场景请不要使用
    ResourceCode   string `json:"resource_code,omitempty" xml:"resource_code,omitempty"`
    // 是否使用智分宝预分拣， 仓库WMS系统对接落地配业务，其它场景请不要使用
    DmsSorting   bool `json:"dms_sorting,omitempty" xml:"dms_sorting,omitempty"`
    // 订单上是否带3PLtiming属性, 该属性需要严格与订单上属性保持一致，如果不确定，请使用默认false。
    ThreePlTiming   bool `json:"three_pl_timing,omitempty" xml:"three_pl_timing,omitempty"`
    // 设定取号返回的云打印报文是否加密
    NeedEncrypt   bool `json:"need_encrypt,omitempty" xml:"need_encrypt,omitempty"`
    // 快递公司支持一票多件，快运公司子母件请勿使用该参数
    MultiPackagesShipment   bool `json:"multi_packages_shipment,omitempty" xml:"multi_packages_shipment,omitempty"`
}
