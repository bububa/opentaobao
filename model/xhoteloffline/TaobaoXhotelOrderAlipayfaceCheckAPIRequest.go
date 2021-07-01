package xhoteloffline

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelOrderAlipayfaceCheckAPIRequest
线下信用住买家资格校验接口 API请求
taobao.xhotel.order.alipayface.check

接口用于校验买家是否具有使用酒店信用住的资格 */
type TaobaoXhotelOrderAlipayfaceCheckAPIRequest struct {
	model.Params
	// 总的收费金额，单位为分
	_totalFee int64
	// 参数必填，发布到阿里旅行的酒店编码
	_hotelCode string
	// 证件号, 如果加密方式设置为1, 传入加密后的证件号
	_idNumber string
	// 加密方式, 默认0: 不加密, 信息会通过淘宝开放平台传输, 阿里旅行可以获取到具体信息;  1: SHA-1不可逆加密,  阿里旅行方面无法解析到具体信息, 只用于做信息匹配.注意加密后生成40字节长度的字符串
	_encryptType int64
	// 验证类型.可以不设置. 默认0-信用住下单资格校验;1-已经创建的信用住订单,入住人在checkIn时的资格复审
	_type int64
	// 证件类型, 默认0:身份证; 1: 护照; 2:警官证; 3:士兵证; 4: 回乡证
	_idType int64
	// 不清楚请留空, 用于和outHid共同定位一个酒店
	_vendor string
	// 入住人姓名
	_guestName string
	// 客人手机号
	_mobileNo string
}

// New
