package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterReservecondCreateAPIRequest
创建主动预约开通条件 API请求
tmall.servicecenter.reservecond.create

1、设置主动预约开通条件 */
type TmallServicecenterReservecondCreateAPIRequest struct {
	model.Params
	// 主动预约开通条件
	_rocList []ReserveOpenConditionDto
}

// NewTmallServicecenterReservecondCreateRequest 初始化TmallServicecenterReservecondCreateAPIRequest对象
func NewTmallServicecenterReservecondCreateRequest() *TmallServicecenterReservecondCreateAPIRequest {
	return &TmallServicecenterReservecondCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterReservecondCreateAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.reservecond.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterReservecondCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RocList Setter
// 主动预约开通条件
func (r *TmallServicecenterReservecondCreateAPIRequest) SetRocList(_rocList []ReserveOpenConditionDto) error {
	r._rocList = _rocList
	r.Set("roc_list", _rocList)
	return nil
}

// Get RocList Getter
func (r TmallServicecenterReservecondCreateAPIRequest) GetRocList() []ReserveOpenConditionDto {
	return r._rocList
}
