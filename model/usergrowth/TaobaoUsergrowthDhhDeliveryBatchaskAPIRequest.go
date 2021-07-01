package usergrowth

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUsergrowthDhhDeliveryBatchaskAPIRequest
广告曝光前判定批量接口V2 API请求
taobao.usergrowth.dhh.delivery.batchask

广告曝光前判定批量接口V2 */
type TaobaoUsergrowthDhhDeliveryBatchaskAPIRequest struct {
	model.Params
	// md5加密后的oaid列表， 32位小写， 多个使用,分隔， 最多支持20个
	_oaidMd5 string
	// md5加密后的oaid列表， 32位小写， 多个使用,分隔， 最多支持20个
	_idfaMd5 string
	// md5加密后的imei列表， 32位小写， 多个使用,分隔， 最多支持20个
	_imeiMd5 string
	// 巨浪广告位id,在巨浪平台申请
	_advertisingSpaceId string
	// 巨浪渠道id,在巨浪平台申请
	_channel string
}

// New
