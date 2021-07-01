package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusSpaceUnitGetspaceunitwithattrAPIRequest
空间单元id查业务属性实例 API请求
alibaba.campus.space.unit.getspaceunitwithattr

空间单元id查业务属性实例 */
type AlibabaCampusSpaceUnitGetspaceunitwithattrAPIRequest struct {
	model.Params
	// 操作用户上下文
	_context *WorkBenchContext
	// 空间单元id
	_spaceUnitId int64
}

// New
