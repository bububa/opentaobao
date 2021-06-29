package shenjing

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
pad获取二维码 APIRequest
alibaba.ib.shenjing.visitor.pad.getqrcodelink

pad获取二维码链接。扫码录入人脸。
*/
type AlibabaIbShenjingVisitorPadGetqrcodelinkRequest struct {
    model.Params

    // 终端id
    termId   string 

}

func NewAlibabaIbShenjingVisitorPadGetqrcodelinkRequest() *AlibabaIbShenjingVisitorPadGetqrcodelinkRequest{
    return &AlibabaIbShenjingVisitorPadGetqrcodelinkRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIbShenjingVisitorPadGetqrcodelinkRequest) GetApiMethodName() string {
    return "alibaba.ib.shenjing.visitor.pad.getqrcodelink"
}

func (r AlibabaIbShenjingVisitorPadGetqrcodelinkRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIbShenjingVisitorPadGetqrcodelinkRequest) SetTermId(termId string) error {
    r.termId = termId
    r.Set("term_id", termId)
    return nil
}

func (r AlibabaIbShenjingVisitorPadGetqrcodelinkRequest) GetTermId() string {
    return r.termId
}

