package alisports

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlisportsPassportParterSynccardAPIRequest 阿里体育-卡信息同步接口 API请求
// alibaba.alisports.passport.parter.synccard
//
// 运享通修改卡号的时候，通知更新到阿里体育和支付宝卡包中
type AlibabaAlisportsPassportParterSynccardAPIRequest struct {
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

// NewAlibabaAlisportsPassportParterSynccardRequest 初始化AlibabaAlisportsPassportParterSynccardAPIRequest对象
func NewAlibabaAlisportsPassportParterSynccardRequest() *AlibabaAlisportsPassportParterSynccardAPIRequest {
	return &AlibabaAlisportsPassportParterSynccardAPIRequest{
		Params: model.NewParams(7),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlisportsPassportParterSynccardAPIRequest) Reset() {
	r._aliuid = ""
	r._type = ""
	r._oldCardNum = ""
	r._alispTime = ""
	r._alispAppKey = ""
	r._alispSign = ""
	r._centerNum = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsPassportParterSynccardAPIRequest) GetApiMethodName() string {
	return "alibaba.alisports.passport.parter.synccard"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsPassportParterSynccardAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlisportsPassportParterSynccardAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAliuid is Aliuid Setter
// 用户的id
func (r *AlibabaAlisportsPassportParterSynccardAPIRequest) SetAliuid(_aliuid string) error {
	r._aliuid = _aliuid
	r.Set("aliuid", _aliuid)
	return nil
}

// GetAliuid Aliuid Getter
func (r AlibabaAlisportsPassportParterSynccardAPIRequest) GetAliuid() string {
	return r._aliuid
}

// SetType is Type Setter
// 类型：1.修改，2.删除
func (r *AlibabaAlisportsPassportParterSynccardAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaAlisportsPassportParterSynccardAPIRequest) GetType() string {
	return r._type
}

// SetOldCardNum is OldCardNum Setter
// 用户的老卡号(修改或删除之前的卡号)
func (r *AlibabaAlisportsPassportParterSynccardAPIRequest) SetOldCardNum(_oldCardNum string) error {
	r._oldCardNum = _oldCardNum
	r.Set("old_card_num", _oldCardNum)
	return nil
}

// GetOldCardNum OldCardNum Getter
func (r AlibabaAlisportsPassportParterSynccardAPIRequest) GetOldCardNum() string {
	return r._oldCardNum
}

// SetAlispTime is AlispTime Setter
// 时间戳
func (r *AlibabaAlisportsPassportParterSynccardAPIRequest) SetAlispTime(_alispTime string) error {
	r._alispTime = _alispTime
	r.Set("alisp_time", _alispTime)
	return nil
}

// GetAlispTime AlispTime Getter
func (r AlibabaAlisportsPassportParterSynccardAPIRequest) GetAlispTime() string {
	return r._alispTime
}

// SetAlispAppKey is AlispAppKey Setter
// appkey
func (r *AlibabaAlisportsPassportParterSynccardAPIRequest) SetAlispAppKey(_alispAppKey string) error {
	r._alispAppKey = _alispAppKey
	r.Set("alisp_app_key", _alispAppKey)
	return nil
}

// GetAlispAppKey AlispAppKey Getter
func (r AlibabaAlisportsPassportParterSynccardAPIRequest) GetAlispAppKey() string {
	return r._alispAppKey
}

// SetAlispSign is AlispSign Setter
// 签名字符串
func (r *AlibabaAlisportsPassportParterSynccardAPIRequest) SetAlispSign(_alispSign string) error {
	r._alispSign = _alispSign
	r.Set("alisp_sign", _alispSign)
	return nil
}

// GetAlispSign AlispSign Getter
func (r AlibabaAlisportsPassportParterSynccardAPIRequest) GetAlispSign() string {
	return r._alispSign
}

// SetCenterNum is CenterNum Setter
// 改卡的中心id，如果卡号唯一则不需要传
func (r *AlibabaAlisportsPassportParterSynccardAPIRequest) SetCenterNum(_centerNum string) error {
	r._centerNum = _centerNum
	r.Set("center_num", _centerNum)
	return nil
}

// GetCenterNum CenterNum Getter
func (r AlibabaAlisportsPassportParterSynccardAPIRequest) GetCenterNum() string {
	return r._centerNum
}

var poolAlibabaAlisportsPassportParterSynccardAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlisportsPassportParterSynccardRequest()
	},
}

// GetAlibabaAlisportsPassportParterSynccardRequest 从 sync.Pool 获取 AlibabaAlisportsPassportParterSynccardAPIRequest
func GetAlibabaAlisportsPassportParterSynccardAPIRequest() *AlibabaAlisportsPassportParterSynccardAPIRequest {
	return poolAlibabaAlisportsPassportParterSynccardAPIRequest.Get().(*AlibabaAlisportsPassportParterSynccardAPIRequest)
}

// ReleaseAlibabaAlisportsPassportParterSynccardAPIRequest 将 AlibabaAlisportsPassportParterSynccardAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlisportsPassportParterSynccardAPIRequest(v *AlibabaAlisportsPassportParterSynccardAPIRequest) {
	v.Reset()
	poolAlibabaAlisportsPassportParterSynccardAPIRequest.Put(v)
}
