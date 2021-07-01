package retail

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractRetailSaveshelflocationAPIRequest
保存地理位置和货架关系 API请求
alibaba.interact.retail.saveshelflocation

保存地理位置和货架关系 */
type AlibabaInteractRetailSaveshelflocationAPIRequest struct {
	model.Params
	// 门店code
	_storeCode string
	// 货架编号
	_shelfNo string
	// 经度
	_lng string
	// 纬度
	_lat string
	// POI
	_poiId string
	// 地址
	_address string
}

// New
