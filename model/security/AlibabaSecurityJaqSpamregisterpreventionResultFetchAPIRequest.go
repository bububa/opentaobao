package security

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest
获取垃圾注册防控结果 API请求
alibaba.security.jaq.spamregisterprevention.result.fetch

获取垃圾注册防控结果 */
type AlibabaSecurityJaqSpamregisterpreventionResultFetchAPIRequest struct {
	model.Params
	// 账号绑定的手机号码
	_phoneNumber string
	// 账号绑定的IP地址
	_ip string
	// 场景描述，String必须在下面可以使用的场景中选择。可以使用的场景为：Login：登录事件；Register：注册事件；Trade:交易事件；Payment:支付事件;Refund：退款事件	；Loan：借款事件；Transfer：转账事件；Withdraw	：提现事件；Modify：修改事件；Click：点击事件；Activate：激活事件；	Other：其他事件。
	_context string
	// 登录来源。可以输入的参数如下：1：PC网页；2：移动网页；3：APP；4：其他
	_source int64
	// 账号绑定的email地址
	_email string
	// 账号的全局唯一标识
	_userId string
	// 当前注册的页面URL，Source为1，2时，该参数必选
	_registerUrl string
	// agent，发送HTTP请求的代理
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
	_nickName string
	// 账号绑定的公司名字
	_companyName string
	// 账号绑定的注册的地址
	_address string
	// 账号绑定的身份证号
	_idNumber string
	// 账号绑定的银行卡号
	_bankCardNumber string
	// 接入JS后从服务端获取的token
	_jsToken string
	// 接入SDK后从服务端获取的token
	_sdkToken string
}

// New
