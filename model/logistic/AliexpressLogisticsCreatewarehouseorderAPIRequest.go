package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressLogisticsCreatewarehouseorderAPIRequest 创建线上物流订单 API请求
// aliexpress.logistics.createwarehouseorder
//
// 创建线上发货物流订单
type AliexpressLogisticsCreatewarehouseorderAPIRequest struct {
	model.Params
	// 申报产品信息,列表类型，以json格式来表达。{productId为产品ID(必填,如为礼品,则设置为0);categoryCnDesc为申报中文名称(必填,长度1-20);categoryEnDesc为申报英文名称(必填,长度1-60);productNum产品件数(必填1-999);productDeclareAmount为产品申报金额(必填,0.01-10000.00);productWeight为产品申报重量(必填0.001-2.000);isContainsBattery为是否包含锂电池(必填0/1);scItemId为仓储发货属性代码（团购订单，仓储发货必填，物流服务为RUSTON 哈尔滨备货仓 HRB_WLB_RUSTONHEB，属性代码对应AE商品的sku属性一级，暂时没有提供接口查询属性代码，可以在仓储管理--库存管理页面查看，例如： 团购产品的sku属性White对应属性代码 40414943126）;skuValue为属性名称（团购订单，仓储发货必填，例如：White）;hsCode为产品海关编码，获取相关数据请至：http://www.customs.gov.cn/Tabid/67737/Default.aspx};isAneroidMarkup为是否含非液体化妆品（必填，填0代表不含非液体化妆品；填1代表含非液体化妆品；默认为0）;isOnlyBattery为是否含纯电池产品（必填，填0代表不含纯电池产品；填1代表含纯电池产品；默认为0）;
	_declareProductDTOs []AeopWlDeclareProductForTopDto
	// 国内快递公司名称,物流公司Id为-1时,必填
	_domesticLogisticsCompany string
	// 国内快递运单号,长度1-32
	_domesticTrackingNo string
	// 订单来源
	_tradeOrderFrom string
	// ”根据订单号获取线上发货物流方案“API获取用户选择的实际发货物流服务（物流服务key,即仓库服务名称)例如：HRB_WLB_ZTOGZ是 中俄航空 Ruston广州仓库； HRB_WLB_RUSTONHEB为哈尔滨备货仓暂不支持，该渠道请做忽略。
	_warehouseCarrierService string
	// 发票号（可空）
	_invoiceNumber string
	// ISV用户唯一标识，一般为userId,最大长度为16个字符
	_topUserKey string
	// addresses
	_addressDTOs *Addressdtos
	// 国内快递ID(物流公司是other时,ID为-1)
	_domesticLogisticsCompanyId int64
	// 包裹数量： 创建国家小包订单时非必填，创建国家快递订单时必填
	_packageNum int64
	// 交易订单号
	_tradeOrderId int64
	// 不可达处理(退回:0/销毁:1) 。详情请参考：http://bbs.seller.aliexpress.com/bbs/read.php?tid=514111
	_undeliverableDecision int64
	// 包裹保额
	_insuranceCoverage *Money
	// 是否同意升级逆向高货值保险
	_isAgreeUpgradeReverseParcelInsure bool
}

// NewAliexpressLogisticsCreatewarehouseorderRequest 初始化AliexpressLogisticsCreatewarehouseorderAPIRequest对象
func NewAliexpressLogisticsCreatewarehouseorderRequest() *AliexpressLogisticsCreatewarehouseorderAPIRequest {
	return &AliexpressLogisticsCreatewarehouseorderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressLogisticsCreatewarehouseorderAPIRequest) GetApiMethodName() string {
	return "aliexpress.logistics.createwarehouseorder"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressLogisticsCreatewarehouseorderAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetDeclareProductDTOs is DeclareProductDTOs Setter
// 申报产品信息,列表类型，以json格式来表达。{productId为产品ID(必填,如为礼品,则设置为0);categoryCnDesc为申报中文名称(必填,长度1-20);categoryEnDesc为申报英文名称(必填,长度1-60);productNum产品件数(必填1-999);productDeclareAmount为产品申报金额(必填,0.01-10000.00);productWeight为产品申报重量(必填0.001-2.000);isContainsBattery为是否包含锂电池(必填0/1);scItemId为仓储发货属性代码（团购订单，仓储发货必填，物流服务为RUSTON 哈尔滨备货仓 HRB_WLB_RUSTONHEB，属性代码对应AE商品的sku属性一级，暂时没有提供接口查询属性代码，可以在仓储管理--库存管理页面查看，例如： 团购产品的sku属性White对应属性代码 40414943126）;skuValue为属性名称（团购订单，仓储发货必填，例如：White）;hsCode为产品海关编码，获取相关数据请至：http://www.customs.gov.cn/Tabid/67737/Default.aspx};isAneroidMarkup为是否含非液体化妆品（必填，填0代表不含非液体化妆品；填1代表含非液体化妆品；默认为0）;isOnlyBattery为是否含纯电池产品（必填，填0代表不含纯电池产品；填1代表含纯电池产品；默认为0）;
func (r *AliexpressLogisticsCreatewarehouseorderAPIRequest) SetDeclareProductDTOs(_declareProductDTOs []AeopWlDeclareProductForTopDto) error {
	r._declareProductDTOs = _declareProductDTOs
	r.Set("declare_product_d_t_os", _declareProductDTOs)
	return nil
}

// GetDeclareProductDTOs DeclareProductDTOs Getter
func (r AliexpressLogisticsCreatewarehouseorderAPIRequest) GetDeclareProductDTOs() []AeopWlDeclareProductForTopDto {
	return r._declareProductDTOs
}

// SetDomesticLogisticsCompany is DomesticLogisticsCompany Setter
// 国内快递公司名称,物流公司Id为-1时,必填
func (r *AliexpressLogisticsCreatewarehouseorderAPIRequest) SetDomesticLogisticsCompany(_domesticLogisticsCompany string) error {
	r._domesticLogisticsCompany = _domesticLogisticsCompany
	r.Set("domestic_logistics_company", _domesticLogisticsCompany)
	return nil
}

// GetDomesticLogisticsCompany DomesticLogisticsCompany Getter
func (r AliexpressLogisticsCreatewarehouseorderAPIRequest) GetDomesticLogisticsCompany() string {
	return r._domesticLogisticsCompany
}

// SetDomesticTrackingNo is DomesticTrackingNo Setter
// 国内快递运单号,长度1-32
func (r *AliexpressLogisticsCreatewarehouseorderAPIRequest) SetDomesticTrackingNo(_domesticTrackingNo string) error {
	r._domesticTrackingNo = _domesticTrackingNo
	r.Set("domestic_tracking_no", _domesticTrackingNo)
	return nil
}

// GetDomesticTrackingNo DomesticTrackingNo Getter
func (r AliexpressLogisticsCreatewarehouseorderAPIRequest) GetDomesticTrackingNo() string {
	return r._domesticTrackingNo
}

// SetTradeOrderFrom is TradeOrderFrom Setter
// 订单来源
func (r *AliexpressLogisticsCreatewarehouseorderAPIRequest) SetTradeOrderFrom(_tradeOrderFrom string) error {
	r._tradeOrderFrom = _tradeOrderFrom
	r.Set("trade_order_from", _tradeOrderFrom)
	return nil
}

// GetTradeOrderFrom TradeOrderFrom Getter
func (r AliexpressLogisticsCreatewarehouseorderAPIRequest) GetTradeOrderFrom() string {
	return r._tradeOrderFrom
}

// SetWarehouseCarrierService is WarehouseCarrierService Setter
// ”根据订单号获取线上发货物流方案“API获取用户选择的实际发货物流服务（物流服务key,即仓库服务名称)例如：HRB_WLB_ZTOGZ是 中俄航空 Ruston广州仓库； HRB_WLB_RUSTONHEB为哈尔滨备货仓暂不支持，该渠道请做忽略。
func (r *AliexpressLogisticsCreatewarehouseorderAPIRequest) SetWarehouseCarrierService(_warehouseCarrierService string) error {
	r._warehouseCarrierService = _warehouseCarrierService
	r.Set("warehouse_carrier_service", _warehouseCarrierService)
	return nil
}

// GetWarehouseCarrierService WarehouseCarrierService Getter
func (r AliexpressLogisticsCreatewarehouseorderAPIRequest) GetWarehouseCarrierService() string {
	return r._warehouseCarrierService
}

// SetInvoiceNumber is InvoiceNumber Setter
// 发票号（可空）
func (r *AliexpressLogisticsCreatewarehouseorderAPIRequest) SetInvoiceNumber(_invoiceNumber string) error {
	r._invoiceNumber = _invoiceNumber
	r.Set("invoice_number", _invoiceNumber)
	return nil
}

// GetInvoiceNumber InvoiceNumber Getter
func (r AliexpressLogisticsCreatewarehouseorderAPIRequest) GetInvoiceNumber() string {
	return r._invoiceNumber
}

// SetTopUserKey is TopUserKey Setter
// ISV用户唯一标识，一般为userId,最大长度为16个字符
func (r *AliexpressLogisticsCreatewarehouseorderAPIRequest) SetTopUserKey(_topUserKey string) error {
	r._topUserKey = _topUserKey
	r.Set("top_user_key", _topUserKey)
	return nil
}

// GetTopUserKey TopUserKey Getter
func (r AliexpressLogisticsCreatewarehouseorderAPIRequest) GetTopUserKey() string {
	return r._topUserKey
}

// SetAddressDTOs is AddressDTOs Setter
// addresses
func (r *AliexpressLogisticsCreatewarehouseorderAPIRequest) SetAddressDTOs(_addressDTOs *Addressdtos) error {
	r._addressDTOs = _addressDTOs
	r.Set("address_d_t_os", _addressDTOs)
	return nil
}

// GetAddressDTOs AddressDTOs Getter
func (r AliexpressLogisticsCreatewarehouseorderAPIRequest) GetAddressDTOs() *Addressdtos {
	return r._addressDTOs
}

// SetDomesticLogisticsCompanyId is DomesticLogisticsCompanyId Setter
// 国内快递ID(物流公司是other时,ID为-1)
func (r *AliexpressLogisticsCreatewarehouseorderAPIRequest) SetDomesticLogisticsCompanyId(_domesticLogisticsCompanyId int64) error {
	r._domesticLogisticsCompanyId = _domesticLogisticsCompanyId
	r.Set("domestic_logistics_company_id", _domesticLogisticsCompanyId)
	return nil
}

// GetDomesticLogisticsCompanyId DomesticLogisticsCompanyId Getter
func (r AliexpressLogisticsCreatewarehouseorderAPIRequest) GetDomesticLogisticsCompanyId() int64 {
	return r._domesticLogisticsCompanyId
}

// SetPackageNum is PackageNum Setter
// 包裹数量： 创建国家小包订单时非必填，创建国家快递订单时必填
func (r *AliexpressLogisticsCreatewarehouseorderAPIRequest) SetPackageNum(_packageNum int64) error {
	r._packageNum = _packageNum
	r.Set("package_num", _packageNum)
	return nil
}

// GetPackageNum PackageNum Getter
func (r AliexpressLogisticsCreatewarehouseorderAPIRequest) GetPackageNum() int64 {
	return r._packageNum
}

// SetTradeOrderId is TradeOrderId Setter
// 交易订单号
func (r *AliexpressLogisticsCreatewarehouseorderAPIRequest) SetTradeOrderId(_tradeOrderId int64) error {
	r._tradeOrderId = _tradeOrderId
	r.Set("trade_order_id", _tradeOrderId)
	return nil
}

// GetTradeOrderId TradeOrderId Getter
func (r AliexpressLogisticsCreatewarehouseorderAPIRequest) GetTradeOrderId() int64 {
	return r._tradeOrderId
}

// SetUndeliverableDecision is UndeliverableDecision Setter
// 不可达处理(退回:0/销毁:1) 。详情请参考：http://bbs.seller.aliexpress.com/bbs/read.php?tid=514111
func (r *AliexpressLogisticsCreatewarehouseorderAPIRequest) SetUndeliverableDecision(_undeliverableDecision int64) error {
	r._undeliverableDecision = _undeliverableDecision
	r.Set("undeliverable_decision", _undeliverableDecision)
	return nil
}

// GetUndeliverableDecision UndeliverableDecision Getter
func (r AliexpressLogisticsCreatewarehouseorderAPIRequest) GetUndeliverableDecision() int64 {
	return r._undeliverableDecision
}

// SetInsuranceCoverage is InsuranceCoverage Setter
// 包裹保额
func (r *AliexpressLogisticsCreatewarehouseorderAPIRequest) SetInsuranceCoverage(_insuranceCoverage *Money) error {
	r._insuranceCoverage = _insuranceCoverage
	r.Set("insurance_coverage", _insuranceCoverage)
	return nil
}

// GetInsuranceCoverage InsuranceCoverage Getter
func (r AliexpressLogisticsCreatewarehouseorderAPIRequest) GetInsuranceCoverage() *Money {
	return r._insuranceCoverage
}

// SetIsAgreeUpgradeReverseParcelInsure is IsAgreeUpgradeReverseParcelInsure Setter
// 是否同意升级逆向高货值保险
func (r *AliexpressLogisticsCreatewarehouseorderAPIRequest) SetIsAgreeUpgradeReverseParcelInsure(_isAgreeUpgradeReverseParcelInsure bool) error {
	r._isAgreeUpgradeReverseParcelInsure = _isAgreeUpgradeReverseParcelInsure
	r.Set("is_agree_upgrade_reverse_parcel_insure", _isAgreeUpgradeReverseParcelInsure)
	return nil
}

// GetIsAgreeUpgradeReverseParcelInsure IsAgreeUpgradeReverseParcelInsure Getter
func (r AliexpressLogisticsCreatewarehouseorderAPIRequest) GetIsAgreeUpgradeReverseParcelInsure() bool {
	return r._isAgreeUpgradeReverseParcelInsure
}
