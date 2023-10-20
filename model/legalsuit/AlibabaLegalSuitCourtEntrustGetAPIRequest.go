package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegalsuitcourtentrustgetAPIRequest 委托开庭服务查询 API请求
// alibaba.legal.suit.court.entrust.get
//
// 查询委托开庭信息
type AlibabalegalsuitcourtentrustgetAPIRequest struct {
	model.Params
	// 案件ID
	_suitId int64
	// 委托ID
	_entrustId int64
}

// NewAlibabalegalsuitcourtentrustgetRequest 初始化AlibabalegalsuitcourtentrustgetAPIRequest对象
func NewAlibabalegalsuitcourtentrustgetRequest() *AlibabalegalsuitcourtentrustgetAPIRequest {
	return &AlibabalegalsuitcourtentrustgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalegalsuitcourtentrustgetAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.suit.court.entrust.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalegalsuitcourtentrustgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalegalsuitcourtentrustgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSuitId is SuitId Setter
// 案件ID
func (r *AlibabalegalsuitcourtentrustgetAPIRequest) SetSuitId(_suitId int64) error {
	r._suitId = _suitId
	r.Set("suit_id", _suitId)
	return nil
}

// GetSuitId SuitId Getter
func (r AlibabalegalsuitcourtentrustgetAPIRequest) GetSuitId() int64 {
	return r._suitId
}

// SetEntrustId is EntrustId Setter
// 委托ID
func (r *AlibabalegalsuitcourtentrustgetAPIRequest) SetEntrustId(_entrustId int64) error {
	r._entrustId = _entrustId
	r.Set("entrust_id", _entrustId)
	return nil
}

// GetEntrustId EntrustId Getter
func (r AlibabalegalsuitcourtentrustgetAPIRequest) GetEntrustId() int64 {
	return r._entrustId
}
