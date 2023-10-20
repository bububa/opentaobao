package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimevopendeletestandAPIRequest 大麦换验平台-第三方对外开放-看台接口deleteStand API请求
// alibaba.damai.mev.open.deletestand
//
// deleteStand
type AlibabadamaimevopendeletestandAPIRequest struct {
	model.Params
	// 入参deleteStandParam
	_deleteStandParam *StandIdOpenParam
}

// NewAlibabadamaimevopendeletestandRequest 初始化AlibabadamaimevopendeletestandAPIRequest对象
func NewAlibabadamaimevopendeletestandRequest() *AlibabadamaimevopendeletestandAPIRequest {
	return &AlibabadamaimevopendeletestandAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadamaimevopendeletestandAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.deletestand"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadamaimevopendeletestandAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadamaimevopendeletestandAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeleteStandParam is DeleteStandParam Setter
// 入参deleteStandParam
func (r *AlibabadamaimevopendeletestandAPIRequest) SetDeleteStandParam(_deleteStandParam *StandIdOpenParam) error {
	r._deleteStandParam = _deleteStandParam
	r.Set("delete_stand_param", _deleteStandParam)
	return nil
}

// GetDeleteStandParam DeleteStandParam Getter
func (r AlibabadamaimevopendeletestandAPIRequest) GetDeleteStandParam() *StandIdOpenParam {
	return r._deleteStandParam
}
