package alihealthcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthUicSwipefaceSyncdataAPIRequest
刷脸测睡眠数据同步 API请求
alibaba.alihealth.uic.swipeface.syncdata

刷脸测睡眠数据同步，三方数据回传 */
type AlibabaAlihealthUicSwipefaceSyncdataAPIRequest struct {
	model.Params
	// 用户ID
	_tpUserId int64
	// 缺觉小时数
	_lackSleepHours int64
	// 黑圆圈级别
	_blackEyeLevel int64
	// 眼袋级别
	_eyeBagSwollenLevel int64
	// 鱼尾纹数
	_fishTail int64
	// 嘴唇颜色
	_lipColor string
}

// New
