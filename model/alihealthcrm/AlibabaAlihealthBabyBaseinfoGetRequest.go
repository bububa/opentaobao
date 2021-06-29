package alihealthcrm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
三方从我们这获取宝宝信息 API请求
alibaba.alihealth.baby.baseinfo.get

三方从我们这获取宝宝信息
*/
type AlibabaAlihealthBabyBaseinfoGetRequest struct {
    model.Params
    // 宝宝id
    babyId   int64
    // 宝宝所属的用户
    tpUserId   int64
}

// 初始化AlibabaAlihealthBabyBaseinfoGetRequest对象
func NewAlibabaAlihealthBabyBaseinfoGetRequest() *AlibabaAlihealthBabyBaseinfoGetRequest{
    return &AlibabaAlihealthBabyBaseinfoGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthBabyBaseinfoGetRequest) GetApiMethodName() string {
    return "alibaba.alihealth.baby.baseinfo.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthBabyBaseinfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BabyId Setter
// 宝宝id
func (r *AlibabaAlihealthBabyBaseinfoGetRequest) SetBabyId(babyId int64) error {
    r.babyId = babyId
    r.Set("baby_id", babyId)
    return nil
}

// BabyId Getter
func (r AlibabaAlihealthBabyBaseinfoGetRequest) GetBabyId() int64 {
    return r.babyId
}
// TpUserId Setter
// 宝宝所属的用户
func (r *AlibabaAlihealthBabyBaseinfoGetRequest) SetTpUserId(tpUserId int64) error {
    r.tpUserId = tpUserId
    r.Set("tp_user_id", tpUserId)
    return nil
}

// TpUserId Getter
func (r AlibabaAlihealthBabyBaseinfoGetRequest) GetTpUserId() int64 {
    return r.tpUserId
}
