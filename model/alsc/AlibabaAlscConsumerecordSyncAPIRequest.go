package alsc

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscConsumerecordSyncAPIRequest) Reset() {
	r._paramRecordOpenSyncRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscConsumerecordSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.consumerecord.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscConsumerecordSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscConsumerecordSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAlscConsumerecordSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscConsumerecordSyncRequest()
	},
}

// GetAlibabaAlscConsumerecordSyncRequest 从 sync.Pool 获取 AlibabaAlscConsumerecordSyncAPIRequest
func GetAlibabaAlscConsumerecordSyncAPIRequest() *AlibabaAlscConsumerecordSyncAPIRequest {
	return poolAlibabaAlscConsumerecordSyncAPIRequest.Get().(*AlibabaAlscConsumerecordSyncAPIRequest)
}

// ReleaseAlibabaAlscConsumerecordSyncAPIRequest 将 AlibabaAlscConsumerecordSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscConsumerecordSyncAPIRequest(v *AlibabaAlscConsumerecordSyncAPIRequest) {
	v.Reset()
	poolAlibabaAlscConsumerecordSyncAPIRequest.Put(v)
}
