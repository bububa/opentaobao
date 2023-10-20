package zqs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabazqsfulfillcompleteAPIRequest 周期购履约完成接口 API请求
// alibaba.zqs.fulfill.complete
//
// 周期购履约完成接口
type AlibabazqsfulfillcompleteAPIRequest struct {
	model.Params
	// 第几期
	_sequenceNo int64
	// 交易单号
	_mainBizOrderId int64
	// 交易子单号
	_subBizOrderId int64
}

// NewAlibabazqsfulfillcompleteRequest 初始化AlibabazqsfulfillcompleteAPIRequest对象
func NewAlibabazqsfulfillcompleteRequest() *AlibabazqsfulfillcompleteAPIRequest {
	return &AlibabazqsfulfillcompleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabazqsfulfillcompleteAPIRequest) GetApiMethodName() string {
	return "alibaba.zqs.fulfill.complete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabazqsfulfillcompleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabazqsfulfillcompleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSequenceNo is SequenceNo Setter
// 第几期
func (r *AlibabazqsfulfillcompleteAPIRequest) SetSequenceNo(_sequenceNo int64) error {
	r._sequenceNo = _sequenceNo
	r.Set("sequence_no", _sequenceNo)
	return nil
}

// GetSequenceNo SequenceNo Getter
func (r AlibabazqsfulfillcompleteAPIRequest) GetSequenceNo() int64 {
	return r._sequenceNo
}

// SetMainBizOrderId is MainBizOrderId Setter
// 交易单号
func (r *AlibabazqsfulfillcompleteAPIRequest) SetMainBizOrderId(_mainBizOrderId int64) error {
	r._mainBizOrderId = _mainBizOrderId
	r.Set("main_biz_order_id", _mainBizOrderId)
	return nil
}

// GetMainBizOrderId MainBizOrderId Getter
func (r AlibabazqsfulfillcompleteAPIRequest) GetMainBizOrderId() int64 {
	return r._mainBizOrderId
}

// SetSubBizOrderId is SubBizOrderId Setter
// 交易子单号
func (r *AlibabazqsfulfillcompleteAPIRequest) SetSubBizOrderId(_subBizOrderId int64) error {
	r._subBizOrderId = _subBizOrderId
	r.Set("sub_biz_order_id", _subBizOrderId)
	return nil
}

// GetSubBizOrderId SubBizOrderId Getter
func (r AlibabazqsfulfillcompleteAPIRequest) GetSubBizOrderId() int64 {
	return r._subBizOrderId
}
