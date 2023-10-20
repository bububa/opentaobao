package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahtorderhotelsyncbookingAPIRequest 未来酒店亲橙客栈预订信息同步 API请求
// alibaba.htorder.hotel.sync.booking
//
// 未来酒店亲橙客栈预订信息同步
type AlibabahtorderhotelsyncbookingAPIRequest struct {
	model.Params
	// 预订信息数据
	_dataEntity *SyncHotelBookingDataRequestDto
}

// NewAlibabahtorderhotelsyncbookingRequest 初始化AlibabahtorderhotelsyncbookingAPIRequest对象
func NewAlibabahtorderhotelsyncbookingRequest() *AlibabahtorderhotelsyncbookingAPIRequest {
	return &AlibabahtorderhotelsyncbookingAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahtorderhotelsyncbookingAPIRequest) GetApiMethodName() string {
	return "alibaba.htorder.hotel.sync.booking"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahtorderhotelsyncbookingAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahtorderhotelsyncbookingAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDataEntity is DataEntity Setter
// 预订信息数据
func (r *AlibabahtorderhotelsyncbookingAPIRequest) SetDataEntity(_dataEntity *SyncHotelBookingDataRequestDto) error {
	r._dataEntity = _dataEntity
	r.Set("data_entity", _dataEntity)
	return nil
}

// GetDataEntity DataEntity Getter
func (r AlibabahtorderhotelsyncbookingAPIRequest) GetDataEntity() *SyncHotelBookingDataRequestDto {
	return r._dataEntity
}
