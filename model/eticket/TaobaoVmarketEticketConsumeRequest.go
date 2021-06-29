package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
电子票券消费通知 API请求
taobao.vmarket.eticket.consume

外部合作商家电子票券消费回调接口
*/
type TaobaoVmarketEticketConsumeRequest struct {
    model.Params
    // 进行验码的电子凭证订单的订单ID
    _orderId   int64
    // 核销的码，只支持单个码，多个码核销需要多次调用
    _verifyCode   string
    // 核销份数
    _consumeNum   int64
    // 安全验证token,需要和发码通知中的token一致
    _token   string
    // 码商ID,是码商的话必须传递,如果是信任卖家不需要传
    _codemerchantId   int64
    // 机具ID(此参数信任卖家可不传递，码商必须传递)
    _posid   string
    // 手机后四位(没有特殊说明请不要传该参数)
    _mobile   string
    // 核销后需要重新生成的码，如果不需要重新生成码，不要传该参数
    _newCode   string
    // 自定义核销流水号，如果核销调用失败，可以用该核销流水号进行冲正操作，需要小于等于100个字符(a-zA-Z0-9_)；每次核销都是唯一的流水号
    _serialNum   string
    // 不需要上传二维码图片或者核销后不需重新生成码码商请不要传，需要传入二维码的码商请先调用taobao.vmarket.eticket.qrcode.upload接口，将返回的img_filename文件名称作为参数（如果二维码不变的话，也可将将发码时传入二维码文件名作为参数传入），文件名与参数new_code必须相互对应。
    _qrImages   string
}

// 初始化TaobaoVmarketEticketConsumeRequest对象
func NewTaobaoVmarketEticketConsumeRequest() *TaobaoVmarketEticketConsumeRequest{
    return &TaobaoVmarketEticketConsumeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoVmarketEticketConsumeRequest) GetApiMethodName() string {
    return "taobao.vmarket.eticket.consume"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoVmarketEticketConsumeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 进行验码的电子凭证订单的订单ID
func (r *TaobaoVmarketEticketConsumeRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r TaobaoVmarketEticketConsumeRequest) GetOrderId() int64 {
    return r._orderId
}
// VerifyCode Setter
// 核销的码，只支持单个码，多个码核销需要多次调用
func (r *TaobaoVmarketEticketConsumeRequest) SetVerifyCode(_verifyCode string) error {
    r._verifyCode = _verifyCode
    r.Set("verify_code", _verifyCode)
    return nil
}

// VerifyCode Getter
func (r TaobaoVmarketEticketConsumeRequest) GetVerifyCode() string {
    return r._verifyCode
}
// ConsumeNum Setter
// 核销份数
func (r *TaobaoVmarketEticketConsumeRequest) SetConsumeNum(_consumeNum int64) error {
    r._consumeNum = _consumeNum
    r.Set("consume_num", _consumeNum)
    return nil
}

// ConsumeNum Getter
func (r TaobaoVmarketEticketConsumeRequest) GetConsumeNum() int64 {
    return r._consumeNum
}
// Token Setter
// 安全验证token,需要和发码通知中的token一致
func (r *TaobaoVmarketEticketConsumeRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r TaobaoVmarketEticketConsumeRequest) GetToken() string {
    return r._token
}
// CodemerchantId Setter
// 码商ID,是码商的话必须传递,如果是信任卖家不需要传
func (r *TaobaoVmarketEticketConsumeRequest) SetCodemerchantId(_codemerchantId int64) error {
    r._codemerchantId = _codemerchantId
    r.Set("codemerchant_id", _codemerchantId)
    return nil
}

// CodemerchantId Getter
func (r TaobaoVmarketEticketConsumeRequest) GetCodemerchantId() int64 {
    return r._codemerchantId
}
// Posid Setter
// 机具ID(此参数信任卖家可不传递，码商必须传递)
func (r *TaobaoVmarketEticketConsumeRequest) SetPosid(_posid string) error {
    r._posid = _posid
    r.Set("posid", _posid)
    return nil
}

// Posid Getter
func (r TaobaoVmarketEticketConsumeRequest) GetPosid() string {
    return r._posid
}
// Mobile Setter
// 手机后四位(没有特殊说明请不要传该参数)
func (r *TaobaoVmarketEticketConsumeRequest) SetMobile(_mobile string) error {
    r._mobile = _mobile
    r.Set("mobile", _mobile)
    return nil
}

// Mobile Getter
func (r TaobaoVmarketEticketConsumeRequest) GetMobile() string {
    return r._mobile
}
// NewCode Setter
// 核销后需要重新生成的码，如果不需要重新生成码，不要传该参数
func (r *TaobaoVmarketEticketConsumeRequest) SetNewCode(_newCode string) error {
    r._newCode = _newCode
    r.Set("new_code", _newCode)
    return nil
}

// NewCode Getter
func (r TaobaoVmarketEticketConsumeRequest) GetNewCode() string {
    return r._newCode
}
// SerialNum Setter
// 自定义核销流水号，如果核销调用失败，可以用该核销流水号进行冲正操作，需要小于等于100个字符(a-zA-Z0-9_)；每次核销都是唯一的流水号
func (r *TaobaoVmarketEticketConsumeRequest) SetSerialNum(_serialNum string) error {
    r._serialNum = _serialNum
    r.Set("serial_num", _serialNum)
    return nil
}

// SerialNum Getter
func (r TaobaoVmarketEticketConsumeRequest) GetSerialNum() string {
    return r._serialNum
}
// QrImages Setter
// 不需要上传二维码图片或者核销后不需重新生成码码商请不要传，需要传入二维码的码商请先调用taobao.vmarket.eticket.qrcode.upload接口，将返回的img_filename文件名称作为参数（如果二维码不变的话，也可将将发码时传入二维码文件名作为参数传入），文件名与参数new_code必须相互对应。
func (r *TaobaoVmarketEticketConsumeRequest) SetQrImages(_qrImages string) error {
    r._qrImages = _qrImages
    r.Set("qr_images", _qrImages)
    return nil
}

// QrImages Getter
func (r TaobaoVmarketEticketConsumeRequest) GetQrImages() string {
    return r._qrImages
}
