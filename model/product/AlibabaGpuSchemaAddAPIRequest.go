package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaGpuSchemaAddAPIRequest
使用schema文件发布产品 API请求
alibaba.gpu.schema.add

使用Schema文件发布一个产品 */
type AlibabaGpuSchemaAddAPIRequest struct {
	model.Params
	// 叶子类目ID
	_leafCatId int64
	// 品牌ID
	_brandId int64
	// 根据alibaba.gpu.add.schema.get获取的规则提交上来的schema
	_schemaXmlFields string
	// 当前用户所在渠道如0代表天猫，8代表淘宝
	_providerId int64
}

// New
