package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户费用归属 API请求
alitrip.btrip.cost.center.get

获取差旅申请用户的费用归属列表
*/
type AlitripBtripCostCenterGetRequest struct {
    model.Params
    // 企业id
    corpId   string
    // 用户id
    userId   string
}

// 初始化AlitripBtripCostCenterGetRequest对象
func NewAlitripBtripCostCenterGetRequest() *AlitripBtripCostCenterGetRequest{
    return &AlitripBtripCostCenterGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripCostCenterGetRequest) GetApiMethodName() string {
    return "alitrip.btrip.cost.center.get"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripCostCenterGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CorpId Setter
// 企业id
func (r *AlitripBtripCostCenterGetRequest) SetCorpId(corpId string) error {
    r.corpId = corpId
    r.Set("corp_id", corpId)
    return nil
}

// CorpId Getter
func (r AlitripBtripCostCenterGetRequest) GetCorpId() string {
    return r.corpId
}
// UserId Setter
// 用户id
func (r *AlitripBtripCostCenterGetRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r AlitripBtripCostCenterGetRequest) GetUserId() string {
    return r.userId
}
