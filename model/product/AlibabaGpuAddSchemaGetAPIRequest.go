package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaGpuAddSchemaGetAPIRequest
获取产品发布规则接口 API请求
alibaba.gpu.add.schema.get

获取产品发布规则接口 */
type AlibabaGpuAddSchemaGetAPIRequest struct {
	model.Params
	// 叶子类目ID
	_leafCatId int64
	// 品牌ID
	_brandId int64
	// 当前用户所在渠道如0代表天猫，8代表淘宝
	_providerId int64
}

// New
