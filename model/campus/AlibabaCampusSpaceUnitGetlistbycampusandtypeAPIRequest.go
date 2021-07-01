package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusSpaceUnitGetlistbycampusandtypeAPIRequest
根据园区id及TypeId获取空间单元 API请求
alibaba.campus.space.unit.getlistbycampusandtype

根据园区id及TypeId获取空间单元
HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceUnitApiTopService
HSF方法名称：getListByCampusAndType */
type AlibabaCampusSpaceUnitGetlistbycampusandtypeAPIRequest struct {
	model.Params
	// 系统自动生成
	_param0 *WorkBenchContext
	// 查询参数封装
	_param1 *SpaceUnitQuery
}

// New
