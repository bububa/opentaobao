package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusSpaceUnitGetlistbygroupidAPIRequest
根据分组ID查询相应的空间单元 API请求
alibaba.campus.space.unit.getlistbygroupid

根据分组ID查询相应的空间单元
HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceUnitApiTopService
HSF方法名称：getListByGroupId */
type AlibabaCampusSpaceUnitGetlistbygroupidAPIRequest struct {
	model.Params
	// 分组ID
	_param0 *WorkBenchContext
	// 分组ID
	_param1 int64
}

// New
