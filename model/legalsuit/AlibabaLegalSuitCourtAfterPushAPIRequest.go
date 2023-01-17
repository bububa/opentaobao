package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalSuitCourtAfterPushAPIRequest 更新或者新增庭后信息 API请求
// alibaba.legal.suit.court.after.push
//
// 供外部ISV供应商 推送庭后信息给集团诉讼
type AlibabaLegalSuitCourtAfterPushAPIRequest struct {
	model.Params
	// 庭后信息
	_afterCourtInfoModel *AfterCourtInfoModel
}

// NewAlibabaLegalSuitCourtAfterPushRequest 初始化AlibabaLegalSuitCourtAfterPushAPIRequest对象
func NewAlibabaLegalSuitCourtAfterPushRequest() *AlibabaLegalSuitCourtAfterPushAPIRequest {
	return &AlibabaLegalSuitCourtAfterPushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalSuitCourtAfterPushAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.suit.court.after.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalSuitCourtAfterPushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLegalSuitCourtAfterPushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAfterCourtInfoModel is AfterCourtInfoModel Setter
// 庭后信息
func (r *AlibabaLegalSuitCourtAfterPushAPIRequest) SetAfterCourtInfoModel(_afterCourtInfoModel *AfterCourtInfoModel) error {
	r._afterCourtInfoModel = _afterCourtInfoModel
	r.Set("after_court_info_model", _afterCourtInfoModel)
	return nil
}

// GetAfterCourtInfoModel AfterCourtInfoModel Getter
func (r AlibabaLegalSuitCourtAfterPushAPIRequest) GetAfterCourtInfoModel() *AfterCourtInfoModel {
	return r._afterCourtInfoModel
}
