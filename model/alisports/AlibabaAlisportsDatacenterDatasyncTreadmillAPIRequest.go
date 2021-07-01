package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlisportsDatacenterDatasyncTreadmillAPIRequest
阿里体育同步跑步机设备数据 API请求
alibaba.alisports.datacenter.datasync.treadmill

合作方向阿里体育同步跑步机设备的数据 */
type AlibabaAlisportsDatacenterDatasyncTreadmillAPIRequest struct {
	model.Params
	// 阿里体育用户id
	_userId string
	// 运动开始时间,秒级别时间戳
	_startTime int64
	// 运动结束时间，秒级别时间戳
	_endTime int64
	// 时区编码，不传默认东八区
	_timezone string
	// 运动位置经纬度
	_location string
	// 国家编码，https://zh.wikipedia.org/wiki/%E5%9C%8B%E5%AE%B6%E5%9C%B0%E5%8D%80%E4%BB%A3%E7%A2%BC
	_countryCode int64
	// 省编码， https://www.ip33.com/area_code.html
	_provinceCode int64
	// 城市编码
	_cityCode int64
	// 平均速度，单位km/h
	_speed *BigDecimal
	// 最大速度，单位km/h
	_maxSpeed *BigDecimal
	// 平均步频
	_powerFrequency int64
	// 累计里程，单位：m
	_mileage int64
	// 累计爬升，单位m
	_climb int64
	// 运动持续时常,单位：秒
	_durationTime int64
	// 总步数
	_steps int64
	// 消耗总热量，单位：卡路里
	_calorie *BigDecimal
	// 最大心率
	_maxHeartrate int64
	// 最小心率
	_minHeartrate int64
	// 平均心率
	_avgHeartrate int64
	// 过程数据收集间隔时间
	_collectTimeInterval int64
	// 过程数据收集间隔时间单位，1.毫秒 2.秒 3.分 4.小时
	_collectTimeUnit int64
	// 速度过程数据,单位km/h
	_speedDatas []string
	// 步频/踏频/桨频过程数据
	_motionFrequencyDatas []string
	// 步幅/踏幅/桨幅过程数据
	_hrzMotionRangeDatas []string
	// 配速过程数据
	_tempoDatas []string
	// 心率过程数据
	_heartrateDatas []int64
	// 三方数据唯一id
	_messageId string
	// 设备类型：4.跑步机 5.单车 6.划船机
	_deviceType int64
	// 设备名称
	_deviceName string
	// 设备型号
	_deviceModel string
}

// New
