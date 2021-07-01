package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusSpaceGroupGetlistAPIRequest
多条件查询空间分组信息 API请求
alibaba.campus.space.group.getlist

多条件查询空间分组信息
HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceGroupApiTopService
HSF方法名称：getList */
type AlibabaCampusSpaceGroupGetlistAPIRequest struct {
	model.Params
	// 查询条件封装
	_param0 *WorkBenchContext
	// 查询参数封装
	_param1 *SpaceGroupQuery
}

// New
