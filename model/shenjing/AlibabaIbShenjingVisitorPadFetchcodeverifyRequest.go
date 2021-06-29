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
    visitorCode   int64
    // 终端ID
    termId   string
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
func (r *AlibabaIbShenjingVisitorPadFetchcodeverifyRequest) SetVisitorCode(visitorCode int64) error {
    r.visitorCode = visitorCode
    r.Set("visitor_code", visitorCode)
    return nil
}

// VisitorCode Getter
func (r AlibabaIbShenjingVisitorPadFetchcodeverifyRequest) GetVisitorCode() int64 {
    return r.visitorCode
}
// TermId Setter
// 终端ID
func (r *AlibabaIbShenjingVisitorPadFetchcodeverifyRequest) SetTermId(termId string) error {
    r.termId = termId
    r.Set("term_id", termId)
    return nil
}

// TermId Getter
func (r AlibabaIbShenjingVisitorPadFetchcodeverifyRequest) GetTermId() string {
    return r.termId
}
