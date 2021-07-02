package zqs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaZqsFulfillCompleteAPIRequest 周期购履约完成接口 API请求
// alibaba.zqs.fulfill.complete
//
// 周期购履约完成接口
type AlibabaZqsFulfillCompleteAPIRequest struct {
	model.Params
	// 第几期
	_sequenceNo int64
	// 交易单号
	_mainBizOrderId int64
	// 交易子单号
	_subBizOrderId int64
}

// NewAlibabaZqsFulfillCompleteRequest 初始化AlibabaZqsFulfillCompleteAPIRequest对象
func NewAlibabaZqsFulfillCompleteRequest() *AlibabaZqsFulfillCompleteAPIRequest {
	return &AlibabaZqsFulfillCompleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaZqsFulfillCompleteAPIRequest) GetApiMethodName() string {
	return "alibaba.zqs.fulfill.complete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaZqsFulfillCompleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SequenceNo Setter
// 第几期
func (r *AlibabaZqsFulfillCompleteAPIRequest) SetSequenceNo(_sequenceNo int64) error {
	r._sequenceNo = _sequenceNo
	r.Set("sequence_no", _sequenceNo)
	return nil
}

// Get SequenceNo Getter
func (r AlibabaZqsFulfillCompleteAPIRequest) GetSequenceNo() int64 {
	return r._sequenceNo
}

// Set is MainBizOrderId Setter
// 交易单号
func (r *AlibabaZqsFulfillCompleteAPIRequest) SetMainBizOrderId(_mainBizOrderId int64) error {
	r._mainBizOrderId = _mainBizOrderId
	r.Set("main_biz_order_id", _mainBizOrderId)
	return nil
}

// Get MainBizOrderId Getter
func (r AlibabaZqsFulfillCompleteAPIRequest) GetMainBizOrderId() int64 {
	return r._mainBizOrderId
}

// Set is SubBizOrderId Setter
// 交易子单号
func (r *AlibabaZqsFulfillCompleteAPIRequest) SetSubBizOrderId(_subBizOrderId int64) error {
	r._subBizOrderId = _subBizOrderId
	r.Set("sub_biz_order_id", _subBizOrderId)
	return nil
}

// Get SubBizOrderId Getter
func (r AlibabaZqsFulfillCompleteAPIRequest) GetSubBizOrderId() int64 {
	return r._subBizOrderId
}
