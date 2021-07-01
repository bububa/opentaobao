package alihealthcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthPregnancyProductGetAPIRequest
备孕首页获取达人配置商品 API请求
alibaba.alihealth.pregnancy.product.get

备孕首页获取达人配置商品 */
type AlibabaAlihealthPregnancyProductGetAPIRequest struct {
	model.Params
	// tab页对应id
	_sceneId int64
	// 起始页码，大于0
	_currentPage int64
	// 每页条数
	_pageSize int64
}

// New
