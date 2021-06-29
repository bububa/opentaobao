package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-票纸版式接口pushPaperFormat APIRequest
alibaba.damai.mev.open.push.paperformat

pushPaperFormat
*/
type AlibabaDamaiMevOpenPushPaperformatRequest struct {
    model.Params

    // 入参pushPaperFormatParam
    pushPaperFormatParam   *ThirdPaperFormatPushOpenParam 

}

func NewAlibabaDamaiMevOpenPushPaperformatRequest() *AlibabaDamaiMevOpenPushPaperformatRequest{
    return &AlibabaDamaiMevOpenPushPaperformatRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDamaiMevOpenPushPaperformatRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.push.paperformat"
}

func (r AlibabaDamaiMevOpenPushPaperformatRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDamaiMevOpenPushPaperformatRequest) SetPushPaperFormatParam(pushPaperFormatParam *ThirdPaperFormatPushOpenParam) error {
    r.pushPaperFormatParam = pushPaperFormatParam
    r.Set("push_paper_format_param", pushPaperFormatParam)
    return nil
}

func (r AlibabaDamaiMevOpenPushPaperformatRequest) GetPushPaperFormatParam() *ThirdPaperFormatPushOpenParam {
    return r.pushPaperFormatParam
}

