package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlisportsDataSportsSyncsportsdataAPIRequest
阿里体育数据中心用户运动数据同步接口 API请求
alibaba.alisports.data.sports.syncsportsdata

阿里体育数据中心用户运动数据同步接口 */
type AlibabaAlisportsDataSportsSyncsportsdataAPIRequest struct {
	model.Params
	// 应用appkey
	_alispAppKey string
	// 最高速度 单位：米/每分
	_maxSpeed string
	// 平均心率 单位：次/每分
	_averHeartrate int64
	// 最高心率 单位：次/每分
	_maxHeartrate int64
	// 最低心率 单位：次/每分
	_minHeartrate int64
	// 运动轨迹，有序的经纬度集合，json格式 例：[[1,2],[3,4]] [1,2]是一个坐标点，1是经度 2是纬度（有就传，阿里体育较依赖此字段）
	_trail string
	// 运动开始时间（如果不区分开始结束，两个时间值相同;格式：Y-m-d H:i:s）
	_stime string
	// 运动结束时间（如果不区分开始结束，两个时间值相同;格式：Y-m-d H:i:s）
	_etime string
	// 设备类型 :1.手环;2.手表;3.跑步机;4.智能运动鞋;5.耳机;6.智能运动鞋;7.智能运动Bra8.智能单车;9.智能跳绳10.智能背心11.脚环
	_deviceType int64
	// 设备名（展示会用到）
	_deviceName string
	// 设备型号
	_deviceModel string
	// 平均速度 单位：米/每分
	_averSpeed string
	// 二级运动量单位 定义：1.爬楼层数(跑步、健走、健身、登山);2.划水次数(游泳)
	_subUnit int64
	// 二级运动量
	_subNum string
	// 时间戳精确到秒
	_alispTime string
	// 接口签名
	_alispSign string
	// 阿里体育用户id
	_aliuid string
	// 业务来源二级分类（中英文）
	_source string
	// 三方运动数据主键id（数据唯一标记，去重使用）
	_dataId string
	// 运动类型一级分类 定义：1-跑步;2-健走;3-骑行;4-游泳;5-健身;6-篮球;7-足球;8-羽毛球;9-排球;10-乒乓球;11-瑜伽;12-电竞;13-登山;16-椭圆机;19-健身操;20-太极;
	_sportsClass int64
	// 运动类型二级分类 定义： 1001-室内跑步;1002-室外跑步;2001-室内健走;2002-室外健走;3001-室内骑行;3002-室外骑行;4001-室内游泳;4002-户外游泳
	_sportsType int64
	// 运动量
	_num string
	// 运动量单位 1.步数(跑步、健走、椭圆机、登山);2.趟数(游泳);3.平均踏频(骑行);
	_unit int64
	// 运动消耗卡路里 单位：卡
	_calorie string
	// 所属赛事（有就传，阿里体育较依赖此字段）
	_match string
	// 运动距离 单位:米（有就传，阿里体育较依赖此字段）
	_distance string
	// 运动时长 单位:秒（有就传，阿里体育较依赖此字段）
	_time string
	// 国家(中英文/运动发生地点，如有就传)
	_country string
	// 省份(中英文/运动发生地点，如有就传)
	_province string
	// 城市(中英文/运动发生地点，如有就传)
	_city string
	// 开始运动地点经纬度，格式：1,2 (1为经度，2为纬度)
	_startPoint string
	// 结束运动地点经纬度，格式：1,2 (1为经度，2为纬度)
	_endPoint string
}

// New
