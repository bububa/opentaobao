package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-票纸版式接口deletePaperFormat API请求
alibaba.damai.mev.open.delete.paperformat

deletePaperFormat
*/
type AlibabaDamaiMevOpenDeletePaperformatRequest struct {
    model.Params
    // 入参deletePaperFormatParam
    _deletePaperFormatParam   *TicketPaperFormatIdOpenParam
}

// 初始化AlibabaDamaiMevOpenDeletePaperformatRequest对象
func NewAlibabaDamaiMevOpenDeletePaperformatRequest() *AlibabaDamaiMevOpenDeletePaperformatRequest{
    return &AlibabaDamaiMevOpenDeletePaperformatRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenDeletePaperformatRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.delete.paperformat"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenDeletePaperformatRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeletePaperFormatParam Setter
// 入参deletePaperFormatParam
func (r *AlibabaDamaiMevOpenDeletePaperformatRequest) SetDeletePaperFormatParam(_deletePaperFormatParam *TicketPaperFormatIdOpenParam) error {
    r._deletePaperFormatParam = _deletePaperFormatParam
    r.Set("delete_paper_format_param", _deletePaperFormatParam)
    return nil
}

// DeletePaperFormatParam Getter
func (r AlibabaDamaiMevOpenDeletePaperformatRequest) GetDeletePaperFormatParam() *TicketPaperFormatIdOpenParam {
    return r._deletePaperFormatParam
}
