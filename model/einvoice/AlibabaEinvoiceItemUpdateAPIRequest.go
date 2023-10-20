package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoiceitemupdateAPIRequest 修改商品开票信息 API请求
// alibaba.einvoice.item.update
//
// ERP通过接口将商品的开票信息同步给阿里发票平台，自动开票时将读取这些开票信息，需要联系阿里小二开通对应的权限
type AlibabaeinvoiceitemupdateAPIRequest struct {
	model.Params
	// 商品的开票名称，对应发票的货物劳务名称，值DELETE时表示删除
	_invoiceName string
	// 税收分类编码，需要精确到叶子节点，必须和taxRate同时修改或删除，值DELETE时表示删除
	_itemNo string
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
	// 商品id，优先级高于outerId，商品必须归属于店铺，itemId和outerId不能同时为空
	_itemId int64
	// skuId，必须是itemId下的sku，填写skuId后，修改和删除sku的开票信息
	_skuId int64
	// 是否根据outerId更新所有对应sku的开票信息，true=更新，false=开票信息维护在发票平台；自动开票时，根据skuId获取outerId，再根据outerId查询开票信息。outerId不为空时必填
	_updateSku bool
}

// NewAlibabaeinvoiceitemupdateRequest 初始化AlibabaeinvoiceitemupdateAPIRequest对象
func NewAlibabaeinvoiceitemupdateRequest() *AlibabaeinvoiceitemupdateAPIRequest {
	return &AlibabaeinvoiceitemupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeinvoiceitemupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.item.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeinvoiceitemupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeinvoiceitemupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInvoiceName is InvoiceName Setter
// 商品的开票名称，对应发票的货物劳务名称，值DELETE时表示删除
func (r *AlibabaeinvoiceitemupdateAPIRequest) SetInvoiceName(_invoiceName string) error {
	r._invoiceName = _invoiceName
	r.Set("invoice_name", _invoiceName)
	return nil
}

// GetInvoiceName InvoiceName Getter
func (r AlibabaeinvoiceitemupdateAPIRequest) GetInvoiceName() string {
	return r._invoiceName
}

// SetItemNo is ItemNo Setter
// 税收分类编码，需要精确到叶子节点，必须和taxRate同时修改或删除，值DELETE时表示删除
func (r *AlibabaeinvoiceitemupdateAPIRequest) SetItemNo(_itemNo string) error {
	r._itemNo = _itemNo
	r.Set("item_no", _itemNo)
	return nil
}

// GetItemNo ItemNo Getter
func (r AlibabaeinvoiceitemupdateAPIRequest) GetItemNo() string {
	return r._itemNo
}

// SetSpecification is Specification Setter
// 规格型号，值DELETE时表示删除
func (r *AlibabaeinvoiceitemupdateAPIRequest) SetSpecification(_specification string) error {
	r._specification = _specification
	r.Set("specification", _specification)
	return nil
}

// GetSpecification Specification Getter
func (r AlibabaeinvoiceitemupdateAPIRequest) GetSpecification() string {
	return r._specification
}

// SetTaxRate is TaxRate Setter
// 税率，可选值0，3，4，5，6，10，11，13， 16，17，必须和itemNo同时修改或删除,值为DELETE时表示删除
func (r *AlibabaeinvoiceitemupdateAPIRequest) SetTaxRate(_taxRate string) error {
	r._taxRate = _taxRate
	r.Set("tax_rate", _taxRate)
	return nil
}

// GetTaxRate TaxRate Getter
func (r AlibabaeinvoiceitemupdateAPIRequest) GetTaxRate() string {
	return r._taxRate
}

// SetZeroRateFlag is ZeroRateFlag Setter
// 0税率标识，只有税率为0的情况才有值，0=出口零税率，1=免税，2=不征收，3=普通零税率，值为DELETE时表示删除
func (r *AlibabaeinvoiceitemupdateAPIRequest) SetZeroRateFlag(_zeroRateFlag string) error {
	r._zeroRateFlag = _zeroRateFlag
	r.Set("zero_rate_flag", _zeroRateFlag)
	return nil
}

// GetZeroRateFlag ZeroRateFlag Getter
func (r AlibabaeinvoiceitemupdateAPIRequest) GetZeroRateFlag() string {
	return r._zeroRateFlag
}

// SetUnit is Unit Setter
// 单位，值DELETE时表示删除
func (r *AlibabaeinvoiceitemupdateAPIRequest) SetUnit(_unit string) error {
	r._unit = _unit
	r.Set("unit", _unit)
	return nil
}

// GetUnit Unit Getter
func (r AlibabaeinvoiceitemupdateAPIRequest) GetUnit() string {
	return r._unit
}

// SetOuterId is OuterId Setter
// 商家外部商品id，如果outerId对应了多个天猫sku，则会更新所有的sku开票信息。itemId和outerId不能同时为空
func (r *AlibabaeinvoiceitemupdateAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r AlibabaeinvoiceitemupdateAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetItemId is ItemId Setter
// 商品id，优先级高于outerId，商品必须归属于店铺，itemId和outerId不能同时为空
func (r *AlibabaeinvoiceitemupdateAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlibabaeinvoiceitemupdateAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetSkuId is SkuId Setter
// skuId，必须是itemId下的sku，填写skuId后，修改和删除sku的开票信息
func (r *AlibabaeinvoiceitemupdateAPIRequest) SetSkuId(_skuId int64) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// GetSkuId SkuId Getter
func (r AlibabaeinvoiceitemupdateAPIRequest) GetSkuId() int64 {
	return r._skuId
}

// SetUpdateSku is UpdateSku Setter
// 是否根据outerId更新所有对应sku的开票信息，true=更新，false=开票信息维护在发票平台；自动开票时，根据skuId获取outerId，再根据outerId查询开票信息。outerId不为空时必填
func (r *AlibabaeinvoiceitemupdateAPIRequest) SetUpdateSku(_updateSku bool) error {
	r._updateSku = _updateSku
	r.Set("update_sku", _updateSku)
	return nil
}

// GetUpdateSku UpdateSku Getter
func (r AlibabaeinvoiceitemupdateAPIRequest) GetUpdateSku() bool {
	return r._updateSku
}
