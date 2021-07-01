package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaItemPublishSchemaGetAPIRequest
获取商品发布规则信息 API请求
alibaba.item.publish.schema.get

新商品发布，获取商品发布规则信息 */
type AlibabaItemPublishSchemaGetAPIRequest struct {
	model.Params
	// 业务扩展参数，需与平台约定好
	_bizType string
	// 商品主图链接，最多5张，传入完整URL
	_images []string
	// 商品类型。b:一口价  a:拍卖  默认值b一口价
	_itemType string
	// 商品发布的市场。taobao:淘宝,tmall:天猫,litetao:淘宝特价版
	_market string
	// 商品类目ID
	_catId int64
	// 产品ID，天猫市场(market=tmall)时必填
	_spuId int64
	// 商品条码
	_barcode string
}

// New
