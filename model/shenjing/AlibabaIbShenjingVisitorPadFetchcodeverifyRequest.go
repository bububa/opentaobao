package shenjing

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
访客通过PAD提交访客码 APIRequest
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

func NewAlibabaIbShenjingVisitorPadFetchcodeverifyRequest() *AlibabaIbShenjingVisitorPadFetchcodeverifyRequest{
    return &AlibabaIbShenjingVisitorPadFetchcodeverifyRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIbShenjingVisitorPadFetchcodeverifyRequest) GetApiMethodName() string {
    return "alibaba.ib.shenjing.visitor.pad.fetchcodeverify"
}

func (r AlibabaIbShenjingVisitorPadFetchcodeverifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIbShenjingVisitorPadFetchcodeverifyRequest) SetVisitorCode(visitorCode int64) error {
    r.visitorCode = visitorCode
    r.Set("visitor_code", visitorCode)
    return nil
}

func (r AlibabaIbShenjingVisitorPadFetchcodeverifyRequest) GetVisitorCode() int64 {
    return r.visitorCode
}

func (r *AlibabaIbShenjingVisitorPadFetchcodeverifyRequest) SetTermId(termId string) error {
    r.termId = termId
    r.Set("term_id", termId)
    return nil
}

func (r AlibabaIbShenjingVisitorPadFetchcodeverifyRequest) GetTermId() string {
    return r.termId
}

