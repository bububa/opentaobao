package shenjing

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
访客通过PAD提交访客码 API请求
alibaba.ib.shenjing.visitor.pad.fetchcodeverify

访客通过PAD提交访客码，录脸进入园区。
*/
type AlibabaIbShenjingVisitorPadFetchcodeverifyRequest struct {
    model.Params
    // 访客码
    _visitorCode   int64
    // 终端ID
    _termId   string
}

// 初始化AlibabaIbShenjingVisitorPadFetchcodeverifyRequest对象
func NewAlibabaIbShenjingVisitorPadFetchcodeverifyRequest() *AlibabaIbShenjingVisitorPadFetchcodeverifyRequest{
    return &AlibabaIbShenjingVisitorPadFetchcodeverifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIbShenjingVisitorPadFetchcodeverifyRequest) GetApiMethodName() string {
    return "alibaba.ib.shenjing.visitor.pad.fetchcodeverify"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIbShenjingVisitorPadFetchcodeverifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// VisitorCode Setter
// 访客码
func (r *AlibabaIbShenjingVisitorPadFetchcodeverifyRequest) SetVisitorCode(_visitorCode int64) error {
    r._visitorCode = _visitorCode
    r.Set("visitor_code", _visitorCode)
    return nil
}

// VisitorCode Getter
func (r AlibabaIbShenjingVisitorPadFetchcodeverifyRequest) GetVisitorCode() int64 {
    return r._visitorCode
}
// TermId Setter
// 终端ID
func (r *AlibabaIbShenjingVisitorPadFetchcodeverifyRequest) SetTermId(_termId string) error {
    r._termId = _termId
    r.Set("term_id", _termId)
    return nil
}

// TermId Getter
func (r AlibabaIbShenjingVisitorPadFetchcodeverifyRequest) GetTermId() string {
    return r._termId
}
