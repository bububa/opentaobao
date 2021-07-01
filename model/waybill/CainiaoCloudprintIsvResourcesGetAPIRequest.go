package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoCloudprintIsvResourcesGetAPIRequest
isv资源查询 API请求
cainiao.cloudprint.isv.resources.get

isv资源查询，包括isv模板、打印项、预设的自定义区等 */
type CainiaoCloudprintIsvResourcesGetAPIRequest struct {
	model.Params
	// isv资源类型，分为：TEMPLATE（表示模板），PRINT_ITEM（打印项），CUSTOM_AREA（预设自定义区）
	_isvResourceType string
}

// New
