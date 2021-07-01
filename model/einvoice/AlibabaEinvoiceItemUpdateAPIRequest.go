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

// New
