package tbrefund

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaorprefundsagreeAPIRequest 同意退款 API请求
// taobao.rp.refunds.agree
//
// 卖家同意退款，支持批量退款，只允许子账号操作。淘宝退款一次最多能退20笔，总金额不超过6000元；天猫退款一次最多能退30笔，总金额不超过10000元。
type TaobaorprefundsagreeAPIRequest struct {
	model.Params
	// 短信验证码，如果退款金额达到一定的数量，后端会返回调用失败，并同时往卖家的手机发送一条短信验证码。接下来用收到的短信验证码再次发起API调用即可完成退款操作。
	_code string
	// 退款信息，格式：refund_id|amount|version|phase，其中refund_id为退款编号，amount为退款金额（以分为单位），version为退款最后更新时间（时间戳格式），phase为退款阶段（可选值为：onsale, aftersale，天猫退款必值，淘宝退款不需要传），多个退款以半角逗号分隔。
	_refundInfos string
}

// NewTaobaorprefundsagreeRequest 初始化TaobaorprefundsagreeAPIRequest对象
func NewTaobaorprefundsagreeRequest() *TaobaorprefundsagreeAPIRequest {
	return &TaobaorprefundsagreeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaorprefundsagreeAPIRequest) GetApiMethodName() string {
	return "taobao.rp.refunds.agree"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaorprefundsagreeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaorprefundsagreeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCode is Code Setter
// 短信验证码，如果退款金额达到一定的数量，后端会返回调用失败，并同时往卖家的手机发送一条短信验证码。接下来用收到的短信验证码再次发起API调用即可完成退款操作。
func (r *TaobaorprefundsagreeAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r TaobaorprefundsagreeAPIRequest) GetCode() string {
	return r._code
}

// SetRefundInfos is RefundInfos Setter
// 退款信息，格式：refund_id|amount|version|phase，其中refund_id为退款编号，amount为退款金额（以分为单位），version为退款最后更新时间（时间戳格式），phase为退款阶段（可选值为：onsale, aftersale，天猫退款必值，淘宝退款不需要传），多个退款以半角逗号分隔。
func (r *TaobaorprefundsagreeAPIRequest) SetRefundInfos(_refundInfos string) error {
	r._refundInfos = _refundInfos
	r.Set("refund_infos", _refundInfos)
	return nil
}

// GetRefundInfos RefundInfos Getter
func (r TaobaorprefundsagreeAPIRequest) GetRefundInfos() string {
	return r._refundInfos
}
