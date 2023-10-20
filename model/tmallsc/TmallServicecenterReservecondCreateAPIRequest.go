package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterreservecondcreateAPIRequest 创建主动预约开通条件 API请求
// tmall.servicecenter.reservecond.create
//
// 1、设置主动预约开通条件
type TmallservicecenterreservecondcreateAPIRequest struct {
	model.Params
	// 主动预约开通条件
	_rocList []ReserveOpenConditionDto
}

// NewTmallservicecenterreservecondcreateRequest 初始化TmallservicecenterreservecondcreateAPIRequest对象
func NewTmallservicecenterreservecondcreateRequest() *TmallservicecenterreservecondcreateAPIRequest {
	return &TmallservicecenterreservecondcreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenterreservecondcreateAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.reservecond.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenterreservecondcreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenterreservecondcreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRocList is RocList Setter
// 主动预约开通条件
func (r *TmallservicecenterreservecondcreateAPIRequest) SetRocList(_rocList []ReserveOpenConditionDto) error {
	r._rocList = _rocList
	r.Set("roc_list", _rocList)
	return nil
}

// GetRocList RocList Getter
func (r TmallservicecenterreservecondcreateAPIRequest) GetRocList() []ReserveOpenConditionDto {
	return r._rocList
}
