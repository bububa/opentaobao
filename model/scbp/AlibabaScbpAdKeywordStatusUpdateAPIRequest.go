package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdKeywordStatusUpdateAPIRequest 关键词启动暂停推广 API请求
// alibaba.scbp.ad.keyword.status.update
//
// 关键词启动暂停推广
type AlibabaScbpAdKeywordStatusUpdateAPIRequest struct {
	model.Params
	// 只能取ascci字符
	_adKeyword string
	// 只能去in_promotion或者stopped
	_status string
}

// NewAlibabaScbpAdKeywordStatusUpdateRequest 初始化AlibabaScbpAdKeywordStatusUpdateAPIRequest对象
func NewAlibabaScbpAdKeywordStatusUpdateRequest() *AlibabaScbpAdKeywordStatusUpdateAPIRequest {
	return &AlibabaScbpAdKeywordStatusUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpAdKeywordStatusUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.ad.keyword.status.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpAdKeywordStatusUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AdKeyword Setter
// 只能取ascci字符
func (r *AlibabaScbpAdKeywordStatusUpdateAPIRequest) SetAdKeyword(_adKeyword string) error {
	r._adKeyword = _adKeyword
	r.Set("ad_keyword", _adKeyword)
	return nil
}

// Get AdKeyword Getter
func (r AlibabaScbpAdKeywordStatusUpdateAPIRequest) GetAdKeyword() string {
	return r._adKeyword
}

// Set is Status Setter
// 只能去in_promotion或者stopped
func (r *AlibabaScbpAdKeywordStatusUpdateAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// Get Status Getter
func (r AlibabaScbpAdKeywordStatusUpdateAPIRequest) GetStatus() string {
	return r._status
}
