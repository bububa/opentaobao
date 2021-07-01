package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusSpaceGroupGetspacegroupwithattrAPIRequest
空间分组id查业务属性实例 API请求
alibaba.campus.space.group.getspacegroupwithattr

空间分组id查业务属性实例 */
type AlibabaCampusSpaceGroupGetspacegroupwithattrAPIRequest struct {
	model.Params
	// 操作用户上下文
	_context *WorkBenchContext
	// 空间单元id
	_groupId int64
}

// New
