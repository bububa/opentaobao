package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest
获取登录保护结果 API请求
alibaba.security.jaq.loginprevention.result.fetch

获取登录保护结果 */
type AlibabaSecurityJaqLoginpreventionResultFetchAPIRequest struct {
	model.Params
	// 关联账号的手机号。手机号码的格式为国家码-手机号的格式，如中国手机号86-13088889999或美国手机号001-9096185426，如果不加国家码都视为中国手机号码【phone_number ,email, (user_id,id_type)三种必选其一】
	_phoneNumber string
	// 关联账号的email。【phone_number ,email, (user_id,id_type)三种必选其一】
	_email string
	// 账号的全局唯一标识，为了提高准确率，建议带上该字段。【phone_number ,email, (user_id,id_type)三种必选其一】
	_userId string
	// 账号的全局唯一标识，为了提高准确率，建议带上该字段。【phone_number ,email, (user_id,id_type)三种必选其一】
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
	// 账号在系统里面是否存在。0：不存在；1：存在
	_accountExist int64
	// 登录场景。1：账密登陆；2：扫码登录；3：短信验证码登录；0：其它
	_loginType int64
	// 密码是否正确。0：不正确；1：正确
	_passwordCorrect int64
	// 将密码加盐hash后传递，用于弱密码检测
	_passwordHash string
	// 注册的时间（秒）
	_registerDate int64
	// 注册时候的ip
	_registerIp string
}

// New
