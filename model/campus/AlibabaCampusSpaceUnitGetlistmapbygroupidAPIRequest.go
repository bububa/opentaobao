package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusSpaceUnitGetlistmapbygroupidAPIRequest
新增查询多个分组ID各自相关的空间单元信息 API请求
alibaba.campus.space.unit.getlistmapbygroupid

新增查询多个分组ID各自相关的空间单元信息
HSF接口名称：	com.alibaba.campus.api.space.service.top.SpaceUnitApiTopService
HSF方法名称：	getListMapByGroupIds */
type AlibabaCampusSpaceUnitGetlistmapbygroupidAPIRequest struct {
	model.Params
	// 用户环境
	_param0 *WorkBenchContext
	// 查询封装
	_param1 *SpaceUnitQuery
}

// New
