package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimevopenpushpaperformatAPIRequest 大麦换验平台-第三方对外开放-票纸版式接口pushPaperFormat API请求
// alibaba.damai.mev.open.push.paperformat
//
// pushPaperFormat
type AlibabadamaimevopenpushpaperformatAPIRequest struct {
	model.Params
	// 入参pushPaperFormatParam
	_pushPaperFormatParam *ThirdPaperFormatPushOpenParam
}

// NewAlibabadamaimevopenpushpaperformatRequest 初始化AlibabadamaimevopenpushpaperformatAPIRequest对象
func NewAlibabadamaimevopenpushpaperformatRequest() *AlibabadamaimevopenpushpaperformatAPIRequest {
	return &AlibabadamaimevopenpushpaperformatAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadamaimevopenpushpaperformatAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.push.paperformat"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadamaimevopenpushpaperformatAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadamaimevopenpushpaperformatAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPushPaperFormatParam is PushPaperFormatParam Setter
// 入参pushPaperFormatParam
func (r *AlibabadamaimevopenpushpaperformatAPIRequest) SetPushPaperFormatParam(_pushPaperFormatParam *ThirdPaperFormatPushOpenParam) error {
	r._pushPaperFormatParam = _pushPaperFormatParam
	r.Set("push_paper_format_param", _pushPaperFormatParam)
	return nil
}

// GetPushPaperFormatParam PushPaperFormatParam Getter
func (r AlibabadamaimevopenpushpaperformatAPIRequest) GetPushPaperFormatParam() *ThirdPaperFormatPushOpenParam {
	return r._pushPaperFormatParam
}
