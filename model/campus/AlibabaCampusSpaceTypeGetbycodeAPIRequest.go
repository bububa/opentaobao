package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusSpaceTypeGetbycodeAPIRequest
根据类别编码查询类别 API请求
alibaba.campus.space.type.getbycode

根据类别编码查询类别
HSF接口名称：com.alibaba.campus.space.api.top.SpaceTypeApiTopService
HSF方法名称：getByCode */
type AlibabaCampusSpaceTypeGetbycodeAPIRequest struct {
	model.Params
	// 查询条件封装
	_param0 *WorkBenchContext
	// 空间类别编码
	_typeCode string
}

// New
