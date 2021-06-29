package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取单个审批单 APIRequest
alitrip.btrip.apply.get

获取单个审批单的详情数据
*/
type AlitripBtripApplyGetRequest struct {
    model.Params

    // 外部审批单id
    thirdpartApplyId   string 

    // 阿里商旅审批单id
    applyId   int64 

    // 企业id
    corpId   string 

    // 审批单展示id
    applyShowId   string 

}

func NewAlitripBtripApplyGetRequest() *AlitripBtripApplyGetRequest{
    return &AlitripBtripApplyGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripApplyGetRequest) GetApiMethodName() string {
    return "alitrip.btrip.apply.get"
}

func (r AlitripBtripApplyGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripApplyGetRequest) SetThirdpartApplyId(thirdpartApplyId string) error {
    r.thirdpartApplyId = thirdpartApplyId
    r.Set("thirdpart_apply_id", thirdpartApplyId)
    return nil
}

func (r AlitripBtripApplyGetRequest) GetThirdpartApplyId() string {
    return r.thirdpartApplyId
}

func (r *AlitripBtripApplyGetRequest) SetApplyId(applyId int64) error {
    r.applyId = applyId
    r.Set("apply_id", applyId)
    return nil
}

func (r AlitripBtripApplyGetRequest) GetApplyId() int64 {
    return r.applyId
}

func (r *AlitripBtripApplyGetRequest) SetCorpId(corpId string) error {
    r.corpId = corpId
    r.Set("corp_id", corpId)
    return nil
}

func (r AlitripBtripApplyGetRequest) GetCorpId() string {
    return r.corpId
}

func (r *AlitripBtripApplyGetRequest) SetApplyShowId(applyShowId string) error {
    r.applyShowId = applyShowId
    r.Set("apply_show_id", applyShowId)
    return nil
}

func (r AlitripBtripApplyGetRequest) GetApplyShowId() string {
    return r.applyShowId
}

