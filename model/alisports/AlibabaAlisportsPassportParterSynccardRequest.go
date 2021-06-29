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
    aliuid   string
    // 类型：1.修改，2.删除
    type   string
    // 用户的老卡号(修改或删除之前的卡号)
    oldCardNum   string
    // 时间戳
    alispTime   string
    // appkey
    alispAppKey   string
    // 签名字符串
    alispSign   string
    // 改卡的中心id，如果卡号唯一则不需要传
    centerNum   string
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
func (r *AlibabaAlisportsPassportParterSynccardRequest) SetAliuid(aliuid string) error {
    r.aliuid = aliuid
    r.Set("aliuid", aliuid)
    return nil
}

// Aliuid Getter
func (r AlibabaAlisportsPassportParterSynccardRequest) GetAliuid() string {
    return r.aliuid
}
// Type Setter
// 类型：1.修改，2.删除
func (r *AlibabaAlisportsPassportParterSynccardRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r AlibabaAlisportsPassportParterSynccardRequest) GetType() string {
    return r.type
}
// OldCardNum Setter
// 用户的老卡号(修改或删除之前的卡号)
func (r *AlibabaAlisportsPassportParterSynccardRequest) SetOldCardNum(oldCardNum string) error {
    r.oldCardNum = oldCardNum
    r.Set("old_card_num", oldCardNum)
    return nil
}

// OldCardNum Getter
func (r AlibabaAlisportsPassportParterSynccardRequest) GetOldCardNum() string {
    return r.oldCardNum
}
// AlispTime Setter
// 时间戳
func (r *AlibabaAlisportsPassportParterSynccardRequest) SetAlispTime(alispTime string) error {
    r.alispTime = alispTime
    r.Set("alisp_time", alispTime)
    return nil
}

// AlispTime Getter
func (r AlibabaAlisportsPassportParterSynccardRequest) GetAlispTime() string {
    return r.alispTime
}
// AlispAppKey Setter
// appkey
func (r *AlibabaAlisportsPassportParterSynccardRequest) SetAlispAppKey(alispAppKey string) error {
    r.alispAppKey = alispAppKey
    r.Set("alisp_app_key", alispAppKey)
    return nil
}

// AlispAppKey Getter
func (r AlibabaAlisportsPassportParterSynccardRequest) GetAlispAppKey() string {
    return r.alispAppKey
}
// AlispSign Setter
// 签名字符串
func (r *AlibabaAlisportsPassportParterSynccardRequest) SetAlispSign(alispSign string) error {
    r.alispSign = alispSign
    r.Set("alisp_sign", alispSign)
    return nil
}

// AlispSign Getter
func (r AlibabaAlisportsPassportParterSynccardRequest) GetAlispSign() string {
    return r.alispSign
}
// CenterNum Setter
// 改卡的中心id，如果卡号唯一则不需要传
func (r *AlibabaAlisportsPassportParterSynccardRequest) SetCenterNum(centerNum string) error {
    r.centerNum = centerNum
    r.Set("center_num", centerNum)
    return nil
}

// CenterNum Getter
func (r AlibabaAlisportsPassportParterSynccardRequest) GetCenterNum() string {
    return r.centerNum
}
