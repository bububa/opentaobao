package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest
阿里体育接入体脂秤数据 API请求
alibaba.alisports.datacenter.datasync.fatscaledata

阿里体育数据中心接入体脂秤数据 */
type AlibabaAlisportsDatacenterDatasyncFatscaledataAPIRequest struct {
	model.Params
	// 阿里体育用户id
	_aliuid string
	// 测量时间，秒级别时间戳
	_time int64
	// 年龄
	_age int64
	// 身高，单位cm
	_height int64
	// 体重，单位kg
	_weight *BigDecimal
	// bmi
	_bmi *BigDecimal
	// 基础代谢率,单位卡
	_basalMetabolicRate *BigDecimal
	// 去脂体重，单位kg
	_fatFreeMass *BigDecimal
	// 体脂率，百分比12.4%，传12.4
	_bodyFatRate *BigDecimal
	// 脂肪重量，单位kg
	_fatMass *BigDecimal
	// 皮下脂肪率，百分比
	_subcutaneousFatRate *BigDecimal
	// 内脏脂肪率，百分比
	_visceralFatIndex *BigDecimal
	// 肌肉率，百分比
	_muscleRate *BigDecimal
	// 肌肉重量，单位kg
	_muscleMass *BigDecimal
	// 骨骼肌率，百分比
	_skeletalMuscleRate *BigDecimal
	// 水分率，百分比
	_moistureRate *BigDecimal
	// 蛋白质率，百分比
	_proteinRate *BigDecimal
	// 骨量，单位kg
	_boneMass *BigDecimal
	// 体重指数
	_weightIndex int64
	// 身体年龄
	_bodyAge int64
	// 设备类型：1.体脂秤，2智能手表，3智能手环
	_deviceType int64
	// 设备名称
	_deviceName string
	// 设备编号
	_deviceModel string
	// 三方唯一id
	_messageId string
}

// New
