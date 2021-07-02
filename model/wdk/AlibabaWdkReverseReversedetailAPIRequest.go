package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkReverseReversedetailAPIRequest 退款详情 API请求
// alibaba.wdk.reverse.reversedetail
//
// 退款详情
type AlibabaWdkReverseReversedetailAPIRequest struct {
	model.Params
	// 退款单id
	_reverseId string
}

// NewAlibabaWdkReverseReversedetailRequest 初始化AlibabaWdkReverseReversedetailAPIRequest对象
func NewAlibabaWdkReverseReversedetailRequest() *AlibabaWdkReverseReversedetailAPIRequest {
	return &AlibabaWdkReverseReversedetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkReverseReversedetailAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.reverse.reversedetail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkReverseReversedetailAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetReverseId is ReverseId Setter
// 退款单id
func (r *AlibabaWdkReverseReversedetailAPIRequest) SetReverseId(_reverseId string) error {
	r._reverseId = _reverseId
	r.Set("reverse_id", _reverseId)
	return nil
}

// GetReverseId ReverseId Getter
func (r AlibabaWdkReverseReversedetailAPIRequest) GetReverseId() string {
	return r._reverseId
}
