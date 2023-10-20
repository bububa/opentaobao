package aliospay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliyunAliosPayRecordListAPIRequest 支付记录批量查询接口 API请求
// aliyun.alios.pay.record.list
//
// 商户用来对账的接口
type AliyunAliosPayRecordListAPIRequest struct {
	model.Params
	// 请求参数
	_searchRecordRequest *SearchRecordRequest
}

// NewAliyunAliosPayRecordListRequest 初始化AliyunAliosPayRecordListAPIRequest对象
func NewAliyunAliosPayRecordListRequest() *AliyunAliosPayRecordListAPIRequest {
	return &AliyunAliosPayRecordListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliyunAliosPayRecordListAPIRequest) GetApiMethodName() string {
	return "aliyun.alios.pay.record.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliyunAliosPayRecordListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliyunAliosPayRecordListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSearchRecordRequest is SearchRecordRequest Setter
// 请求参数
func (r *AliyunAliosPayRecordListAPIRequest) SetSearchRecordRequest(_searchRecordRequest *SearchRecordRequest) error {
	r._searchRecordRequest = _searchRecordRequest
	r.Set("search_record_request", _searchRecordRequest)
	return nil
}

// GetSearchRecordRequest SearchRecordRequest Getter
func (r AliyunAliosPayRecordListAPIRequest) GetSearchRecordRequest() *SearchRecordRequest {
	return r._searchRecordRequest
}
