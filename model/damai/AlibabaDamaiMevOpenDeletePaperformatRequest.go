package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-票纸版式接口deletePaperFormat APIRequest
alibaba.damai.mev.open.delete.paperformat

deletePaperFormat
*/
type AlibabaDamaiMevOpenDeletePaperformatRequest struct {
    model.Params

    // 入参deletePaperFormatParam
    deletePaperFormatParam   *TicketPaperFormatIdOpenParam 

}

func NewAlibabaDamaiMevOpenDeletePaperformatRequest() *AlibabaDamaiMevOpenDeletePaperformatRequest{
    return &AlibabaDamaiMevOpenDeletePaperformatRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDamaiMevOpenDeletePaperformatRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.delete.paperformat"
}

func (r AlibabaDamaiMevOpenDeletePaperformatRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDamaiMevOpenDeletePaperformatRequest) SetDeletePaperFormatParam(deletePaperFormatParam *TicketPaperFormatIdOpenParam) error {
    r.deletePaperFormatParam = deletePaperFormatParam
    r.Set("delete_paper_format_param", deletePaperFormatParam)
    return nil
}

func (r AlibabaDamaiMevOpenDeletePaperformatRequest) GetDeletePaperFormatParam() *TicketPaperFormatIdOpenParam {
    return r.deletePaperFormatParam
}

