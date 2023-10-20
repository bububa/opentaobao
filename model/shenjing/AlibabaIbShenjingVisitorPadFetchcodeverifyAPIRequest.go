package shenjing

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaibshenjingvisitorpadfetchcodeverifyAPIRequest 访客通过PAD提交访客码 API请求
// alibaba.ib.shenjing.visitor.pad.fetchcodeverify
//
// 访客通过PAD提交访客码，录脸进入园区。
type AlibabaibshenjingvisitorpadfetchcodeverifyAPIRequest struct {
	model.Params
	// 终端ID
	_termId string
	// 访客码
	_visitorCode int64
}

// NewAlibabaibshenjingvisitorpadfetchcodeverifyRequest 初始化AlibabaibshenjingvisitorpadfetchcodeverifyAPIRequest对象
func NewAlibabaibshenjingvisitorpadfetchcodeverifyRequest() *AlibabaibshenjingvisitorpadfetchcodeverifyAPIRequest {
	return &AlibabaibshenjingvisitorpadfetchcodeverifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaibshenjingvisitorpadfetchcodeverifyAPIRequest) GetApiMethodName() string {
	return "alibaba.ib.shenjing.visitor.pad.fetchcodeverify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaibshenjingvisitorpadfetchcodeverifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaibshenjingvisitorpadfetchcodeverifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTermId is TermId Setter
// 终端ID
func (r *AlibabaibshenjingvisitorpadfetchcodeverifyAPIRequest) SetTermId(_termId string) error {
	r._termId = _termId
	r.Set("term_id", _termId)
	return nil
}

// GetTermId TermId Getter
func (r AlibabaibshenjingvisitorpadfetchcodeverifyAPIRequest) GetTermId() string {
	return r._termId
}

// SetVisitorCode is VisitorCode Setter
// 访客码
func (r *AlibabaibshenjingvisitorpadfetchcodeverifyAPIRequest) SetVisitorCode(_visitorCode int64) error {
	r._visitorCode = _visitorCode
	r.Set("visitor_code", _visitorCode)
	return nil
}

// GetVisitorCode VisitorCode Getter
func (r AlibabaibshenjingvisitorpadfetchcodeverifyAPIRequest) GetVisitorCode() int64 {
	return r._visitorCode
}
