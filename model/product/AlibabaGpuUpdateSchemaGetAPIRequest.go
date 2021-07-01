package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaGpuUpdateSchemaGetAPIRequest
获取产品编辑schema规则的接口 API请求
alibaba.gpu.update.schema.get

获取产品编辑schema规则的接口 */
type AlibabaGpuUpdateSchemaGetAPIRequest struct {
	model.Params
	// 产品ID
	_productId int64
	// 当前用户所在渠道如0代表天猫，8代表淘宝
	_providerId int64
}

// New
