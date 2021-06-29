package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商反馈鉴定结果 API请求
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
    picUrlList   []string
    // 工单id
    workcardId   int64
}

// 初始化TmallServicecenterWorkcardEvaluateRequest对象
func NewTmallServicecenterWorkcardEvaluateRequest() *TmallServicecenterWorkcardEvaluateRequest{
    return &TmallServicecenterWorkcardEvaluateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardEvaluateRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.evaluate"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardEvaluateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PassEvaluation Setter
// 是否鉴定通过
func (r *TmallServicecenterWorkcardEvaluateRequest) SetPassEvaluation(passEvaluation bool) error {
    r.passEvaluation = passEvaluation
    r.Set("pass_evaluation", passEvaluation)
    return nil
}

// PassEvaluation Getter
func (r TmallServicecenterWorkcardEvaluateRequest) GetPassEvaluation() bool {
    return r.passEvaluation
}
// FailCode Setter
// 鉴定不通过时的原因编码
func (r *TmallServicecenterWorkcardEvaluateRequest) SetFailCode(failCode int64) error {
    r.failCode = failCode
    r.Set("fail_code", failCode)
    return nil
}

// FailCode Getter
func (r TmallServicecenterWorkcardEvaluateRequest) GetFailCode() int64 {
    return r.failCode
}
// PicUrlList Setter
// 鉴定结果图片列表
func (r *TmallServicecenterWorkcardEvaluateRequest) SetPicUrlList(picUrlList []string) error {
    r.picUrlList = picUrlList
    r.Set("pic_url_list", picUrlList)
    return nil
}

// PicUrlList Getter
func (r TmallServicecenterWorkcardEvaluateRequest) GetPicUrlList() []string {
    return r.picUrlList
}
// WorkcardId Setter
// 工单id
func (r *TmallServicecenterWorkcardEvaluateRequest) SetWorkcardId(workcardId int64) error {
    r.workcardId = workcardId
    r.Set("workcard_id", workcardId)
    return nil
}

// WorkcardId Getter
func (r TmallServicecenterWorkcardEvaluateRequest) GetWorkcardId() int64 {
    return r.workcardId
}
