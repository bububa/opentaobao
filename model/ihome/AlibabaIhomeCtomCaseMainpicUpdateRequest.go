package ihome

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
方案渲染图修改 API请求
alibaba.ihome.ctom.case.mainpic.update

用于在门店工作台里的编辑器保存方案，由三维家后端调用阿里后端，保存方案信息
此接口只允许ihome业务使用，用于门店的编辑功能，只允许广东三维家信息科技有限公司一家公司调用，不适用于其他业务。
*/
type AlibabaIhomeCtomCaseMainpicUpdateRequest struct {
    model.Params
    // 32位字符串
    _traceId   string
    // 方案id
    _caseId   string
    // 图片的地址
    _picUrl   string
    // 图片类型
    _picType   string
}

// 初始化AlibabaIhomeCtomCaseMainpicUpdateRequest对象
func NewAlibabaIhomeCtomCaseMainpicUpdateRequest() *AlibabaIhomeCtomCaseMainpicUpdateRequest{
    return &AlibabaIhomeCtomCaseMainpicUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIhomeCtomCaseMainpicUpdateRequest) GetApiMethodName() string {
    return "alibaba.ihome.ctom.case.mainpic.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIhomeCtomCaseMainpicUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TraceId Setter
// 32位字符串
func (r *AlibabaIhomeCtomCaseMainpicUpdateRequest) SetTraceId(_traceId string) error {
    r._traceId = _traceId
    r.Set("trace_id", _traceId)
    return nil
}

// TraceId Getter
func (r AlibabaIhomeCtomCaseMainpicUpdateRequest) GetTraceId() string {
    return r._traceId
}
// CaseId Setter
// 方案id
func (r *AlibabaIhomeCtomCaseMainpicUpdateRequest) SetCaseId(_caseId string) error {
    r._caseId = _caseId
    r.Set("case_id", _caseId)
    return nil
}

// CaseId Getter
func (r AlibabaIhomeCtomCaseMainpicUpdateRequest) GetCaseId() string {
    return r._caseId
}
// PicUrl Setter
// 图片的地址
func (r *AlibabaIhomeCtomCaseMainpicUpdateRequest) SetPicUrl(_picUrl string) error {
    r._picUrl = _picUrl
    r.Set("pic_url", _picUrl)
    return nil
}

// PicUrl Getter
func (r AlibabaIhomeCtomCaseMainpicUpdateRequest) GetPicUrl() string {
    return r._picUrl
}
// PicType Setter
// 图片类型
func (r *AlibabaIhomeCtomCaseMainpicUpdateRequest) SetPicType(_picType string) error {
    r._picType = _picType
    r.Set("pic_type", _picType)
    return nil
}

// PicType Getter
func (r AlibabaIhomeCtomCaseMainpicUpdateRequest) GetPicType() string {
    return r._picType
}