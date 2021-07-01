package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTbkDgVegasTljCreateAPIRequest
淘宝客-推广者-淘礼金创建 API请求
taobao.tbk.dg.vegas.tlj.create

创建淘礼金 */
type TaobaoTbkDgVegasTljCreateAPIRequest struct {
	model.Params
	// CPS佣金计划类型
	_campaignType string
	// 妈妈广告位Id
	_adzoneId int64
	// 宝贝id
	_itemId int64
	// 淘礼金总个数
	_totalNum int64
	// 淘礼金名称，最大10个字符
	_name string
	// 单用户累计中奖次数上限
	_userTotalWinNumLimit int64
	// 安全开关，若不进行安全校验，可能放大您的资损风险，请谨慎选择
	_securitySwitch bool
	// 单个淘礼金面额，支持两位小数，单位元
	_perFace string
	// 发放开始时间
	_sendStartTime string
	// 发放截止时间
	_sendEndTime string
	// 使用结束日期。如果是结束时间模式为相对时间，时间格式为1-7直接的整数, 例如，1（相对领取时间1天）； 如果是绝对时间，格式为yyyy-MM-dd，例如，2019-01-29，表示到2019-01-29 23:59:59结束
	_useEndTime string
	// 结束日期的模式,1:相对时间，2:绝对时间
	_useEndTimeMode int64
	// 使用开始日期。相对时间，无需填写，以用户领取时间作为使用开始时间。绝对时间，格式 yyyy-MM-dd，例如，2019-01-29，表示从2019-01-29 00:00:00 开始
	_useStartTime string
	// 安全等级，0：适用于常规淘礼金投放场景；1：更高安全级别，适用于淘礼金面额偏大等安全性较高的淘礼金投放场景，可能导致更多用户拦截。security_switch为true，此字段不填写时，使用0作为默认安全级别。如果security_switch为false，不进行安全判断。
	_securityLevel int64
}

// New
