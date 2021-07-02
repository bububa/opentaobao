package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenDeletestandAPIRequest 大麦换验平台-第三方对外开放-看台接口deleteStand API请求
// alibaba.damai.mev.open.deletestand
//
// deleteStand
type AlibabaDamaiMevOpenDeletestandAPIRequest struct {
	model.Params
	// 入参deleteStandParam
	_deleteStandParam *StandIdOpenParam
}

// NewAlibabaDamaiMevOpenDeletestandRequest 初始化AlibabaDamaiMevOpenDeletestandAPIRequest对象
func NewAlibabaDamaiMevOpenDeletestandRequest() *AlibabaDamaiMevOpenDeletestandAPIRequest {
	return &AlibabaDamaiMevOpenDeletestandAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenDeletestandAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.deletestand"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenDeletestandAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetDeleteStandParam is DeleteStandParam Setter
// 入参deleteStandParam
func (r *AlibabaDamaiMevOpenDeletestandAPIRequest) SetDeleteStandParam(_deleteStandParam *StandIdOpenParam) error {
	r._deleteStandParam = _deleteStandParam
	r.Set("delete_stand_param", _deleteStandParam)
	return nil
}

// GetDeleteStandParam DeleteStandParam Getter
func (r AlibabaDamaiMevOpenDeletestandAPIRequest) GetDeleteStandParam() *StandIdOpenParam {
	return r._deleteStandParam
}
