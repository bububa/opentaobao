package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlisportsDataSportsSyncsleepdataAPIRequest
阿里体育数据中心用户睡眠数据同步接口 API请求
alibaba.alisports.data.sports.syncsleepdata

阿里体育数据中心用户睡眠数据同步接口 */
type AlibabaAlisportsDataSportsSyncsleepdataAPIRequest struct {
	model.Params
	// 应用appkey
	_alispAppKey string
	// 入睡时间，格式：y-m-d h:i:s
	_stime string
	// 清醒时间，单位：小时
	_soberTime string
	// 浅度睡眠时间，单位：小时
	_shallowTime string
	// 深度睡眠时间，单位：小时
	_deepTime string
	// 睡眠总时间，单位：小时
	_allTime string
	// 阿里体育用户id
	_aliuid string
	// 接口签名
	_alispSign string
	// 时间戳精确到秒
	_alispTime string
	// 醒来时间，格式：y-m-d h:i:s
	_etime string
}

// New
