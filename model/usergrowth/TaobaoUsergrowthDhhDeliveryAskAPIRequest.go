package usergrowth

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUsergrowthDhhDeliveryAskAPIRequest
广告曝光前判定接口V2 API请求
taobao.usergrowth.dhh.delivery.ask

提供给媒体在曝光广告前调用 */
type TaobaoUsergrowthDhhDeliveryAskAPIRequest struct {
	model.Params
	// 预留json参数，与手淘团队单独沟通
	_profile string
	// oaid的md5值， 32位小写
	_oaidMd5 string
	// idfa的md5值， 32位小写
	_idfaMd5 string
	// imei的md5值， 32位小写
	_imeiMd5 string
	// oaid的原生值
	_oaid string
	// idfa的原生值
	_idfa string
	// imei的原生值
	_imei string
	// 用户所使用设备的系统， 0： android, 1: ios, 2: windowsphone, 3: other
	_os string
	// 渠道标识，在大航海平台申请
	_channel string
	// 大航海广告位，在大航海平台申请
	_advertisingSpaceId string
}

// New
