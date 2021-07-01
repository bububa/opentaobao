package alihealthlab

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthLabItemTbitemskuRelationSyncAPIRequest
阿里健康检验检测业务，检验检测项目淘宝商品SKU关系同步 API请求
alibaba.alihealth.lab.item.tbitemsku.relation.sync

阿里健康检验检测业务，检验检测项目淘宝商品SKU关系同步 */
type AlibabaAlihealthLabItemTbitemskuRelationSyncAPIRequest struct {
	model.Params
	// EFFECTIVE 有效，INVALID 无效
	_isvRelationStatus string
	// 关联的淘宝商品SKU id，在商品没有sku的情况下传0
	_tbSkuId int64
	// 关联的淘宝商品 id
	_tbItemId int64
	// 检验检测项目isv侧code
	_isvItemCode string
}

// New
