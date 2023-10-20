package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegalsuitsealpushAPIRequest 法宝推送用印 API请求
// alibaba.legal.suit.seal.push
//
// 法宝推送用印
type AlibabalegalsuitsealpushAPIRequest struct {
	model.Params
	// 最外层list
	_sealTaskModel *SealTaskModel
}

// NewAlibabalegalsuitsealpushRequest 初始化AlibabalegalsuitsealpushAPIRequest对象
func NewAlibabalegalsuitsealpushRequest() *AlibabalegalsuitsealpushAPIRequest {
	return &AlibabalegalsuitsealpushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalegalsuitsealpushAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.suit.seal.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalegalsuitsealpushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalegalsuitsealpushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSealTaskModel is SealTaskModel Setter
// 最外层list
func (r *AlibabalegalsuitsealpushAPIRequest) SetSealTaskModel(_sealTaskModel *SealTaskModel) error {
	r._sealTaskModel = _sealTaskModel
	r.Set("seal_task_model", _sealTaskModel)
	return nil
}

// GetSealTaskModel SealTaskModel Getter
func (r AlibabalegalsuitsealpushAPIRequest) GetSealTaskModel() *SealTaskModel {
	return r._sealTaskModel
}
