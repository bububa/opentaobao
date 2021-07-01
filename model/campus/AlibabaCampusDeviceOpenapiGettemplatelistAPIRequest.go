package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusDeviceOpenapiGettemplatelistAPIRequest
查询设备模板 API请求
alibaba.campus.device.openapi.gettemplatelist

查询设备模板信息 */
type AlibabaCampusDeviceOpenapiGettemplatelistAPIRequest struct {
	model.Params
	// 设备模板查询对象
	_query *TemplateApiQuery
}

// New
