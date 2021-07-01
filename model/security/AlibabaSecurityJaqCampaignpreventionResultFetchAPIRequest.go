package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest
获取活动保护结果 API请求
alibaba.security.jaq.campaignprevention.result.fetch

获取活动保护结果 */
type AlibabaSecurityJaqCampaignpreventionResultFetchAPIRequest struct {
	model.Params
	// 电话号码。手机号码的格式为国家码-手机号的格式，如中国手机号86-13088889999或美国手机号001-9096185426，如果不加国家码都视为中国手机号码【phone_number ,email, (user_id,id_type)三种必选其一】
	_phoneNumber string
	// 邮箱地址。【phone_number ,email, (user_id,id_type)三种必选其一】
	_email string
	// 账号的全局唯一标识。【phone_number ,email, (user_id,id_type)三种必选其一】
	_userId string
	// 有user_id时候必填。1：阿里HID；2：用户自有ID；3：openId; 4:其它。【phone_number ,email, (user_id,id_type)三种必选其一】
	_idType int64
	// 登录时候的IP地址
	_ip string
	// 登录来源。1：PC网页；2：移动网页；3：APP;4:其它
	_source int64
	// 当前操作的页面URL。Source为1，2时，该参数必选
	_currentUrl string
	// 发送HTTP请求的代理
	_agent string
	// Cookie
	_cookie string
	// Session id
	_sessionId string
	// 硬件信息
	_macAddress string
	// 上一跳信息
	_referer string
	// 账号绑定的呢称
	_userName string
	// 账号绑定的公司名字
	_companyName string
	// 账号绑定的注册的地址
	_address string
	// 账号绑定的身份证号
	_idNumber string
	// 账号绑定的银行卡号
	_bankCardNumber string
	// 接入JS后获取的token
	_jsToken string
	// 接入无线保镖安全组件后从服务端获取的token
	_sdkToken string
	// 协议版本号。现在的值是1.0
	_protocolVersion string
	// 扩展字段。json格式的字符串，根据具体情况而定 。
	_extendData string
	// 活动描述，场景为活动时提供，活动基本规则描述，以评估活动流程的风险
	_activityDescription string
	// 活动编号
	_activityId string
	// 活动奖品面值，场景为活动时提供，活动抽中的价值
	_prize string
	// 活动奖品类型。1：事物；2：红包；3：优惠券；4：购物券；5：积分；6：代币；0：其它
	_prizeType int64
	// 注册的时间（秒）
	_registerDate int64
	// 注册时候的ip
	_registerIp string
}

// New
