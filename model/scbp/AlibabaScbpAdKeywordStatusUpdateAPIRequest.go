package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadkeywordstatusupdateAPIRequest 关键词启动暂停推广 API请求
// alibaba.scbp.ad.keyword.status.update
//
// 关键词启动暂停推广
type AlibabascbpadkeywordstatusupdateAPIRequest struct {
	model.Params
	// 只能取ascci字符
	_adKeyword string
	// 只能去in_promotion或者stopped
	_status string
}

// NewAlibabascbpadkeywordstatusupdateRequest 初始化AlibabascbpadkeywordstatusupdateAPIRequest对象
func NewAlibabascbpadkeywordstatusupdateRequest() *AlibabascbpadkeywordstatusupdateAPIRequest {
	return &AlibabascbpadkeywordstatusupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabascbpadkeywordstatusupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.status.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabascbpadkeywordstatusupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabascbpadkeywordstatusupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAdKeyword is AdKeyword Setter
// 只能取ascci字符
func (r *AlibabascbpadkeywordstatusupdateAPIRequest) SetAdKeyword(_adKeyword string) error {
	r._adKeyword = _adKeyword
	r.Set("ad_keyword", _adKeyword)
	return nil
}

// GetAdKeyword AdKeyword Getter
func (r AlibabascbpadkeywordstatusupdateAPIRequest) GetAdKeyword() string {
	return r._adKeyword
}

// SetStatus is Status Setter
// 只能去in_promotion或者stopped
func (r *AlibabascbpadkeywordstatusupdateAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabascbpadkeywordstatusupdateAPIRequest) GetStatus() string {
	return r._status
}
