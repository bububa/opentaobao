package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosArkHealthOfflineLevelGetAPIRequest 获取mall的离线等级 API请求
// alibaba.mos.ark.health.offline.level.get
//
// 获取mall的离线等级
type AlibabaMosArkHealthOfflineLevelGetAPIRequest struct {
	model.Params
	// 商场id
	_mallId string
}

// NewAlibabaMosArkHealthOfflineLevelGetRequest 初始化AlibabaMosArkHealthOfflineLevelGetAPIRequest对象
func NewAlibabaMosArkHealthOfflineLevelGetRequest() *AlibabaMosArkHealthOfflineLevelGetAPIRequest {
	return &AlibabaMosArkHealthOfflineLevelGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMosArkHealthOfflineLevelGetAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.ark.health.offline.level.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMosArkHealthOfflineLevelGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetMallId is MallId Setter
// 商场id
func (r *AlibabaMosArkHealthOfflineLevelGetAPIRequest) SetMallId(_mallId string) error {
	r._mallId = _mallId
	r.Set("mall_id", _mallId)
	return nil
}

// GetMallId MallId Getter
func (r AlibabaMosArkHealthOfflineLevelGetAPIRequest) GetMallId() string {
	return r._mallId
}
