package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpentradeGroupDetailAPIRequest 组团购场景查询团详情 API请求
// taobao.opentrade.group.detail
//
// 组团购场景下，查询团详情
type TaobaoOpentradeGroupDetailAPIRequest struct {
	model.Params
	// 团id
	_groupId int64
}

// NewTaobaoOpentradeGroupDetailRequest 初始化TaobaoOpentradeGroupDetailAPIRequest对象
func NewTaobaoOpentradeGroupDetailRequest() *TaobaoOpentradeGroupDetailAPIRequest {
	return &TaobaoOpentradeGroupDetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeGroupDetailAPIRequest) GetApiMethodName() string {
	return "taobao.opentrade.group.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeGroupDetailAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is GroupId Setter
// 团id
func (r *TaobaoOpentradeGroupDetailAPIRequest) SetGroupId(_groupId int64) error {
	r._groupId = _groupId
	r.Set("group_id", _groupId)
	return nil
}

// Get GroupId Getter
func (r TaobaoOpentradeGroupDetailAPIRequest) GetGroupId() int64 {
	return r._groupId
}
