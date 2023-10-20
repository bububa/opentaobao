package tuanhotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitriptuanhoteladaptstoregetAPIRequest 酒店团购套餐关联适用门店 API请求
// alitrip.tuan.hotel.adapt.store.get
//
// 输入shid，返回关联门店详情信息
type AlitriptuanhoteladaptstoregetAPIRequest struct {
	model.Params
	// 标准酒店ID列表,逗号分割。与hid_list二者只能选一
	_shidList []int64
	// 物理酒店ID列表，逗号分割。与shid_list二者只能选一
	_hidList []int64
}

// NewAlitriptuanhoteladaptstoregetRequest 初始化AlitriptuanhoteladaptstoregetAPIRequest对象
func NewAlitriptuanhoteladaptstoregetRequest() *AlitriptuanhoteladaptstoregetAPIRequest {
	return &AlitriptuanhoteladaptstoregetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitriptuanhoteladaptstoregetAPIRequest) GetApiMethodName() string {
	return "alitrip.tuan.hotel.adapt.store.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitriptuanhoteladaptstoregetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitriptuanhoteladaptstoregetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetShidList is ShidList Setter
// 标准酒店ID列表,逗号分割。与hid_list二者只能选一
func (r *AlitriptuanhoteladaptstoregetAPIRequest) SetShidList(_shidList []int64) error {
	r._shidList = _shidList
	r.Set("shid_list", _shidList)
	return nil
}

// GetShidList ShidList Getter
func (r AlitriptuanhoteladaptstoregetAPIRequest) GetShidList() []int64 {
	return r._shidList
}

// SetHidList is HidList Setter
// 物理酒店ID列表，逗号分割。与shid_list二者只能选一
func (r *AlitriptuanhoteladaptstoregetAPIRequest) SetHidList(_hidList []int64) error {
	r._hidList = _hidList
	r.Set("hid_list", _hidList)
	return nil
}

// GetHidList HidList Getter
func (r AlitriptuanhoteladaptstoregetAPIRequest) GetHidList() []int64 {
	return r._hidList
}
