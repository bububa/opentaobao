package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取单个审批单 API请求
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

// 初始化AlitripBtripApplyGetRequest对象
func NewAlitripBtripApplyGetRequest() *AlitripBtripApplyGetRequest{
    return &AlitripBtripApplyGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripApplyGetRequest) GetApiMethodName() string {
    return "alitrip.btrip.apply.get"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripApplyGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ThirdpartApplyId Setter
// 外部审批单id
func (r *AlitripBtripApplyGetRequest) SetThirdpartApplyId(thirdpartApplyId string) error {
    r.thirdpartApplyId = thirdpartApplyId
    r.Set("thirdpart_apply_id", thirdpartApplyId)
    return nil
}

// ThirdpartApplyId Getter
func (r AlitripBtripApplyGetRequest) GetThirdpartApplyId() string {
    return r.thirdpartApplyId
}
// ApplyId Setter
// 阿里商旅审批单id
func (r *AlitripBtripApplyGetRequest) SetApplyId(applyId int64) error {
    r.applyId = applyId
    r.Set("apply_id", applyId)
    return nil
}

// ApplyId Getter
func (r AlitripBtripApplyGetRequest) GetApplyId() int64 {
    return r.applyId
}
// CorpId Setter
// 企业id
func (r *AlitripBtripApplyGetRequest) SetCorpId(corpId string) error {
    r.corpId = corpId
    r.Set("corp_id", corpId)
    return nil
}

// CorpId Getter
func (r AlitripBtripApplyGetRequest) GetCorpId() string {
    return r.corpId
}
// ApplyShowId Setter
// 审批单展示id
func (r *AlitripBtripApplyGetRequest) SetApplyShowId(applyShowId string) error {
    r.applyShowId = applyShowId
    r.Set("apply_show_id", applyShowId)
    return nil
}

// ApplyShowId Getter
func (r AlitripBtripApplyGetRequest) GetApplyShowId() string {
    return r.applyShowId
}
