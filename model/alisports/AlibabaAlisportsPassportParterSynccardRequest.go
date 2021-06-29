package alisports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里体育-卡信息同步接口 API请求
alibaba.alisports.passport.parter.synccard

运享通修改卡号的时候，通知更新到阿里体育和支付宝卡包中
*/
type AlibabaAlisportsPassportParterSynccardRequest struct {
    model.Params
    // 用户的id
    _aliuid   string
    // 类型：1.修改，2.删除
    _type   string
    // 用户的老卡号(修改或删除之前的卡号)
    _oldCardNum   string
    // 时间戳
    _alispTime   string
    // appkey
    _alispAppKey   string
    // 签名字符串
    _alispSign   string
    // 改卡的中心id，如果卡号唯一则不需要传
    _centerNum   string
}

// 初始化AlibabaAlisportsPassportParterSynccardRequest对象
func NewAlibabaAlisportsPassportParterSynccardRequest() *AlibabaAlisportsPassportParterSynccardRequest{
    return &AlibabaAlisportsPassportParterSynccardRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsPassportParterSynccardRequest) GetApiMethodName() string {
    return "alibaba.alisports.passport.parter.synccard"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsPassportParterSynccardRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Aliuid Setter
// 用户的id
func (r *AlibabaAlisportsPassportParterSynccardRequest) SetAliuid(_aliuid string) error {
    r._aliuid = _aliuid
    r.Set("aliuid", _aliuid)
    return nil
}

// Aliuid Getter
func (r AlibabaAlisportsPassportParterSynccardRequest) GetAliuid() string {
    return r._aliuid
}
// Type Setter
// 类型：1.修改，2.删除
func (r *AlibabaAlisportsPassportParterSynccardRequest) SetType(_type string) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r AlibabaAlisportsPassportParterSynccardRequest) GetType() string {
    return r._type
}
// OldCardNum Setter
// 用户的老卡号(修改或删除之前的卡号)
func (r *AlibabaAlisportsPassportParterSynccardRequest) SetOldCardNum(_oldCardNum string) error {
    r._oldCardNum = _oldCardNum
    r.Set("old_card_num", _oldCardNum)
    return nil
}

// OldCardNum Getter
func (r AlibabaAlisportsPassportParterSynccardRequest) GetOldCardNum() string {
    return r._oldCardNum
}
// AlispTime Setter
// 时间戳
func (r *AlibabaAlisportsPassportParterSynccardRequest) SetAlispTime(_alispTime string) error {
    r._alispTime = _alispTime
    r.Set("alisp_time", _alispTime)
    return nil
}

// AlispTime Getter
func (r AlibabaAlisportsPassportParterSynccardRequest) GetAlispTime() string {
    return r._alispTime
}
// AlispAppKey Setter
// appkey
func (r *AlibabaAlisportsPassportParterSynccardRequest) SetAlispAppKey(_alispAppKey string) error {
    r._alispAppKey = _alispAppKey
    r.Set("alisp_app_key", _alispAppKey)
    return nil
}

// AlispAppKey Getter
func (r AlibabaAlisportsPassportParterSynccardRequest) GetAlispAppKey() string {
    return r._alispAppKey
}
// AlispSign Setter
// 签名字符串
func (r *AlibabaAlisportsPassportParterSynccardRequest) SetAlispSign(_alispSign string) error {
    r._alispSign = _alispSign
    r.Set("alisp_sign", _alispSign)
    return nil
}

// AlispSign Getter
func (r AlibabaAlisportsPassportParterSynccardRequest) GetAlispSign() string {
    return r._alispSign
}
// CenterNum Setter
// 改卡的中心id，如果卡号唯一则不需要传
func (r *AlibabaAlisportsPassportParterSynccardRequest) SetCenterNum(_centerNum string) error {
    r._centerNum = _centerNum
    r.Set("center_num", _centerNum)
    return nil
}

// CenterNum Getter
func (r AlibabaAlisportsPassportParterSynccardRequest) GetCenterNum() string {
    return r._centerNum
}
