package tuanhotel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTuanHotelAdaptStoreGetAPIRequest 酒店团购套餐关联适用门店 API请求
// alitrip.tuan.hotel.adapt.store.get
//
// 输入shid，返回关联门店详情信息
type AlitripTuanHotelAdaptStoreGetAPIRequest struct {
	model.Params
	// 标准酒店ID列表,逗号分割。与hid_list二者只能选一
	_shidList []int64
	// 物理酒店ID列表，逗号分割。与shid_list二者只能选一
	_hidList []int64
}

// NewAlitripTuanHotelAdaptStoreGetRequest 初始化AlitripTuanHotelAdaptStoreGetAPIRequest对象
func NewAlitripTuanHotelAdaptStoreGetRequest() *AlitripTuanHotelAdaptStoreGetAPIRequest {
	return &AlitripTuanHotelAdaptStoreGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripTuanHotelAdaptStoreGetAPIRequest) Reset() {
	r._shidList = r._shidList[:0]
	r._hidList = r._hidList[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripTuanHotelAdaptStoreGetAPIRequest) GetApiMethodName() string {
	return "alitrip.tuan.hotel.adapt.store.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripTuanHotelAdaptStoreGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripTuanHotelAdaptStoreGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetShidList is ShidList Setter
// 标准酒店ID列表,逗号分割。与hid_list二者只能选一
func (r *AlitripTuanHotelAdaptStoreGetAPIRequest) SetShidList(_shidList []int64) error {
	r._shidList = _shidList
	r.Set("shid_list", _shidList)
	return nil
}

// GetShidList ShidList Getter
func (r AlitripTuanHotelAdaptStoreGetAPIRequest) GetShidList() []int64 {
	return r._shidList
}

// SetHidList is HidList Setter
// 物理酒店ID列表，逗号分割。与shid_list二者只能选一
func (r *AlitripTuanHotelAdaptStoreGetAPIRequest) SetHidList(_hidList []int64) error {
	r._hidList = _hidList
	r.Set("hid_list", _hidList)
	return nil
}

// GetHidList HidList Getter
func (r AlitripTuanHotelAdaptStoreGetAPIRequest) GetHidList() []int64 {
	return r._hidList
}

var poolAlitripTuanHotelAdaptStoreGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripTuanHotelAdaptStoreGetRequest()
	},
}

// GetAlitripTuanHotelAdaptStoreGetRequest 从 sync.Pool 获取 AlitripTuanHotelAdaptStoreGetAPIRequest
func GetAlitripTuanHotelAdaptStoreGetAPIRequest() *AlitripTuanHotelAdaptStoreGetAPIRequest {
	return poolAlitripTuanHotelAdaptStoreGetAPIRequest.Get().(*AlitripTuanHotelAdaptStoreGetAPIRequest)
}

// ReleaseAlitripTuanHotelAdaptStoreGetAPIRequest 将 AlitripTuanHotelAdaptStoreGetAPIRequest 放入 sync.Pool
func ReleaseAlitripTuanHotelAdaptStoreGetAPIRequest(v *AlitripTuanHotelAdaptStoreGetAPIRequest) {
	v.Reset()
	poolAlitripTuanHotelAdaptStoreGetAPIRequest.Put(v)
}
