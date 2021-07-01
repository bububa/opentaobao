package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlisportsDataSportsSyncstatdataAPIRequest
阿里体育数据中心用户当天累积数据同步接口 API请求
alibaba.alisports.data.sports.syncstatdata

阿里体育数据中心用户当天累积数据同步接口 */
type AlibabaAlisportsDataSportsSyncstatdataAPIRequest struct {
	model.Params
	// 应用appkey
	_alispAppKey string
	// 时间戳精确到秒
	_alispTime string
	// 签名
	_alispSign string
	// 阿里体育用户id
	_aliuid string
	// 运动步数
	_steps string
	// 消耗卡路里 单位：卡
	_calorie string
	// 运动距离 单位：米
	_distance string
	// 日期 格式：y-m-d h:i:s
	_time string
	// 设备类型：1手环
	_deviceType string
	// 设备名
	_deviceName string
	// 设备型号
	_deviceModel string
}

// New
