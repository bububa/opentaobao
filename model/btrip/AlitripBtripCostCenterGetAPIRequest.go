package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripcostcentergetAPIRequest 获取用户费用归属 API请求
// alitrip.btrip.cost.center.get
//
// 获取差旅申请用户的费用归属列表
type AlitripbtripcostcentergetAPIRequest struct {
	model.Params
	// 企业id
	_corpId string
	// 用户id
	_userId string
}

// NewAlitripbtripcostcentergetRequest 初始化AlitripbtripcostcentergetAPIRequest对象
func NewAlitripbtripcostcentergetRequest() *AlitripbtripcostcentergetAPIRequest {
	return &AlitripbtripcostcentergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripcostcentergetAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.cost.center.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripcostcentergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripcostcentergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCorpId is CorpId Setter
// 企业id
func (r *AlitripbtripcostcentergetAPIRequest) SetCorpId(_corpId string) error {
	r._corpId = _corpId
	r.Set("corp_id", _corpId)
	return nil
}

// GetCorpId CorpId Getter
func (r AlitripbtripcostcentergetAPIRequest) GetCorpId() string {
	return r._corpId
}

// SetUserId is UserId Setter
// 用户id
func (r *AlitripbtripcostcentergetAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlitripbtripcostcentergetAPIRequest) GetUserId() string {
	return r._userId
}
