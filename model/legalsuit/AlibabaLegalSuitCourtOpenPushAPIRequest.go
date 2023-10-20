package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegalsuitcourtopenpushAPIRequest 开庭信息推送接口 API请求
// alibaba.legal.suit.court.open.push
//
// 供ISV推送开庭信息给集团诉讼
type AlibabalegalsuitcourtopenpushAPIRequest struct {
	model.Params
	// 开庭信息
	_courtInfoModel *CourtInfoModel
}

// NewAlibabalegalsuitcourtopenpushRequest 初始化AlibabalegalsuitcourtopenpushAPIRequest对象
func NewAlibabalegalsuitcourtopenpushRequest() *AlibabalegalsuitcourtopenpushAPIRequest {
	return &AlibabalegalsuitcourtopenpushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalegalsuitcourtopenpushAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.suit.court.open.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalegalsuitcourtopenpushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalegalsuitcourtopenpushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCourtInfoModel is CourtInfoModel Setter
// 开庭信息
func (r *AlibabalegalsuitcourtopenpushAPIRequest) SetCourtInfoModel(_courtInfoModel *CourtInfoModel) error {
	r._courtInfoModel = _courtInfoModel
	r.Set("court_info_model", _courtInfoModel)
	return nil
}

// GetCourtInfoModel CourtInfoModel Getter
func (r AlibabalegalsuitcourtopenpushAPIRequest) GetCourtInfoModel() *CourtInfoModel {
	return r._courtInfoModel
}
