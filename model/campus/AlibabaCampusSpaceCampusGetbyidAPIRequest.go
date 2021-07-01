package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusSpaceCampusGetbyidAPIRequest
根据园区id获取园区信息 API请求
alibaba.campus.space.campus.getbyid

根据园区id获取园区信息
HSF接口名称：com.alibaba.campus.api.space.service.top.CampusApiTopService
HSF方法名称：getCampusById */
type AlibabaCampusSpaceCampusGetbyidAPIRequest struct {
	model.Params
	// 园区ID
	_param0 *WorkBenchContext
	// 园区ID
	_param1 int64
}

// New
