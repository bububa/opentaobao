package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusSpaceTypeGetbycodeAPIRequest 根据类别编码查询类别 API请求
// alibaba.campus.space.type.getbycode
//
// 根据类别编码查询类别
// HSF接口名称：com.alibaba.campus.space.api.top.SpaceTypeApiTopService
// HSF方法名称：getByCode
type AlibabaCampusSpaceTypeGetbycodeAPIRequest struct {
	model.Params
	// 查询条件封装
	_param0 *WorkBenchContext
	// 空间类别编码
	_typeCode string
}

// NewAlibabaCampusSpaceTypeGetbycodeRequest 初始化AlibabaCampusSpaceTypeGetbycodeAPIRequest对象
func NewAlibabaCampusSpaceTypeGetbycodeRequest() *AlibabaCampusSpaceTypeGetbycodeAPIRequest {
	return &AlibabaCampusSpaceTypeGetbycodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusSpaceTypeGetbycodeAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.space.type.getbycode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusSpaceTypeGetbycodeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param0 Setter
// 查询条件封装
func (r *AlibabaCampusSpaceTypeGetbycodeAPIRequest) SetParam0(_param0 *WorkBenchContext) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// Get Param0 Getter
func (r AlibabaCampusSpaceTypeGetbycodeAPIRequest) GetParam0() *WorkBenchContext {
	return r._param0
}

// Set is TypeCode Setter
// 空间类别编码
func (r *AlibabaCampusSpaceTypeGetbycodeAPIRequest) SetTypeCode(_typeCode string) error {
	r._typeCode = _typeCode
	r.Set("type_code", _typeCode)
	return nil
}

// Get TypeCode Getter
func (r AlibabaCampusSpaceTypeGetbycodeAPIRequest) GetTypeCode() string {
	return r._typeCode
}
