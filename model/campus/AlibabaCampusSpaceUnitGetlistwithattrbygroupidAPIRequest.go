package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusSpaceUnitGetlistwithattrbygroupidAPIRequest
根据空间分组id、ids查空间单元信息【带空间单元业务属性信息】 API请求
alibaba.campus.space.unit.getlistwithattrbygroupid

根据空间分组id、ids查空间单元信息【带空间单元业务属性信息】 */
type AlibabaCampusSpaceUnitGetlistwithattrbygroupidAPIRequest struct {
	model.Params
	// 操作用户上下文
	_context *WorkBenchContext
	// 分组id
	_groupId int64
}

// New
