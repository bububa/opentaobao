package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIRequest
根据园区id及TypeId获取空间分组 API请求
alibaba.campus.space.group.getlistbycampusandtype

根据园区id及TypeId获取空间分组
HSF接口名称：com.alibaba.campus.api.space.service.top.SpaceGroupApiTopService
HSF方法名称：getListByCampusAndType */
type AlibabaCampusSpaceGroupGetlistbycampusandtypeAPIRequest struct {
	model.Params
	// 系统自动生成
	_param0 *WorkBenchContext
	// 查询参数封装
	_param1 *SpaceGroupQuery
}

// New
