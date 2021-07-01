package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUmpRangeGetAPIRequest
查询活动范围 API请求
taobao.ump.range.get

查询某个卖家所有参加或者不参加某项活动的物品 */
type TaobaoUmpRangeGetAPIRequest struct {
	model.Params
	// 活动id
	_actId int64
}

// NewTaobaoUmpRangeGetRequest 初始化TaobaoUmpRangeGetAPIRequest对象
func NewTaobaoUmpRangeGetRequest() *TaobaoUmpRangeGetAPIRequest {
	return &TaobaoUmpRangeGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUmpRangeGetAPIRequest) GetApiMethodName() string {
	return "taobao.ump.range.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUmpRangeGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ActId Setter
// 活动id
func (r *TaobaoUmpRangeGetAPIRequest) SetActId(_actId int64) error {
	r._actId = _actId
	r.Set("act_id", _actId)
	return nil
}

// Get ActId Getter
func (r TaobaoUmpRangeGetAPIRequest) GetActId() int64 {
	return r._actId
}
