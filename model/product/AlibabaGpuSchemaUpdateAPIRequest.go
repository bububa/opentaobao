package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaGpuSchemaUpdateAPIRequest
产品更新接口 API请求
alibaba.gpu.schema.update

产品更新接口 */
type AlibabaGpuSchemaUpdateAPIRequest struct {
	model.Params
	// 产品ID
	_productId int64
	// 更新产品提交的schema数据
	_schemaXmlFields string
	// 当前用户所在渠道如0代表天猫，8代表淘宝
	_providerId int64
}

// New
