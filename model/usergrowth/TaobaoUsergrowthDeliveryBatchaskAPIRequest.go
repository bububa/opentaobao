package usergrowth

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUsergrowthDeliveryBatchaskAPIRequest
广告投放批量询问 API请求
taobao.usergrowth.delivery.batchask

提供给媒体在曝光广告前调用， 返回是否曝光以及报价 */
type TaobaoUsergrowthDeliveryBatchaskAPIRequest struct {
	model.Params
	// 渠道标识，向淘宝技术申请
	_channel string
	// 广告id，淘宝和媒体协商
	_adid string
	// idfa的md5值， 32位小写， 多个使用,分隔， 最多支持20个
	_idfaMd5 string
	// imei的md5值， 32位小写， 多个之间使用,分隔， 最多支持20个
	_imeiMd5 string
}

// New
