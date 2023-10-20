package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalscconsumerecordsyncAPIRequest 外域订单同步本地生活消费记录 API请求
// alibaba.alsc.consumerecord.sync
//
// 外部第三方将本地生活app端下单数据同步到本地生活消费记录中心
type AlibabaalscconsumerecordsyncAPIRequest struct {
	model.Params
	// 外部订单同步本地生活消费记录请求
	_paramRecordOpenSyncRequest *RecordOpenSyncRequest
}

// NewAlibabaalscconsumerecordsyncRequest 初始化AlibabaalscconsumerecordsyncAPIRequest对象
func NewAlibabaalscconsumerecordsyncRequest() *AlibabaalscconsumerecordsyncAPIRequest {
	return &AlibabaalscconsumerecordsyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalscconsumerecordsyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.consumerecord.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalscconsumerecordsyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalscconsumerecordsyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamRecordOpenSyncRequest is ParamRecordOpenSyncRequest Setter
// 外部订单同步本地生活消费记录请求
func (r *AlibabaalscconsumerecordsyncAPIRequest) SetParamRecordOpenSyncRequest(_paramRecordOpenSyncRequest *RecordOpenSyncRequest) error {
	r._paramRecordOpenSyncRequest = _paramRecordOpenSyncRequest
	r.Set("param_record_open_sync_request", _paramRecordOpenSyncRequest)
	return nil
}

// GetParamRecordOpenSyncRequest ParamRecordOpenSyncRequest Getter
func (r AlibabaalscconsumerecordsyncAPIRequest) GetParamRecordOpenSyncRequest() *RecordOpenSyncRequest {
	return r._paramRecordOpenSyncRequest
}
