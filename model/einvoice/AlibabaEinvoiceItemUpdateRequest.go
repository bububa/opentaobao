package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改商品开票信息 API请求
alibaba.einvoice.item.update

ERP通过接口将商品的开票信息同步给阿里发票平台，自动开票时将读取这些开票信息，需要联系阿里小二开通对应的权限
*/
type AlibabaEinvoiceItemUpdateRequest struct {
    model.Params
    // 商品的开票名称，对应发票的货物劳务名称，值DELETE时表示删除
    invoiceName   string
    // 商品id，优先级高于outerId，商品必须归属于店铺，itemId和outerId不能同时为空
    itemId   int64
    // 税收分类编码，需要精确到叶子节点，必须和taxRate同时修改或删除，值DELETE时表示删除
    itemNo   string
    // skuId，必须是itemId下的sku，填写skuId后，修改和删除sku的开票信息
    skuId   int64
    // 规格型号，值DELETE时表示删除
    specification   string
    // 税率，可选值0，3，4，5，6，10，11，13， 16，17，必须和itemNo同时修改或删除,值为DELETE时表示删除
    taxRate   string
    // 0税率标识，只有税率为0的情况才有值，0=出口零税率，1=免税，2=不征收，3=普通零税率，值为DELETE时表示删除
    zeroRateFlag   string
    // 单位，值DELETE时表示删除
    unit   string
    // 商家外部商品id，如果outerId对应了多个天猫sku，则会更新所有的sku开票信息。itemId和outerId不能同时为空
    outerId   string
    // 是否根据outerId更新所有对应sku的开票信息，true=更新，false=开票信息维护在发票平台；自动开票时，根据skuId获取outerId，再根据outerId查询开票信息。outerId不为空时必填
    updateSku   bool
}

// 初始化AlibabaEinvoiceItemUpdateRequest对象
func NewAlibabaEinvoiceItemUpdateRequest() *AlibabaEinvoiceItemUpdateRequest{
    return &AlibabaEinvoiceItemUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceItemUpdateRequest) GetApiMethodName() string {
    return "alibaba.einvoice.item.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceItemUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// InvoiceName Setter
// 商品的开票名称，对应发票的货物劳务名称，值DELETE时表示删除
func (r *AlibabaEinvoiceItemUpdateRequest) SetInvoiceName(invoiceName string) error {
    r.invoiceName = invoiceName
    r.Set("invoice_name", invoiceName)
    return nil
}

// InvoiceName Getter
func (r AlibabaEinvoiceItemUpdateRequest) GetInvoiceName() string {
    return r.invoiceName
}
// ItemId Setter
// 商品id，优先级高于outerId，商品必须归属于店铺，itemId和outerId不能同时为空
func (r *AlibabaEinvoiceItemUpdateRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r AlibabaEinvoiceItemUpdateRequest) GetItemId() int64 {
    return r.itemId
}
// ItemNo Setter
// 税收分类编码，需要精确到叶子节点，必须和taxRate同时修改或删除，值DELETE时表示删除
func (r *AlibabaEinvoiceItemUpdateRequest) SetItemNo(itemNo string) error {
    r.itemNo = itemNo
    r.Set("item_no", itemNo)
    return nil
}

// ItemNo Getter
func (r AlibabaEinvoiceItemUpdateRequest) GetItemNo() string {
    return r.itemNo
}
// SkuId Setter
// skuId，必须是itemId下的sku，填写skuId后，修改和删除sku的开票信息
func (r *AlibabaEinvoiceItemUpdateRequest) SetSkuId(skuId int64) error {
    r.skuId = skuId
    r.Set("sku_id", skuId)
    return nil
}

// SkuId Getter
func (r AlibabaEinvoiceItemUpdateRequest) GetSkuId() int64 {
    return r.skuId
}
// Specification Setter
// 规格型号，值DELETE时表示删除
func (r *AlibabaEinvoiceItemUpdateRequest) SetSpecification(specification string) error {
    r.specification = specification
    r.Set("specification", specification)
    return nil
}

// Specification Getter
func (r AlibabaEinvoiceItemUpdateRequest) GetSpecification() string {
    return r.specification
}
// TaxRate Setter
// 税率，可选值0，3，4，5，6，10，11，13， 16，17，必须和itemNo同时修改或删除,值为DELETE时表示删除
func (r *AlibabaEinvoiceItemUpdateRequest) SetTaxRate(taxRate string) error {
    r.taxRate = taxRate
    r.Set("tax_rate", taxRate)
    return nil
}

// TaxRate Getter
func (r AlibabaEinvoiceItemUpdateRequest) GetTaxRate() string {
    return r.taxRate
}
// ZeroRateFlag Setter
// 0税率标识，只有税率为0的情况才有值，0=出口零税率，1=免税，2=不征收，3=普通零税率，值为DELETE时表示删除
func (r *AlibabaEinvoiceItemUpdateRequest) SetZeroRateFlag(zeroRateFlag string) error {
    r.zeroRateFlag = zeroRateFlag
    r.Set("zero_rate_flag", zeroRateFlag)
    return nil
}

// ZeroRateFlag Getter
func (r AlibabaEinvoiceItemUpdateRequest) GetZeroRateFlag() string {
    return r.zeroRateFlag
}
// Unit Setter
// 单位，值DELETE时表示删除
func (r *AlibabaEinvoiceItemUpdateRequest) SetUnit(unit string) error {
    r.unit = unit
    r.Set("unit", unit)
    return nil
}

// Unit Getter
func (r AlibabaEinvoiceItemUpdateRequest) GetUnit() string {
    return r.unit
}
// OuterId Setter
// 商家外部商品id，如果outerId对应了多个天猫sku，则会更新所有的sku开票信息。itemId和outerId不能同时为空
func (r *AlibabaEinvoiceItemUpdateRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

// OuterId Getter
func (r AlibabaEinvoiceItemUpdateRequest) GetOuterId() string {
    return r.outerId
}
// UpdateSku Setter
// 是否根据outerId更新所有对应sku的开票信息，true=更新，false=开票信息维护在发票平台；自动开票时，根据skuId获取outerId，再根据outerId查询开票信息。outerId不为空时必填
func (r *AlibabaEinvoiceItemUpdateRequest) SetUpdateSku(updateSku bool) error {
    r.updateSku = updateSku
    r.Set("update_sku", updateSku)
    return nil
}

// UpdateSku Getter
func (r AlibabaEinvoiceItemUpdateRequest) GetUpdateSku() bool {
    return r.updateSku
}
