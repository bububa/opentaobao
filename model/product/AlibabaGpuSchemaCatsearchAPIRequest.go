package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaGpuSchemaCatsearchAPIRequest
按类目查询spu接口 API请求
alibaba.gpu.schema.catsearch

按类目查询spu的schema接口 */
type AlibabaGpuSchemaCatsearchAPIRequest struct {
	model.Params
	// 叶子类目ID
	_leafCatId int64
	// 当前页
	_currentPage int64
	// 每页大小
	_pageSize int64
	// 渠道Id，如0代表天猫，8代表淘宝
	_providerId int64
}

// New
