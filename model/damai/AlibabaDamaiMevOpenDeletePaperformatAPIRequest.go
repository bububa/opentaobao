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
type AlibabaDamaiMevOpenDeletePaperformatAPIRequest struct {
    model.Params
    // 入参deletePaperFormatParam
    _deletePaperFormatParam   *TicketPaperFormatIdOpenParam
}

// 初始化AlibabaDamaiMevOpenDeletePaperformatAPIRequest对象
func NewAlibabaDamaiMevOpenDeletePaperformatRequest() *AlibabaDamaiMevOpenDeletePaperformatAPIRequest{
    return &AlibabaDamaiMevOpenDeletePaperformatAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenDeletePaperformatAPIRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.delete.paperformat"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenDeletePaperformatAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeletePaperFormatParam Setter
// 入参deletePaperFormatParam
func (r *AlibabaDamaiMevOpenDeletePaperformatAPIRequest) SetDeletePaperFormatParam(_deletePaperFormatParam *TicketPaperFormatIdOpenParam) error {
    r._deletePaperFormatParam = _deletePaperFormatParam
    r.Set("delete_paper_format_param", _deletePaperFormatParam)
    return nil
}

// DeletePaperFormatParam Getter
func (r AlibabaDamaiMevOpenDeletePaperformatAPIRequest) GetDeletePaperFormatParam() *TicketPaperFormatIdOpenParam {
    return r._deletePaperFormatParam
}
