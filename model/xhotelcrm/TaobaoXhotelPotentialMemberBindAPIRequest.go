package xhotelcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelPotentialMemberBindAPIRequest
飞猪酒店商家会员绑定 API请求
taobao.xhotel.potential.member.bind

支持互通商家发起会员绑定 */
type TaobaoXhotelPotentialMemberBindAPIRequest struct {
	model.Params
	// 名
	_firstName string
	// 姓
	_lastName string
	// 电话
	_phone string
	// 邮箱
	_email string
	// 卡号
	_cardNo string
	// 等级(V1,V2,V3)
	_grade string
	// 注册时间
	_registerDate string
	// 生效时间
	_fromDate string
	// 截止时间
	_toDate string
	// 性别(M,F,U-未知)
	_sex string
	// 城市
	_city string
	// 年龄
	_age int64
	// 籍贯
	_nativePlace string
}

// New
