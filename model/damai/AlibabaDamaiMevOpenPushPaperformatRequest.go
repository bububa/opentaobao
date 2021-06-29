package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-票纸版式接口pushPaperFormat API请求
alibaba.damai.mev.open.push.paperformat

pushPaperFormat
*/
type AlibabaDamaiMevOpenPushPaperformatRequest struct {
    model.Params
    // 入参pushPaperFormatParam
    _pushPaperFormatParam   *ThirdPaperFormatPushOpenParam
}

// 初始化AlibabaDamaiMevOpenPushPaperformatRequest对象
func NewAlibabaDamaiMevOpenPushPaperformatRequest() *AlibabaDamaiMevOpenPushPaperformatRequest{
    return &AlibabaDamaiMevOpenPushPaperformatRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenPushPaperformatRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.push.paperformat"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenPushPaperformatRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PushPaperFormatParam Setter
// 入参pushPaperFormatParam
func (r *AlibabaDamaiMevOpenPushPaperformatRequest) SetPushPaperFormatParam(_pushPaperFormatParam *ThirdPaperFormatPushOpenParam) error {
    r._pushPaperFormatParam = _pushPaperFormatParam
    r.Set("push_paper_format_param", _pushPaperFormatParam)
    return nil
}

// PushPaperFormatParam Getter
func (r AlibabaDamaiMevOpenPushPaperformatRequest) GetPushPaperFormatParam() *ThirdPaperFormatPushOpenParam {
    return r._pushPaperFormatParam
}
