package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusSpaceGroupGetbyidAPIRequest
根据分组ID查询相关的空间分组信息 API请求
alibaba.campus.space.group.getbyid

根据分组ID查询相关的空间分组信息
HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceGroupApiTopService
HSF方法名称：getById */
type AlibabaCampusSpaceGroupGetbyidAPIRequest struct {
	model.Params
	// 用户环境
	_param0 *WorkBenchContext
	// 分组ID
	_param1 int64
}

// New
