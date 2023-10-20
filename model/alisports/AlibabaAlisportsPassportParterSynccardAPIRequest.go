package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalisportspassportpartersynccardAPIRequest 阿里体育-卡信息同步接口 API请求
// alibaba.alisports.passport.parter.synccard
//
// 运享通修改卡号的时候，通知更新到阿里体育和支付宝卡包中
type AlibabaalisportspassportpartersynccardAPIRequest struct {
	model.Params
	// 用户的id
	_aliuid string
	// 类型：1.修改，2.删除
	_type string
	// 用户的老卡号(修改或删除之前的卡号)
	_oldCardNum string
	// 时间戳
	_alispTime string
	// appkey
	_alispAppKey string
	// 签名字符串
	_alispSign string
	// 改卡的中心id，如果卡号唯一则不需要传
	_centerNum string
}

// NewAlibabaalisportspassportpartersynccardRequest 初始化AlibabaalisportspassportpartersynccardAPIRequest对象
func NewAlibabaalisportspassportpartersynccardRequest() *AlibabaalisportspassportpartersynccardAPIRequest {
	return &AlibabaalisportspassportpartersynccardAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalisportspassportpartersynccardAPIRequest) GetApiMethodName() string {
	return "alibaba.alisports.passport.parter.synccard"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalisportspassportpartersynccardAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalisportspassportpartersynccardAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAliuid is Aliuid Setter
// 用户的id
func (r *AlibabaalisportspassportpartersynccardAPIRequest) SetAliuid(_aliuid string) error {
	r._aliuid = _aliuid
	r.Set("aliuid", _aliuid)
	return nil
}

// GetAliuid Aliuid Getter
func (r AlibabaalisportspassportpartersynccardAPIRequest) GetAliuid() string {
	return r._aliuid
}

// SetType is Type Setter
// 类型：1.修改，2.删除
func (r *AlibabaalisportspassportpartersynccardAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaalisportspassportpartersynccardAPIRequest) GetType() string {
	return r._type
}

// SetOldCardNum is OldCardNum Setter
// 用户的老卡号(修改或删除之前的卡号)
func (r *AlibabaalisportspassportpartersynccardAPIRequest) SetOldCardNum(_oldCardNum string) error {
	r._oldCardNum = _oldCardNum
	r.Set("old_card_num", _oldCardNum)
	return nil
}

// GetOldCardNum OldCardNum Getter
func (r AlibabaalisportspassportpartersynccardAPIRequest) GetOldCardNum() string {
	return r._oldCardNum
}

// SetAlispTime is AlispTime Setter
// 时间戳
func (r *AlibabaalisportspassportpartersynccardAPIRequest) SetAlispTime(_alispTime string) error {
	r._alispTime = _alispTime
	r.Set("alisp_time", _alispTime)
	return nil
}

// GetAlispTime AlispTime Getter
func (r AlibabaalisportspassportpartersynccardAPIRequest) GetAlispTime() string {
	return r._alispTime
}

// SetAlispAppKey is AlispAppKey Setter
// appkey
func (r *AlibabaalisportspassportpartersynccardAPIRequest) SetAlispAppKey(_alispAppKey string) error {
	r._alispAppKey = _alispAppKey
	r.Set("alisp_app_key", _alispAppKey)
	return nil
}

// GetAlispAppKey AlispAppKey Getter
func (r AlibabaalisportspassportpartersynccardAPIRequest) GetAlispAppKey() string {
	return r._alispAppKey
}

// SetAlispSign is AlispSign Setter
// 签名字符串
func (r *AlibabaalisportspassportpartersynccardAPIRequest) SetAlispSign(_alispSign string) error {
	r._alispSign = _alispSign
	r.Set("alisp_sign", _alispSign)
	return nil
}

// GetAlispSign AlispSign Getter
func (r AlibabaalisportspassportpartersynccardAPIRequest) GetAlispSign() string {
	return r._alispSign
}

// SetCenterNum is CenterNum Setter
// 改卡的中心id，如果卡号唯一则不需要传
func (r *AlibabaalisportspassportpartersynccardAPIRequest) SetCenterNum(_centerNum string) error {
	r._centerNum = _centerNum
	r.Set("center_num", _centerNum)
	return nil
}

// GetCenterNum CenterNum Getter
func (r AlibabaalisportspassportpartersynccardAPIRequest) GetCenterNum() string {
	return r._centerNum
}
