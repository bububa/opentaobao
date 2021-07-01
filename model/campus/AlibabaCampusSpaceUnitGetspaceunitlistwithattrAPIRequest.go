package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusSpaceUnitGetspaceunitlistwithattrAPIRequest
空间单元列表带业务属性实例 API请求
alibaba.campus.space.unit.getspaceunitlistwithattr

空间单元列表带业务属性实例 */
type AlibabaCampusSpaceUnitGetspaceunitlistwithattrAPIRequest struct {
	model.Params
	// 操作用户上下文
	_context *WorkBenchContext
	// 空间单元查询对象
	_unitQuery *SpaceUnitQuery
}

// New
