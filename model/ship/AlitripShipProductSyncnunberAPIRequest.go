package ship

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripShipProductSyncnunberAPIRequest
船票班次变更回调 API请求
alitrip.ship.product.syncnunber

船票班次变更回调 */
type AlitripShipProductSyncnunberAPIRequest struct {
	model.Params
	// 出发城市
	_cityName string
	// 出发城市code
	_cityCode string
	// 出发港口
	_fromStationName string
	// 出发港口编号
	_fromStationCode string
}

// New
