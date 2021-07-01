package tuanhotel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripTuanHotelAdaptStoreGetAPIRequest
酒店团购套餐关联适用门店 API请求
alitrip.tuan.hotel.adapt.store.get

输入shid，返回关联门店详情信息 */
type AlitripTuanHotelAdaptStoreGetAPIRequest struct {
	model.Params
	// 标准酒店ID列表,逗号分割。与hid_list二者只能选一
	_shidList []int64
	// 物理酒店ID列表，逗号分割。与shid_list二者只能选一
	_hidList []int64
}

// New
