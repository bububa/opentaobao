package alisports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里体育-卡信息同步接口 APIRequest
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

func NewAlibabaAlisportsPassportParterSynccardRequest() *AlibabaAlisportsPassportParterSynccardRequest{
    return &AlibabaAlisportsPassportParterSynccardRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlisportsPassportParterSynccardRequest) GetApiMethodName() string {
    return "alibaba.alisports.passport.parter.synccard"
}

func (r AlibabaAlisportsPassportParterSynccardRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlisportsPassportParterSynccardRequest) SetAliuid(aliuid string) error {
    r.aliuid = aliuid
    r.Set("aliuid", aliuid)
    return nil
}

func (r AlibabaAlisportsPassportParterSynccardRequest) GetAliuid() string {
    return r.aliuid
}

func (r *AlibabaAlisportsPassportParterSynccardRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r AlibabaAlisportsPassportParterSynccardRequest) GetType() string {
    return r.type
}

func (r *AlibabaAlisportsPassportParterSynccardRequest) SetOldCardNum(oldCardNum string) error {
    r.oldCardNum = oldCardNum
    r.Set("old_card_num", oldCardNum)
    return nil
}

func (r AlibabaAlisportsPassportParterSynccardRequest) GetOldCardNum() string {
    return r.oldCardNum
}

func (r *AlibabaAlisportsPassportParterSynccardRequest) SetAlispTime(alispTime string) error {
    r.alispTime = alispTime
    r.Set("alisp_time", alispTime)
    return nil
}

func (r AlibabaAlisportsPassportParterSynccardRequest) GetAlispTime() string {
    return r.alispTime
}

func (r *AlibabaAlisportsPassportParterSynccardRequest) SetAlispAppKey(alispAppKey string) error {
    r.alispAppKey = alispAppKey
    r.Set("alisp_app_key", alispAppKey)
    return nil
}

func (r AlibabaAlisportsPassportParterSynccardRequest) GetAlispAppKey() string {
    return r.alispAppKey
}

func (r *AlibabaAlisportsPassportParterSynccardRequest) SetAlispSign(alispSign string) error {
    r.alispSign = alispSign
    r.Set("alisp_sign", alispSign)
    return nil
}

func (r AlibabaAlisportsPassportParterSynccardRequest) GetAlispSign() string {
    return r.alispSign
}

func (r *AlibabaAlisportsPassportParterSynccardRequest) SetCenterNum(centerNum string) error {
    r.centerNum = centerNum
    r.Set("center_num", centerNum)
    return nil
}

func (r AlibabaAlisportsPassportParterSynccardRequest) GetCenterNum() string {
    return r.centerNum
}

