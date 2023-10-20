package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripapplygetAPIRequest 获取单个审批单 API请求
// alitrip.btrip.apply.get
//
// 获取单个审批单的详情数据
type AlitripbtripapplygetAPIRequest struct {
	model.Params
	// 外部审批单id
	_thirdpartApplyId string
	// 企业id
	_corpId string
	// 审批单展示id
	_applyShowId string
	// 阿里商旅审批单id
	_applyId int64
}

// NewAlitripbtripapplygetRequest 初始化AlitripbtripapplygetAPIRequest对象
func NewAlitripbtripapplygetRequest() *AlitripbtripapplygetAPIRequest {
	return &AlitripbtripapplygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripapplygetAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.apply.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripapplygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripapplygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetThirdpartApplyId is ThirdpartApplyId Setter
// 外部审批单id
func (r *AlitripbtripapplygetAPIRequest) SetThirdpartApplyId(_thirdpartApplyId string) error {
	r._thirdpartApplyId = _thirdpartApplyId
	r.Set("thirdpart_apply_id", _thirdpartApplyId)
	return nil
}

// GetThirdpartApplyId ThirdpartApplyId Getter
func (r AlitripbtripapplygetAPIRequest) GetThirdpartApplyId() string {
	return r._thirdpartApplyId
}

// SetCorpId is CorpId Setter
// 企业id
func (r *AlitripbtripapplygetAPIRequest) SetCorpId(_corpId string) error {
	r._corpId = _corpId
	r.Set("corp_id", _corpId)
	return nil
}

// GetCorpId CorpId Getter
func (r AlitripbtripapplygetAPIRequest) GetCorpId() string {
	return r._corpId
}

// SetApplyShowId is ApplyShowId Setter
// 审批单展示id
func (r *AlitripbtripapplygetAPIRequest) SetApplyShowId(_applyShowId string) error {
	r._applyShowId = _applyShowId
	r.Set("apply_show_id", _applyShowId)
	return nil
}

// GetApplyShowId ApplyShowId Getter
func (r AlitripbtripapplygetAPIRequest) GetApplyShowId() string {
	return r._applyShowId
}

// SetApplyId is ApplyId Setter
// 阿里商旅审批单id
func (r *AlitripbtripapplygetAPIRequest) SetApplyId(_applyId int64) error {
	r._applyId = _applyId
	r.Set("apply_id", _applyId)
	return nil
}

// GetApplyId ApplyId Getter
func (r AlitripbtripapplygetAPIRequest) GetApplyId() int64 {
	return r._applyId
}
