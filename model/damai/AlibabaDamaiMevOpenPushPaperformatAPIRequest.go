package damai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDamaiMevOpenPushPaperformatAPIRequest
大麦换验平台-第三方对外开放-票纸版式接口pushPaperFormat API请求
alibaba.damai.mev.open.push.paperformat

pushPaperFormat */
type AlibabaDamaiMevOpenPushPaperformatAPIRequest struct {
	model.Params
	// 入参pushPaperFormatParam
	_pushPaperFormatParam *ThirdPaperFormatPushOpenParam
}

// NewAlibabaDamaiMevOpenPushPaperformatRequest 初始化AlibabaDamaiMevOpenPushPaperformatAPIRequest对象
func NewAlibabaDamaiMevOpenPushPaperformatRequest() *AlibabaDamaiMevOpenPushPaperformatAPIRequest {
	return &AlibabaDamaiMevOpenPushPaperformatAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenPushPaperformatAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.mev.open.push.paperformat"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenPushPaperformatAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is PushPaperFormatParam Setter
// 入参pushPaperFormatParam
func (r *AlibabaDamaiMevOpenPushPaperformatAPIRequest) SetPushPaperFormatParam(_pushPaperFormatParam *ThirdPaperFormatPushOpenParam) error {
	r._pushPaperFormatParam = _pushPaperFormatParam
	r.Set("push_paper_format_param", _pushPaperFormatParam)
	return nil
}

// Get PushPaperFormatParam Getter
func (r AlibabaDamaiMevOpenPushPaperformatAPIRequest) GetPushPaperFormatParam() *ThirdPaperFormatPushOpenParam {
	return r._pushPaperFormatParam
}
