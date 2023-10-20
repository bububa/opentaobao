package baodian

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaodegusergamegiftqueryAPIRequest 用户数娱游戏礼包查询 API请求
// taobao.deg.user.gamegift.query
//
// 查询用户数娱礼包列表
type TaobaodegusergamegiftqueryAPIRequest struct {
	model.Params
	// cp item id列表
	_cpItemIds []string
	// 状态，1为待发放，2为已发放，3为过期
	_status int64
}

// NewTaobaodegusergamegiftqueryRequest 初始化TaobaodegusergamegiftqueryAPIRequest对象
func NewTaobaodegusergamegiftqueryRequest() *TaobaodegusergamegiftqueryAPIRequest {
	return &TaobaodegusergamegiftqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaodegusergamegiftqueryAPIRequest) GetApiMethodName() string {
	return "taobao.deg.user.gamegift.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaodegusergamegiftqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaodegusergamegiftqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCpItemIds is CpItemIds Setter
// cp item id列表
func (r *TaobaodegusergamegiftqueryAPIRequest) SetCpItemIds(_cpItemIds []string) error {
	r._cpItemIds = _cpItemIds
	r.Set("cp_item_ids", _cpItemIds)
	return nil
}

// GetCpItemIds CpItemIds Getter
func (r TaobaodegusergamegiftqueryAPIRequest) GetCpItemIds() []string {
	return r._cpItemIds
}

// SetStatus is Status Setter
// 状态，1为待发放，2为已发放，3为过期
func (r *TaobaodegusergamegiftqueryAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaodegusergamegiftqueryAPIRequest) GetStatus() int64 {
	return r._status
}
