package legalsuit

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLegalSuitCourtOpenPushAPIRequest) Reset() {
	r._courtInfoModel = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalSuitCourtOpenPushAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.suit.court.open.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalSuitCourtOpenPushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLegalSuitCourtOpenPushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCourtInfoModel is CourtInfoModel Setter
// 开庭信息
func (r *AlibabaLegalSuitCourtOpenPushAPIRequest) SetCourtInfoModel(_courtInfoModel *CourtInfoModel) error {
	r._courtInfoModel = _courtInfoModel
	r.Set("court_info_model", _courtInfoModel)
	return nil
}

// GetCourtInfoModel CourtInfoModel Getter
func (r AlibabaLegalSuitCourtOpenPushAPIRequest) GetCourtInfoModel() *CourtInfoModel {
	return r._courtInfoModel
}

var poolAlibabaLegalSuitCourtOpenPushAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLegalSuitCourtOpenPushRequest()
	},
}

// GetAlibabaLegalSuitCourtOpenPushRequest 从 sync.Pool 获取 AlibabaLegalSuitCourtOpenPushAPIRequest
func GetAlibabaLegalSuitCourtOpenPushAPIRequest() *AlibabaLegalSuitCourtOpenPushAPIRequest {
	return poolAlibabaLegalSuitCourtOpenPushAPIRequest.Get().(*AlibabaLegalSuitCourtOpenPushAPIRequest)
}

// ReleaseAlibabaLegalSuitCourtOpenPushAPIRequest 将 AlibabaLegalSuitCourtOpenPushAPIRequest 放入 sync.Pool
func ReleaseAlibabaLegalSuitCourtOpenPushAPIRequest(v *AlibabaLegalSuitCourtOpenPushAPIRequest) {
	v.Reset()
	poolAlibabaLegalSuitCourtOpenPushAPIRequest.Put(v)
}
