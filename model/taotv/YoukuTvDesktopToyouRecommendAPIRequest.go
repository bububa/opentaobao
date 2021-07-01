package taotv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YoukuTvDesktopToyouRecommendAPIRequest
TV桌面为你推荐接口 API请求
youku.tv.desktop.toyou.recommend

提供为你推荐数据 */
type YoukuTvDesktopToyouRecommendAPIRequest struct {
	model.Params
	// 用户登陆token
	_token string
	// 播控方，1:华数 7:CIBN
	_bcp string
	// 终端设备型号
	_deviceModel string
	// 桌面版本号
	_versionCode int64
	// 终端设备mac
	_mac string
	// 终端设备uuid
	_uuid string
	// 影视来源，0-淘tv 7-优酷免费 9-优酷付费 多项用“,”隔开
	_from string
	// 支持收费方式，0-免费 1-限免 2-单点 3-包月 4-红包 5-单包,多项用“,”隔开
	_chargeType string
	// 获取的最大节目数
	_maxSize int64
	// 分辨率，sw720 sw1080
	_sw string
	// 支持媒体类型,dts,drm,dolby,h265
	_deviceMedia string
	// 请求IP
	_ip string
}

// New
