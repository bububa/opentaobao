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
type AlibabaIbShenjingVisitorPadFetchcodeverifyAPIRequest struct {
    model.Params
    // 访客码
    _visitorCode   int64
    // 终端ID
    _termId   string
}

// 初始化AlibabaIbShenjingVisitorPadFetchcodeverifyAPIRequest对象
func NewAlibabaIbShenjingVisitorPadFetchcodeverifyRequest() *AlibabaIbShenjingVisitorPadFetchcodeverifyAPIRequest{
    return &AlibabaIbShenjingVisitorPadFetchcodeverifyAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIbShenjingVisitorPadFetchcodeverifyAPIRequest) GetApiMethodName() string {
    return "alibaba.ib.shenjing.visitor.pad.fetchcodeverify"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIbShenjingVisitorPadFetchcodeverifyAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// VisitorCode Setter
// 访客码
func (r *AlibabaIbShenjingVisitorPadFetchcodeverifyAPIRequest) SetVisitorCode(_visitorCode int64) error {
    r._visitorCode = _visitorCode
    r.Set("visitor_code", _visitorCode)
    return nil
}

// VisitorCode Getter
func (r AlibabaIbShenjingVisitorPadFetchcodeverifyAPIRequest) GetVisitorCode() int64 {
    return r._visitorCode
}
// TermId Setter
// 终端ID
func (r *AlibabaIbShenjingVisitorPadFetchcodeverifyAPIRequest) SetTermId(_termId string) error {
    r._termId = _termId
    r.Set("term_id", _termId)
    return nil
}

// TermId Getter
func (r AlibabaIbShenjingVisitorPadFetchcodeverifyAPIRequest) GetTermId() string {
    return r._termId
}
