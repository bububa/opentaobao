package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMjOcBigposBanksaleQueryAPIRequest 大pos银行卡查账接口 API请求
// alibaba.mj.oc.bigpos.banksale.query
//
// 大pos银行卡查账接口，给收银员查询银行卡销售记录，便于调账
type AlibabaMjOcBigposBanksaleQueryAPIRequest struct {
	model.Params
	// 开始时间
	_startTime string
	// 卡号
	_cardNo string
	// 外部门店号
	_outStoreNo string
	// 结束时间
	_endTime string
}

// NewAlibabaMjOcBigposBanksaleQueryRequest 初始化AlibabaMjOcBigposBanksaleQueryAPIRequest对象
func NewAlibabaMjOcBigposBanksaleQueryRequest() *AlibabaMjOcBigposBanksaleQueryAPIRequest {
	return &AlibabaMjOcBigposBanksaleQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMjOcBigposBanksaleQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.mj.oc.bigpos.banksale.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMjOcBigposBanksaleQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMjOcBigposBanksaleQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStartTime is StartTime Setter
// 开始时间
func (r *AlibabaMjOcBigposBanksaleQueryAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r AlibabaMjOcBigposBanksaleQueryAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetCardNo is CardNo Setter
// 卡号
func (r *AlibabaMjOcBigposBanksaleQueryAPIRequest) SetCardNo(_cardNo string) error {
	r._cardNo = _cardNo
	r.Set("card_no", _cardNo)
	return nil
}

// GetCardNo CardNo Getter
func (r AlibabaMjOcBigposBanksaleQueryAPIRequest) GetCardNo() string {
	return r._cardNo
}

// SetOutStoreNo is OutStoreNo Setter
// 外部门店号
func (r *AlibabaMjOcBigposBanksaleQueryAPIRequest) SetOutStoreNo(_outStoreNo string) error {
	r._outStoreNo = _outStoreNo
	r.Set("out_store_no", _outStoreNo)
	return nil
}

// GetOutStoreNo OutStoreNo Getter
func (r AlibabaMjOcBigposBanksaleQueryAPIRequest) GetOutStoreNo() string {
	return r._outStoreNo
}

// SetEndTime is EndTime Setter
// 结束时间
func (r *AlibabaMjOcBigposBanksaleQueryAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r AlibabaMjOcBigposBanksaleQueryAPIRequest) GetEndTime() string {
	return r._endTime
}
