package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaItemCategoryPredictAPIRequest
商品发布类目预测 API请求
alibaba.item.category.predict

<font color='red'>商品发布类目预测接口，预测匹配的结果存在一定误差，需要商家二次确认，避免类目配置错误产生其他影响。</font> */
type AlibabaItemCategoryPredictAPIRequest struct {
	model.Params
	// 商品主图链接，最多5张，传入完整URL
	_images []string
	// 商品发布的市场。taobao:淘宝,tmall:天猫,litetao:淘宝特价版
	_market string
	// 商品条码
	_barcode string
	// 商品条码图片
	_barcodeImage string
	// 商品介绍
	_itemDesc string
}

// New
