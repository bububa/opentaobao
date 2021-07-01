package alihealthcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaFmhealthWeightLossplanSynclossplanAPIRequest
减重计划--同步减重计划 API请求
alibaba.fmhealth.weight.lossplan.synclossplan

减重计划--三方同步用户初始化减重计划给我们 */
type AlibabaFmhealthWeightLossplanSynclossplanAPIRequest struct {
	model.Params
	// 用户id
	_tpUserId int64
	// 性别0女 1男
	_gender int64
	// 生日
	_birthday string
	// 身高170 即一米七
	_height int64
	// 当前体重（今天的体重），单位kg
	_weight string
	// 0创建减肥计划调用；1修改调用；
	_type int64
	// 体年龄
	_bodyAge int64
	// 完成时间
	_finishDate string
	// 每周减重
	_lossPerWeek string
	// 目标体重
	_weightGoal string
	// 减重类型0保持 1减肥
	_lossLevel int64
	// 每日可以摄入的标准总量
	_totalCalorie int64
	// 减重计划开始时间
	_beginDate string
}

// New
