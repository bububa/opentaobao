package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商反馈鉴定结果 APIRequest
tmall.servicecenter.workcard.evaluate

服务商反馈鉴定结果
*/
type TmallServicecenterWorkcardEvaluateRequest struct {
    model.Params

    // 是否鉴定通过
    passEvaluation   bool 

    // 鉴定不通过时的原因编码
    failCode   int64 

    // 鉴定结果图片列表
    picUrlList   []String 

    // 工单id
    workcardId   int64 

}

func NewTmallServicecenterWorkcardEvaluateRequest() *TmallServicecenterWorkcardEvaluateRequest{
    return &TmallServicecenterWorkcardEvaluateRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterWorkcardEvaluateRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.evaluate"
}

func (r TmallServicecenterWorkcardEvaluateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterWorkcardEvaluateRequest) SetPassEvaluation(passEvaluation bool) error {
    r.passEvaluation = passEvaluation
    r.Set("pass_evaluation", passEvaluation)
    return nil
}

func (r TmallServicecenterWorkcardEvaluateRequest) GetPassEvaluation() bool {
    return r.passEvaluation
}

func (r *TmallServicecenterWorkcardEvaluateRequest) SetFailCode(failCode int64) error {
    r.failCode = failCode
    r.Set("fail_code", failCode)
    return nil
}

func (r TmallServicecenterWorkcardEvaluateRequest) GetFailCode() int64 {
    return r.failCode
}

func (r *TmallServicecenterWorkcardEvaluateRequest) SetPicUrlList(picUrlList []String) error {
    r.picUrlList = picUrlList
    r.Set("pic_url_list", picUrlList)
    return nil
}

func (r TmallServicecenterWorkcardEvaluateRequest) GetPicUrlList() []String {
    return r.picUrlList
}

func (r *TmallServicecenterWorkcardEvaluateRequest) SetWorkcardId(workcardId int64) error {
    r.workcardId = workcardId
    r.Set("workcard_id", workcardId)
    return nil
}

func (r TmallServicecenterWorkcardEvaluateRequest) GetWorkcardId() int64 {
    return r.workcardId
}

