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
    _passEvaluation   bool
    // 鉴定不通过时的原因编码
    _failCode   int64
    // 鉴定结果图片列表
    _picUrlList   []string
    // 工单id
    _workcardId   int64
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
func (r *TmallServicecenterWorkcardEvaluateRequest) SetPassEvaluation(_passEvaluation bool) error {
    r._passEvaluation = _passEvaluation
    r.Set("pass_evaluation", _passEvaluation)
    return nil
}

// PassEvaluation Getter
func (r TmallServicecenterWorkcardEvaluateRequest) GetPassEvaluation() bool {
    return r._passEvaluation
}
// FailCode Setter
// 鉴定不通过时的原因编码
func (r *TmallServicecenterWorkcardEvaluateRequest) SetFailCode(_failCode int64) error {
    r._failCode = _failCode
    r.Set("fail_code", _failCode)
    return nil
}

// FailCode Getter
func (r TmallServicecenterWorkcardEvaluateRequest) GetFailCode() int64 {
    return r._failCode
}
// PicUrlList Setter
// 鉴定结果图片列表
func (r *TmallServicecenterWorkcardEvaluateRequest) SetPicUrlList(_picUrlList []string) error {
    r._picUrlList = _picUrlList
    r.Set("pic_url_list", _picUrlList)
    return nil
}

// PicUrlList Getter
func (r TmallServicecenterWorkcardEvaluateRequest) GetPicUrlList() []string {
    return r._picUrlList
}
// WorkcardId Setter
// 工单id
func (r *TmallServicecenterWorkcardEvaluateRequest) SetWorkcardId(_workcardId int64) error {
    r._workcardId = _workcardId
    r.Set("workcard_id", _workcardId)
    return nil
}

// WorkcardId Getter
func (r TmallServicecenterWorkcardEvaluateRequest) GetWorkcardId() int64 {
    return r._workcardId
}
