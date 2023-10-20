package damai

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenDeletePaperformatAPIRequest 大麦换验平台-第三方对外开放-票纸版式接口deletePaperFormat API请求
// alibaba.damai.mev.open.delete.paperformat
//
// deletePaperFormat
type AlibabaDamaiMevOpenDeletePaperformatAPIRequest struct {
	model.Params
	// 入参deletePaperFormatParam
	_deletePaperFormatParam *TicketPaperFormatIdOpenParam
}

// NewAlibabaDamaiMevOpenDeletePaperformatRequest 初始化AlibabaDamaiMevOpenDeletePaperformatAPIRequest对象
func NewAlibabaDamaiMevOpenDeletePaperformatRequest() *AlibabaDamaiMevOpenDeletePaperformatAPIRequest {
	return &AlibabaDamaiMevOpenDeletePaperformatAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDamaiMevOpenDeletePaperformatAPIRequest) Reset() {
	r._deletePaperFormatParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenDeletePaperformatAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.delete.paperformat"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenDeletePaperformatAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDamaiMevOpenDeletePaperformatAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeletePaperFormatParam is DeletePaperFormatParam Setter
// 入参deletePaperFormatParam
func (r *AlibabaDamaiMevOpenDeletePaperformatAPIRequest) SetDeletePaperFormatParam(_deletePaperFormatParam *TicketPaperFormatIdOpenParam) error {
	r._deletePaperFormatParam = _deletePaperFormatParam
	r.Set("delete_paper_format_param", _deletePaperFormatParam)
	return nil
}

// GetDeletePaperFormatParam DeletePaperFormatParam Getter
func (r AlibabaDamaiMevOpenDeletePaperformatAPIRequest) GetDeletePaperFormatParam() *TicketPaperFormatIdOpenParam {
	return r._deletePaperFormatParam
}

var poolAlibabaDamaiMevOpenDeletePaperformatAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDamaiMevOpenDeletePaperformatRequest()
	},
}

// GetAlibabaDamaiMevOpenDeletePaperformatRequest 从 sync.Pool 获取 AlibabaDamaiMevOpenDeletePaperformatAPIRequest
func GetAlibabaDamaiMevOpenDeletePaperformatAPIRequest() *AlibabaDamaiMevOpenDeletePaperformatAPIRequest {
	return poolAlibabaDamaiMevOpenDeletePaperformatAPIRequest.Get().(*AlibabaDamaiMevOpenDeletePaperformatAPIRequest)
}

// ReleaseAlibabaDamaiMevOpenDeletePaperformatAPIRequest 将 AlibabaDamaiMevOpenDeletePaperformatAPIRequest 放入 sync.Pool
func ReleaseAlibabaDamaiMevOpenDeletePaperformatAPIRequest(v *AlibabaDamaiMevOpenDeletePaperformatAPIRequest) {
	v.Reset()
	poolAlibabaDamaiMevOpenDeletePaperformatAPIRequest.Put(v)
}
