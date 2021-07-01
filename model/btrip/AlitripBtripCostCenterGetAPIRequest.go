package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripCostCenterGetAPIRequest
获取用户费用归属 API请求
alitrip.btrip.cost.center.get

获取差旅申请用户的费用归属列表 */
type AlitripBtripCostCenterGetAPIRequest struct {
	model.Params
	// 企业id
	_corpId string
	// 用户id
	_userId string
}

// NewAlitripBtripCostCenterGetRequest 初始化AlitripBtripCostCenterGetAPIRequest对象
func NewAlitripBtripCostCenterGetRequest() *AlitripBtripCostCenterGetAPIRequest {
	return &AlitripBtripCostCenterGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripCostCenterGetAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.cost.center.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripCostCenterGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CorpId Setter
// 企业id
func (r *AlitripBtripCostCenterGetAPIRequest) SetCorpId(_corpId string) error {
	r._corpId = _corpId
	r.Set("corp_id", _corpId)
	return nil
}

// Get CorpId Getter
func (r AlitripBtripCostCenterGetAPIRequest) GetCorpId() string {
	return r._corpId
}

// Set is UserId Setter
// 用户id
func (r *AlitripBtripCostCenterGetAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// Get UserId Getter
func (r AlitripBtripCostCenterGetAPIRequest) GetUserId() string {
	return r._userId
}
