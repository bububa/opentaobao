package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlisportsDataSportsSyncmatchdataAPIRequest
阿里体育数据中心用户比赛数据同步接口 API请求
alibaba.alisports.data.sports.syncmatchdata

阿里体育数据中心用户比赛数据同步接口 */
type AlibabaAlisportsDataSportsSyncmatchdataAPIRequest struct {
	model.Params
	// 应用appkey
	_alispAppKey string
	// 成绩(比赛用时)
	_score string
	// 组别 1001半程马拉松  1002全程马拉松
	_matchGroup int64
	// 国家
	_country string
	// 出生日期 格式：Y-m-d
	_birthday string
	// 手机号
	_mobile string
	// 证件ID
	_cardId string
	// 证件类型 1身份证 2军官证 4护照 8台胞证 16港澳通行证 32未设置  64 其他
	_cardType int64
	// 类型：1专业 2业余
	_type int64
	// 性别 0未知 1男 2女
	_gender int64
	// 姓名
	_name string
	// 排名
	_ranking int64
	// 比赛名（展示用）
	_match string
	// 比赛类型 1马拉松
	_matchType int64
	// 参赛号
	_num string
	// 阿里体育用户id
	_aliuid string
	// 接口签名
	_alispSign string
	// 时间戳精确到秒
	_alispTime string
	// 比赛日期 格式：Y-m-d
	_matchTime string
	// 枪声成绩
	_gunshotScore string
	// 枪声排名
	_gunshotRanking int64
	// 平均配速
	_speed string
	// 展示用，比如：男子半程马拉松
	_project string
	// 比如马拉松 5KM 10KM 15KM分段成绩，json key->value 格式
	_subScore string
	// 比如马拉松 5KM 10KM 15KM分段时间点，json key->value 格式
	_subPoint string
}

// New
