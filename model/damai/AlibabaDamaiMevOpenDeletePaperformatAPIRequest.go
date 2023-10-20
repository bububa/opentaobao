package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimevopendeletepaperformatAPIRequest 大麦换验平台-第三方对外开放-票纸版式接口deletePaperFormat API请求
// alibaba.damai.mev.open.delete.paperformat
//
// deletePaperFormat
type AlibabadamaimevopendeletepaperformatAPIRequest struct {
	model.Params
	// 入参deletePaperFormatParam
	_deletePaperFormatParam *TicketPaperFormatIdOpenParam
}

// NewAlibabadamaimevopendeletepaperformatRequest 初始化AlibabadamaimevopendeletepaperformatAPIRequest对象
func NewAlibabadamaimevopendeletepaperformatRequest() *AlibabadamaimevopendeletepaperformatAPIRequest {
	return &AlibabadamaimevopendeletepaperformatAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadamaimevopendeletepaperformatAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.delete.paperformat"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadamaimevopendeletepaperformatAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadamaimevopendeletepaperformatAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeletePaperFormatParam is DeletePaperFormatParam Setter
// 入参deletePaperFormatParam
func (r *AlibabadamaimevopendeletepaperformatAPIRequest) SetDeletePaperFormatParam(_deletePaperFormatParam *TicketPaperFormatIdOpenParam) error {
	r._deletePaperFormatParam = _deletePaperFormatParam
	r.Set("delete_paper_format_param", _deletePaperFormatParam)
	return nil
}

// GetDeletePaperFormatParam DeletePaperFormatParam Getter
func (r AlibabadamaimevopendeletepaperformatAPIRequest) GetDeletePaperFormatParam() *TicketPaperFormatIdOpenParam {
	return r._deletePaperFormatParam
}
