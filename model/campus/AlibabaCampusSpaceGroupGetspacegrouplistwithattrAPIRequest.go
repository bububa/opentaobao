package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIRequest
分页查询空间分组业务属性 API请求
alibaba.campus.space.group.getspacegrouplistwithattr

分页查询空间分组业务属性 */
type AlibabaCampusSpaceGroupGetspacegrouplistwithattrAPIRequest struct {
	model.Params
	// 操作用户上下文
	_context *WorkBenchContext
	// 查询对象
	_groupQuery *SpaceGroupQuery
}

// New
