package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusSpaceTypeGetpageresultAPIRequest
分页查询空间类别接口 API请求
alibaba.campus.space.type.getpageresult

分页查询空间类别接口
HSF接口名称：com.alibaba.campus.space.api.top.SpaceTypeApiTopService
HSF方法名称：getPageResult */
type AlibabaCampusSpaceTypeGetpageresultAPIRequest struct {
	model.Params
	// 环境参数
	_param0 *WorkBenchContext
	// 查询参数
	_param1 *SpaceTypeQuery
}

// New
