package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceItemUpdateAPIRequest
修改商品开票信息 API请求
alibaba.einvoice.item.update

ERP通过接口将商品的开票信息同步给阿里发票平台，自动开票时将读取这些开票信息，需要联系阿里小二开通对应的权限 */
type AlibabaEinvoiceItemUpdateAPIRequest struct {
	model.Params
	// 商品的开票名称，对应发票的货物劳务名称，值DELETE时表示删除
	_invoiceName string
	// 商品id，优先级高于outerId，商品必须归属于店铺，itemId和outerId不能同时为空
	_itemId int64
	// 税收分类编码，需要精确到叶子节点，必须和taxRate同时修改或删除，值DELETE时表示删除
	_itemNo string
	// skuId，必须是itemId下的sku，填写skuId后，修改和删除sku的开票信息
	_skuId int64
	// 规格型号，值DELETE时表示删除
	_specification string
	// 税率，可选值0，3，4，5，6，10，11，13， 16，17，必须和itemNo同时修改或删除,值为DELETE时表示删除
	_taxRate string
	// 0税率标识，只有税率为0的情况才有值，0=出口零税率，1=免税，2=不征收，3=普通零税率，值为DELETE时表示删除
	_zeroRateFlag string
	// 单位，值DELETE时表示删除
	_unit string
	// 商家外部商品id，如果outerId对应了多个天猫sku，则会更新所有的sku开票信息。itemId和outerId不能同时为空
	_outerId string
	// 是否根据outerId更新所有对应sku的开票信息，true=更新，false=开票信息维护在发票平台；自动开票时，根据skuId获取outerId，再根据outerId查询开票信息。outerId不为空时必填
	_updateSku bool
}

// NewAlibabaEinvoiceItemUpdateRequest 初始化AlibabaEinvoiceItemUpdateAPIRequest对象
func NewAlibabaEinvoiceItemUpdateRequest() *AlibabaEinvoiceItemUpdateAPIRequest {
	return &AlibabaEinvoiceItemUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceItemUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.item.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceItemUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is InvoiceName Setter
// 商品的开票名称，对应发票的货物劳务名称，值DELETE时表示删除
func (r *AlibabaEinvoiceItemUpdateAPIRequest) SetInvoiceName(_invoiceName string) error {
	r._invoiceName = _invoiceName
	r.Set("invoice_name", _invoiceName)
	return nil
}

// Get InvoiceName Getter
func (r AlibabaEinvoiceItemUpdateAPIRequest) GetInvoiceName() string {
	return r._invoiceName
}

// Set is ItemId Setter
// 商品id，优先级高于outerId，商品必须归属于店铺，itemId和outerId不能同时为空
func (r *AlibabaEinvoiceItemUpdateAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// Get ItemId Getter
func (r AlibabaEinvoiceItemUpdateAPIRequest) GetItemId() int64 {
	return r._itemId
}

// Set is ItemNo Setter
// 税收分类编码，需要精确到叶子节点，必须和taxRate同时修改或删除，值DELETE时表示删除
func (r *AlibabaEinvoiceItemUpdateAPIRequest) SetItemNo(_itemNo string) error {
	r._itemNo = _itemNo
	r.Set("item_no", _itemNo)
	return nil
}

// Get ItemNo Getter
func (r AlibabaEinvoiceItemUpdateAPIRequest) GetItemNo() string {
	return r._itemNo
}

// Set is SkuId Setter
// skuId，必须是itemId下的sku，填写skuId后，修改和删除sku的开票信息
func (r *AlibabaEinvoiceItemUpdateAPIRequest) SetSkuId(_skuId int64) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// Get SkuId Getter
func (r AlibabaEinvoiceItemUpdateAPIRequest) GetSkuId() int64 {
	return r._skuId
}

// Set is Specification Setter
// 规格型号，值DELETE时表示删除
func (r *AlibabaEinvoiceItemUpdateAPIRequest) SetSpecification(_specification string) error {
	r._specification = _specification
	r.Set("specification", _specification)
	return nil
}

// Get Specification Getter
func (r AlibabaEinvoiceItemUpdateAPIRequest) GetSpecification() string {
	return r._specification
}

// Set is TaxRate Setter
// 税率，可选值0，3，4，5，6，10，11，13， 16，17，必须和itemNo同时修改或删除,值为DELETE时表示删除
func (r *AlibabaEinvoiceItemUpdateAPIRequest) SetTaxRate(_taxRate string) error {
	r._taxRate = _taxRate
	r.Set("tax_rate", _taxRate)
	return nil
}

// Get TaxRate Getter
func (r AlibabaEinvoiceItemUpdateAPIRequest) GetTaxRate() string {
	return r._taxRate
}

// Set is ZeroRateFlag Setter
// 0税率标识，只有税率为0的情况才有值，0=出口零税率，1=免税，2=不征收，3=普通零税率，值为DELETE时表示删除
func (r *AlibabaEinvoiceItemUpdateAPIRequest) SetZeroRateFlag(_zeroRateFlag string) error {
	r._zeroRateFlag = _zeroRateFlag
	r.Set("zero_rate_flag", _zeroRateFlag)
	return nil
}

// Get ZeroRateFlag Getter
func (r AlibabaEinvoiceItemUpdateAPIRequest) GetZeroRateFlag() string {
	return r._zeroRateFlag
}

// Set is Unit Setter
// 单位，值DELETE时表示删除
func (r *AlibabaEinvoiceItemUpdateAPIRequest) SetUnit(_unit string) error {
	r._unit = _unit
	r.Set("unit", _unit)
	return nil
}

// Get Unit Getter
func (r AlibabaEinvoiceItemUpdateAPIRequest) GetUnit() string {
	return r._unit
}

// Set is OuterId Setter
// 商家外部商品id，如果outerId对应了多个天猫sku，则会更新所有的sku开票信息。itemId和outerId不能同时为空
func (r *AlibabaEinvoiceItemUpdateAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// Get OuterId Getter
func (r AlibabaEinvoiceItemUpdateAPIRequest) GetOuterId() string {
	return r._outerId
}

// Set is UpdateSku Setter
// 是否根据outerId更新所有对应sku的开票信息，true=更新，false=开票信息维护在发票平台；自动开票时，根据skuId获取outerId，再根据outerId查询开票信息。outerId不为空时必填
func (r *AlibabaEinvoiceItemUpdateAPIRequest) SetUpdateSku(_updateSku bool) error {
	r._updateSku = _updateSku
	r.Set("update_sku", _updateSku)
	return nil
}

// Get UpdateSku Getter
func (r AlibabaEinvoiceItemUpdateAPIRequest) GetUpdateSku() bool {
	return r._updateSku
}
