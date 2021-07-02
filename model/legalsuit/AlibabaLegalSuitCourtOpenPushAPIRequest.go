package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalSuitCourtOpenPushAPIRequest 开庭信息推送接口 API请求
// alibaba.legal.suit.court.open.push
//
// 供ISV推送开庭信息给集团诉讼
type AlibabaLegalSuitCourtOpenPushAPIRequest struct {
	model.Params
	// 开庭信息
	_courtInfoModel *CourtInfoModel
}

// NewAlibabaLegalSuitCourtOpenPushRequest 初始化AlibabaLegalSuitCourtOpenPushAPIRequest对象
func NewAlibabaLegalSuitCourtOpenPushRequest() *AlibabaLegalSuitCourtOpenPushAPIRequest {
	return &AlibabaLegalSuitCourtOpenPushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalSuitCourtOpenPushAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.suit.court.open.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalSuitCourtOpenPushAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CourtInfoModel Setter
// 开庭信息
func (r *AlibabaLegalSuitCourtOpenPushAPIRequest) SetCourtInfoModel(_courtInfoModel *CourtInfoModel) error {
	r._courtInfoModel = _courtInfoModel
	r.Set("court_info_model", _courtInfoModel)
	return nil
}

// Get CourtInfoModel Getter
func (r AlibabaLegalSuitCourtOpenPushAPIRequest) GetCourtInfoModel() *CourtInfoModel {
	return r._courtInfoModel
}
