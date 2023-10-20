package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpRangeGetAPIRequest 查询活动范围 API请求
// taobao.ump.range.get
//
// 查询某个卖家所有参加或者不参加某项活动的物品
type TaobaoUmpRangeGetAPIRequest struct {
	model.Params
	// 活动id
	_actId int64
}

// NewTaobaoUmpRangeGetRequest 初始化TaobaoUmpRangeGetAPIRequest对象
func NewTaobaoUmpRangeGetRequest() *TaobaoUmpRangeGetAPIRequest {
	return &TaobaoUmpRangeGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUmpRangeGetAPIRequest) Reset() {
	r._actId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUmpRangeGetAPIRequest) GetApiMethodName() string {
	return "taobao.ump.range.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUmpRangeGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUmpRangeGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActId is ActId Setter
// 活动id
func (r *TaobaoUmpRangeGetAPIRequest) SetActId(_actId int64) error {
	r._actId = _actId
	r.Set("act_id", _actId)
	return nil
}

// GetActId ActId Getter
func (r TaobaoUmpRangeGetAPIRequest) GetActId() int64 {
	return r._actId
}

var poolTaobaoUmpRangeGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUmpRangeGetRequest()
	},
}

// GetTaobaoUmpRangeGetRequest 从 sync.Pool 获取 TaobaoUmpRangeGetAPIRequest
func GetTaobaoUmpRangeGetAPIRequest() *TaobaoUmpRangeGetAPIRequest {
	return poolTaobaoUmpRangeGetAPIRequest.Get().(*TaobaoUmpRangeGetAPIRequest)
}

// ReleaseTaobaoUmpRangeGetAPIRequest 将 TaobaoUmpRangeGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoUmpRangeGetAPIRequest(v *TaobaoUmpRangeGetAPIRequest) {
	v.Reset()
	poolTaobaoUmpRangeGetAPIRequest.Put(v)
}
