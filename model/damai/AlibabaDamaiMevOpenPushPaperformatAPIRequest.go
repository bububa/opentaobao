package damai

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMevOpenPushPaperformatAPIRequest 大麦换验平台-第三方对外开放-票纸版式接口pushPaperFormat API请求
// alibaba.damai.mev.open.push.paperformat
//
// pushPaperFormat
type AlibabaDamaiMevOpenPushPaperformatAPIRequest struct {
	model.Params
	// 入参pushPaperFormatParam
	_pushPaperFormatParam *ThirdPaperFormatPushOpenParam
}

// NewAlibabaDamaiMevOpenPushPaperformatRequest 初始化AlibabaDamaiMevOpenPushPaperformatAPIRequest对象
func NewAlibabaDamaiMevOpenPushPaperformatRequest() *AlibabaDamaiMevOpenPushPaperformatAPIRequest {
	return &AlibabaDamaiMevOpenPushPaperformatAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDamaiMevOpenPushPaperformatAPIRequest) Reset() {
	r._pushPaperFormatParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenPushPaperformatAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.push.paperformat"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenPushPaperformatAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDamaiMevOpenPushPaperformatAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPushPaperFormatParam is PushPaperFormatParam Setter
// 入参pushPaperFormatParam
func (r *AlibabaDamaiMevOpenPushPaperformatAPIRequest) SetPushPaperFormatParam(_pushPaperFormatParam *ThirdPaperFormatPushOpenParam) error {
	r._pushPaperFormatParam = _pushPaperFormatParam
	r.Set("push_paper_format_param", _pushPaperFormatParam)
	return nil
}

// GetPushPaperFormatParam PushPaperFormatParam Getter
func (r AlibabaDamaiMevOpenPushPaperformatAPIRequest) GetPushPaperFormatParam() *ThirdPaperFormatPushOpenParam {
	return r._pushPaperFormatParam
}

var poolAlibabaDamaiMevOpenPushPaperformatAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDamaiMevOpenPushPaperformatRequest()
	},
}

// GetAlibabaDamaiMevOpenPushPaperformatRequest 从 sync.Pool 获取 AlibabaDamaiMevOpenPushPaperformatAPIRequest
func GetAlibabaDamaiMevOpenPushPaperformatAPIRequest() *AlibabaDamaiMevOpenPushPaperformatAPIRequest {
	return poolAlibabaDamaiMevOpenPushPaperformatAPIRequest.Get().(*AlibabaDamaiMevOpenPushPaperformatAPIRequest)
}

// ReleaseAlibabaDamaiMevOpenPushPaperformatAPIRequest 将 AlibabaDamaiMevOpenPushPaperformatAPIRequest 放入 sync.Pool
func ReleaseAlibabaDamaiMevOpenPushPaperformatAPIRequest(v *AlibabaDamaiMevOpenPushPaperformatAPIRequest) {
	v.Reset()
	poolAlibabaDamaiMevOpenPushPaperformatAPIRequest.Put(v)
}
