package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterreservecondupdateAPIRequest 主动预约条件更新 API请求
// tmall.servicecenter.reservecond.update
//
// 1、设置主动预约开通条件
type TmallservicecenterreservecondupdateAPIRequest struct {
	model.Params
	// 主动预约开通条件
	_rocList []ReserveOpenConditionDto
}

// NewTmallservicecenterreservecondupdateRequest 初始化TmallservicecenterreservecondupdateAPIRequest对象
func NewTmallservicecenterreservecondupdateRequest() *TmallservicecenterreservecondupdateAPIRequest {
	return &TmallservicecenterreservecondupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenterreservecondupdateAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.reservecond.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenterreservecondupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenterreservecondupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRocList is RocList Setter
// 主动预约开通条件
func (r *TmallservicecenterreservecondupdateAPIRequest) SetRocList(_rocList []ReserveOpenConditionDto) error {
	r._rocList = _rocList
	r.Set("roc_list", _rocList)
	return nil
}

// GetRocList RocList Getter
func (r TmallservicecenterreservecondupdateAPIRequest) GetRocList() []ReserveOpenConditionDto {
	return r._rocList
}
