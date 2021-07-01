package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusSpaceUnitGetlistAPIRequest
多条件查询空间单元信息 API请求
alibaba.campus.space.unit.getlist

多条件查询空间单元信息
HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceUnitApiTopService
HSF方法名称：getList */
type AlibabaCampusSpaceUnitGetlistAPIRequest struct {
	model.Params
	// 查询条件封装
	_param0 *WorkBenchContext
	// 查询参数封装
	_param1 *SpaceUnitQuery
}

// New
