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

// New
