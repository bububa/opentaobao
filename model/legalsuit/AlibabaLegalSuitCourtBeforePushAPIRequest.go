package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegalsuitcourtbeforepushAPIRequest 更新或保存庭前信息 API请求
// alibaba.legal.suit.court.before.push
//
// 更新或者保存庭前信息
type AlibabalegalsuitcourtbeforepushAPIRequest struct {
	model.Params
	// 庭前信息
	_beforeCourtModel *BeforeCourtModel
}

// NewAlibabalegalsuitcourtbeforepushRequest 初始化AlibabalegalsuitcourtbeforepushAPIRequest对象
func NewAlibabalegalsuitcourtbeforepushRequest() *AlibabalegalsuitcourtbeforepushAPIRequest {
	return &AlibabalegalsuitcourtbeforepushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalegalsuitcourtbeforepushAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.suit.court.before.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalegalsuitcourtbeforepushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalegalsuitcourtbeforepushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBeforeCourtModel is BeforeCourtModel Setter
// 庭前信息
func (r *AlibabalegalsuitcourtbeforepushAPIRequest) SetBeforeCourtModel(_beforeCourtModel *BeforeCourtModel) error {
	r._beforeCourtModel = _beforeCourtModel
	r.Set("before_court_model", _beforeCourtModel)
	return nil
}

// GetBeforeCourtModel BeforeCourtModel Getter
func (r AlibabalegalsuitcourtbeforepushAPIRequest) GetBeforeCourtModel() *BeforeCourtModel {
	return r._beforeCourtModel
}
