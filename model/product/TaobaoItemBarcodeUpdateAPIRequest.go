package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoItemBarcodeUpdateAPIRequest
更新商品条形码信息 API请求
taobao.item.barcode.update

通过该接口，将商品以及SKU上得条形码信息补全 */
type TaobaoItemBarcodeUpdateAPIRequest struct {
	model.Params
	// 被更新商品的ID
	_itemId int64
	// 商品条形码，如果不用更新，可选择不填
	_itemBarcode string
	// 被更新SKU的ID列表，中间以英文逗号进行分隔。如果没有SKU或者不需要更新SKU的条形码，不需要设置
	_skuIds string
	// SKU维度的条形码，和sku_ids字段一一对应，中间以英文逗号分隔
	_skuBarcodes string
	// 是否强制保存商品条码。true：强制保存false ：需要执行条码库校验
	_isforce bool
	// 访问来源，这字段提供给千牛扫码枪用，其他调用方，不需要填写
	_src string
}

// New
