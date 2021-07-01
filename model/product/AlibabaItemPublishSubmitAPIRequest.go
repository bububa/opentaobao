package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaItemPublishSubmitAPIRequest
商品发布 API请求
alibaba.item.publish.submit

新商品发布，提交商品发布信息 */
type AlibabaItemPublishSubmitAPIRequest struct {
	model.Params
	// 业务扩展参数，需与平台约定好
	_bizType string
	// 商品发布的市场。taobao:淘宝,tmall:天猫,litetao:淘宝特价版
	_market string
	// 商品类目ID
	_catId int64
	// 产品ID，天猫市场(market=tmall)时必填
	_spuId int64
	// 商品条码
	_barcode string
	// 商品schema信息，通过alibaba.item.publish.props.get获取并补全后提交
	_schema string
}

// New
