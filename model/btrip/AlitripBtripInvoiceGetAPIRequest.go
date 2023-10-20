package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripinvoicegetAPIRequest 获取用户可用发票列表 API请求
// alitrip.btrip.invoice.get
//
// 差旅申请用户获取可用发票列表
type AlitripbtripinvoicegetAPIRequest struct {
	model.Params
	// 企业id
	_corpId string
	// 用户id
	_userId string
}

// NewAlitripbtripinvoicegetRequest 初始化AlitripbtripinvoicegetAPIRequest对象
func NewAlitripbtripinvoicegetRequest() *AlitripbtripinvoicegetAPIRequest {
	return &AlitripbtripinvoicegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripinvoicegetAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.invoice.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripinvoicegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripinvoicegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCorpId is CorpId Setter
// 企业id
func (r *AlitripbtripinvoicegetAPIRequest) SetCorpId(_corpId string) error {
	r._corpId = _corpId
	r.Set("corp_id", _corpId)
	return nil
}

// GetCorpId CorpId Getter
func (r AlitripbtripinvoicegetAPIRequest) GetCorpId() string {
	return r._corpId
}

// SetUserId is UserId Setter
// 用户id
func (r *AlitripbtripinvoicegetAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlitripbtripinvoicegetAPIRequest) GetUserId() string {
	return r._userId
}
