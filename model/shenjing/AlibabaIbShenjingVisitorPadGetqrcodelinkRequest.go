package shenjing

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
pad获取二维码 API请求
alibaba.ib.shenjing.visitor.pad.getqrcodelink

pad获取二维码链接。扫码录入人脸。
*/
type AlibabaIbShenjingVisitorPadGetqrcodelinkRequest struct {
    model.Params
    // 终端id
    _termId   string
}

// 初始化AlibabaIbShenjingVisitorPadGetqrcodelinkRequest对象
func NewAlibabaIbShenjingVisitorPadGetqrcodelinkRequest() *AlibabaIbShenjingVisitorPadGetqrcodelinkRequest{
    return &AlibabaIbShenjingVisitorPadGetqrcodelinkRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIbShenjingVisitorPadGetqrcodelinkRequest) GetApiMethodName() string {
    return "alibaba.ib.shenjing.visitor.pad.getqrcodelink"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIbShenjingVisitorPadGetqrcodelinkRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TermId Setter
// 终端id
func (r *AlibabaIbShenjingVisitorPadGetqrcodelinkRequest) SetTermId(_termId string) error {
    r._termId = _termId
    r.Set("term_id", _termId)
    return nil
}

// TermId Getter
func (r AlibabaIbShenjingVisitorPadGetqrcodelinkRequest) GetTermId() string {
    return r._termId
}
