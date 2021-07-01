package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlisportsDataSportsSyncuserdataAPIRequest
阿里体育数据中心用户个人信息同步接口 API请求
alibaba.alisports.data.sports.syncuserdata

阿里体育数据中心用户个人信息同步接口 */
type AlibabaAlisportsDataSportsSyncuserdataAPIRequest struct {
	model.Params
	// 应用appkey
	_alispAppKey string
	// 基础代谢率
	_metabolize string
	// 蛋白质含量
	_protein string
	// 骨量
	_bone string
	// 水分率
	_water string
	// 肌肉率
	_muscle string
	// 体脂率
	_fat string
	// 静息心率，单位：次/每分
	_heartRate int64
	// 体重，单位kg
	_weight string
	// 身高，单位m
	_height string
	// 年龄
	_age int64
	// 三方主键id，唯一标识数据
	_dataId string
	// 阿里体育用户id
	_aliuid string
	// 接口签名
	_alispSign string
	// 时间戳精确到秒
	_alispTime string
	// 日期 格式：y-m-d h:i:s
	_time string
}

// New
