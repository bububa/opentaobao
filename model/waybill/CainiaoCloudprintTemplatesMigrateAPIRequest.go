package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoCloudprintTemplatesMigrateAPIRequest
云打印模板迁移接口 API请求
cainiao.cloudprint.templates.migrate

云打印模板迁移接口 */
type CainiaoCloudprintTemplatesMigrateAPIRequest struct {
	model.Params
	// 标准电子面单模板的id
	_tempalteId int64
	// 自定义区名称
	_customAreaName string
	// 自定义区内容
	_customAreaContent string
}

// New
