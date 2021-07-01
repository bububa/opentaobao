package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusSpaceGetbyidsAPIRequest
根据ids和类型查询空间列表 API请求
alibaba.campus.space.getbyids

根据ids和类型查询空间列表 */
type AlibabaCampusSpaceGetbyidsAPIRequest struct {
	model.Params
	// 上下文
	_context *WorkBenchContext
	// 查询条件
	_query *SpaceIdsQuery
}

// New
