package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHtorderHotelSyncConfigAPIRequest
同步配置信息 API请求
alibaba.htorder.hotel.sync.config

同步配置信息 */
type AlibabaHtorderHotelSyncConfigAPIRequest struct {
	model.Params
	// 配置信息
	_dataEntity *HotelMessageConfigDto
}

// NewAlibabaHtorderHotelSyncConfigRequest 初始化AlibabaHtorderHotelSyncConfigAPIRequest对象
func NewAlibabaHtorderHotelSyncConfigRequest() *AlibabaHtorderHotelSyncConfigAPIRequest {
	return &AlibabaHtorderHotelSyncConfigAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHtorderHotelSyncConfigAPIRequest) GetApiMethodName() string {
	return "alibaba.htorder.hotel.sync.config"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHtorderHotelSyncConfigAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is DataEntity Setter
// 配置信息
func (r *AlibabaHtorderHotelSyncConfigAPIRequest) SetDataEntity(_dataEntity *HotelMessageConfigDto) error {
	r._dataEntity = _dataEntity
	r.Set("data_entity", _dataEntity)
	return nil
}

// Get DataEntity Getter
func (r AlibabaHtorderHotelSyncConfigAPIRequest) GetDataEntity() *HotelMessageConfigDto {
	return r._dataEntity
}
