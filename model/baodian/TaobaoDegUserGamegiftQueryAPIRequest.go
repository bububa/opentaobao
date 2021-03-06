package baodian

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDegUserGamegiftQueryAPIRequest 用户数娱游戏礼包查询 API请求
// taobao.deg.user.gamegift.query
//
// 查询用户数娱礼包列表
type TaobaoDegUserGamegiftQueryAPIRequest struct {
	model.Params
	// cp item id列表
	_cpItemIds []string
	// 状态，1为待发放，2为已发放，3为过期
	_status int64
}

// NewTaobaoDegUserGamegiftQueryRequest 初始化TaobaoDegUserGamegiftQueryAPIRequest对象
func NewTaobaoDegUserGamegiftQueryRequest() *TaobaoDegUserGamegiftQueryAPIRequest {
	return &TaobaoDegUserGamegiftQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoDegUserGamegiftQueryAPIRequest) GetApiMethodName() string {
	return "taobao.deg.user.gamegift.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoDegUserGamegiftQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCpItemIds is CpItemIds Setter
// cp item id列表
func (r *TaobaoDegUserGamegiftQueryAPIRequest) SetCpItemIds(_cpItemIds []string) error {
	r._cpItemIds = _cpItemIds
	r.Set("cp_item_ids", _cpItemIds)
	return nil
}

// GetCpItemIds CpItemIds Getter
func (r TaobaoDegUserGamegiftQueryAPIRequest) GetCpItemIds() []string {
	return r._cpItemIds
}

// SetStatus is Status Setter
// 状态，1为待发放，2为已发放，3为过期
func (r *TaobaoDegUserGamegiftQueryAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoDegUserGamegiftQueryAPIRequest) GetStatus() int64 {
	return r._status
}
