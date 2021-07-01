package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusSpaceUnitGetbyidAPIRequest
根据ID查询指定空间单元信息 API请求
alibaba.campus.space.unit.getbyid

根据ID查询指定空间单元信息
HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceUnitApiTopService
HSF方法名称：getById */
type AlibabaCampusSpaceUnitGetbyidAPIRequest struct {
	model.Params
	// 空间单元ID
	_param0 *WorkBenchContext
	// 空间单元ID
	_param1 int64
}

// New
