package happytrip

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHtorderHotelSyncConfigAPIRequest 同步配置信息 API请求
// alibaba.htorder.hotel.sync.config
//
// 同步配置信息
type AlibabaHtorderHotelSyncConfigAPIRequest struct {
	model.Params
	// 配置信息
	_dataEntity *HotelMessageConfigDto
}

// NewAlibabaHtorderHotelSyncConfigRequest 初始化AlibabaHtorderHotelSyncConfigAPIRequest对象
func NewAlibabaHtorderHotelSyncConfigRequest() *AlibabaHtorderHotelSyncConfigAPIRequest {
	return &AlibabaHtorderHotelSyncConfigAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHtorderHotelSyncConfigAPIRequest) Reset() {
	r._dataEntity = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHtorderHotelSyncConfigAPIRequest) GetApiMethodName() string {
	return "alibaba.htorder.hotel.sync.config"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHtorderHotelSyncConfigAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHtorderHotelSyncConfigAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDataEntity is DataEntity Setter
// 配置信息
func (r *AlibabaHtorderHotelSyncConfigAPIRequest) SetDataEntity(_dataEntity *HotelMessageConfigDto) error {
	r._dataEntity = _dataEntity
	r.Set("data_entity", _dataEntity)
	return nil
}

// GetDataEntity DataEntity Getter
func (r AlibabaHtorderHotelSyncConfigAPIRequest) GetDataEntity() *HotelMessageConfigDto {
	return r._dataEntity
}

var poolAlibabaHtorderHotelSyncConfigAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHtorderHotelSyncConfigRequest()
	},
}

// GetAlibabaHtorderHotelSyncConfigRequest 从 sync.Pool 获取 AlibabaHtorderHotelSyncConfigAPIRequest
func GetAlibabaHtorderHotelSyncConfigAPIRequest() *AlibabaHtorderHotelSyncConfigAPIRequest {
	return poolAlibabaHtorderHotelSyncConfigAPIRequest.Get().(*AlibabaHtorderHotelSyncConfigAPIRequest)
}

// ReleaseAlibabaHtorderHotelSyncConfigAPIRequest 将 AlibabaHtorderHotelSyncConfigAPIRequest 放入 sync.Pool
func ReleaseAlibabaHtorderHotelSyncConfigAPIRequest(v *AlibabaHtorderHotelSyncConfigAPIRequest) {
	v.Reset()
	poolAlibabaHtorderHotelSyncConfigAPIRequest.Put(v)
}
