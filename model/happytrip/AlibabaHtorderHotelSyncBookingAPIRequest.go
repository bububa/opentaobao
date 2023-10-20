package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHtorderHotelSyncBookingAPIRequest 未来酒店亲橙客栈预订信息同步 API请求
// alibaba.htorder.hotel.sync.booking
//
// 未来酒店亲橙客栈预订信息同步
type AlibabaHtorderHotelSyncBookingAPIRequest struct {
	model.Params
	// 预订信息数据
	_dataEntity *SyncHotelBookingDataRequestDto
}

// NewAlibabaHtorderHotelSyncBookingRequest 初始化AlibabaHtorderHotelSyncBookingAPIRequest对象
func NewAlibabaHtorderHotelSyncBookingRequest() *AlibabaHtorderHotelSyncBookingAPIRequest {
	return &AlibabaHtorderHotelSyncBookingAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHtorderHotelSyncBookingAPIRequest) GetApiMethodName() string {
	return "alibaba.htorder.hotel.sync.booking"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHtorderHotelSyncBookingAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHtorderHotelSyncBookingAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDataEntity is DataEntity Setter
// 预订信息数据
func (r *AlibabaHtorderHotelSyncBookingAPIRequest) SetDataEntity(_dataEntity *SyncHotelBookingDataRequestDto) error {
	r._dataEntity = _dataEntity
	r.Set("data_entity", _dataEntity)
	return nil
}

// GetDataEntity DataEntity Getter
func (r AlibabaHtorderHotelSyncBookingAPIRequest) GetDataEntity() *SyncHotelBookingDataRequestDto {
	return r._dataEntity
}
