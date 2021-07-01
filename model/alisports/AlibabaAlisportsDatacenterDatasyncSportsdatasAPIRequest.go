package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest
阿里体育运动数据同步统一接口 API请求
alibaba.alisports.datacenter.datasync.sportsdatas

给单方提供同步运动数据到阿里体育的接口 */
type AlibabaAlisportsDatacenterDatasyncSportsdatasAPIRequest struct {
	model.Params
	// 用户阿里体育id
	_userId string
	// 运动一级分类
	_sportsCat1Id int64
	// 运动二级分类
	_sportsCat2Id int64
	// 运动三级分类
	_sportsCat3Id string
	// 运动开始时间，单位：毫秒
	_sportsStartTime int64
	// 运动结束时间，单位：毫秒
	_sportsEndTime int64
	// 时区
	_timezone int64
	// 最小心率
	_minHeartrate int64
	// 最大心率
	_maxHeartrate int64
	// 平均心率
	_avgHeartrate int64
	// 速度，单位：千米/小时
	_speed string
	// 动作计数，如：步数、滑水次数
	_actionCount string
	// 路径
	_path string
	// 数据原始来源
	_subChannel string
	// 里程，单位：米
	_mileage int64
	// 爬高，单位：米
	_climb int64
	// 运动持续时间，单位：毫秒
	_durationTime int64
	// 开始位置，格式：经度,维度
	_startPoint string
	// 预留字段
	_resultOther string
	// 最大速度，单位：千米/小时
	_maxSpeed string
	// 结束位置,格式[经度,纬度]
	_endPoint string
	// 过程数据Json
	_stage string
	// 频率
	_powerFrequency int64
	// 消耗卡路里，单位：千卡
	_calorie string
	// 路径节点数据下载地址
	_pathDataUrl string
	// 过程数据下载地址
	_stageDataUrl string
	// 数据类型：0.普通数据 1.赛事数据
	_dataType int64
	// 设备类型
	_deviceType int64
	// 设备型号(厂商)
	_deviceModel string
	// 设备名称
	_deviceName string
	// 三方数据唯一码
	_messageId string
	// 版本号
	_version string
}

// New
