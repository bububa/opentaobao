package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscConsumerecordSyncAPIRequest 外域订单同步本地生活消费记录 API请求
// alibaba.alsc.consumerecord.sync
//
// 外部第三方将本地生活app端下单数据同步到本地生活消费记录中心
type AlibabaAlscConsumerecordSyncAPIRequest struct {
	model.Params
	// 外部订单同步本地生活消费记录请求
	_paramRecordOpenSyncRequest *RecordOpenSyncRequest
}

// NewAlibabaAlscConsumerecordSyncRequest 初始化AlibabaAlscConsumerecordSyncAPIRequest对象
func NewAlibabaAlscConsumerecordSyncRequest() *AlibabaAlscConsumerecordSyncAPIRequest {
	return &AlibabaAlscConsumerecordSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscConsumerecordSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.consumerecord.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscConsumerecordSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamRecordOpenSyncRequest is ParamRecordOpenSyncRequest Setter
// 外部订单同步本地生活消费记录请求
func (r *AlibabaAlscConsumerecordSyncAPIRequest) SetParamRecordOpenSyncRequest(_paramRecordOpenSyncRequest *RecordOpenSyncRequest) error {
	r._paramRecordOpenSyncRequest = _paramRecordOpenSyncRequest
	r.Set("param_record_open_sync_request", _paramRecordOpenSyncRequest)
	return nil
}

// GetParamRecordOpenSyncRequest ParamRecordOpenSyncRequest Getter
func (r AlibabaAlscConsumerecordSyncAPIRequest) GetParamRecordOpenSyncRequest() *RecordOpenSyncRequest {
	return r._paramRecordOpenSyncRequest
}
