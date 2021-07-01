package alihealthcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthBabyBaseinfoGetAPIRequest
三方从我们这获取宝宝信息 API请求
alibaba.alihealth.baby.baseinfo.get

三方从我们这获取宝宝信息 */
type AlibabaAlihealthBabyBaseinfoGetAPIRequest struct {
	model.Params
	// 宝宝id
	_babyId int64
	// 宝宝所属的用户
	_tpUserId int64
}

// NewAlibabaAlihealthBabyBaseinfoGetRequest 初始化AlibabaAlihealthBabyBaseinfoGetAPIRequest对象
func NewAlibabaAlihealthBabyBaseinfoGetRequest() *AlibabaAlihealthBabyBaseinfoGetAPIRequest {
	return &AlibabaAlihealthBabyBaseinfoGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthBabyBaseinfoGetAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.baby.baseinfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthBabyBaseinfoGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is BabyId Setter
// 宝宝id
func (r *AlibabaAlihealthBabyBaseinfoGetAPIRequest) SetBabyId(_babyId int64) error {
	r._babyId = _babyId
	r.Set("baby_id", _babyId)
	return nil
}

// Get BabyId Getter
func (r AlibabaAlihealthBabyBaseinfoGetAPIRequest) GetBabyId() int64 {
	return r._babyId
}

// Set is TpUserId Setter
// 宝宝所属的用户
func (r *AlibabaAlihealthBabyBaseinfoGetAPIRequest) SetTpUserId(_tpUserId int64) error {
	r._tpUserId = _tpUserId
	r.Set("tp_user_id", _tpUserId)
	return nil
}

// Get TpUserId Getter
func (r AlibabaAlihealthBabyBaseinfoGetAPIRequest) GetTpUserId() int64 {
	return r._tpUserId
}
