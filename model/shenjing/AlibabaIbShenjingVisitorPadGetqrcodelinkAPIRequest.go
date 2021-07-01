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
type AlibabaIbShenjingVisitorPadGetqrcodelinkAPIRequest struct {
    model.Params
    // 终端id
    _termId   string
}

// 初始化AlibabaIbShenjingVisitorPadGetqrcodelinkAPIRequest对象
func NewAlibabaIbShenjingVisitorPadGetqrcodelinkRequest() *AlibabaIbShenjingVisitorPadGetqrcodelinkAPIRequest{
    return &AlibabaIbShenjingVisitorPadGetqrcodelinkAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIbShenjingVisitorPadGetqrcodelinkAPIRequest) GetApiMethodName() string {
    return "alibaba.ib.shenjing.visitor.pad.getqrcodelink"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIbShenjingVisitorPadGetqrcodelinkAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TermId Setter
// 终端id
func (r *AlibabaIbShenjingVisitorPadGetqrcodelinkAPIRequest) SetTermId(_termId string) error {
    r._termId = _termId
    r.Set("term_id", _termId)
    return nil
}

// TermId Getter
func (r AlibabaIbShenjingVisitorPadGetqrcodelinkAPIRequest) GetTermId() string {
    return r._termId
}
