package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaAdgroupDeleteAPIRequest 删除一个推广组 API请求
// taobao.simba.adgroup.delete
//
// 删除一个推广组
type TaobaoSimbaAdgroupDeleteAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广组Id
	_adgroupId int64
}

// NewTaobaoSimbaAdgroupDeleteRequest 初始化TaobaoSimbaAdgroupDeleteAPIRequest对象
func NewTaobaoSimbaAdgroupDeleteRequest() *TaobaoSimbaAdgroupDeleteAPIRequest {
	return &TaobaoSimbaAdgroupDeleteAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaAdgroupDeleteAPIRequest) Reset() {
	r._nick = ""
	r._adgroupId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaAdgroupDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.simba.adgroup.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaAdgroupDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaAdgroupDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaoSimbaAdgroupDeleteAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaAdgroupDeleteAPIRequest) GetNick() string {
	return r._nick
}

// SetAdgroupId is AdgroupId Setter
// 推广组Id
func (r *TaobaoSimbaAdgroupDeleteAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaoSimbaAdgroupDeleteAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

var poolTaobaoSimbaAdgroupDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaAdgroupDeleteRequest()
	},
}

// GetTaobaoSimbaAdgroupDeleteRequest 从 sync.Pool 获取 TaobaoSimbaAdgroupDeleteAPIRequest
func GetTaobaoSimbaAdgroupDeleteAPIRequest() *TaobaoSimbaAdgroupDeleteAPIRequest {
	return poolTaobaoSimbaAdgroupDeleteAPIRequest.Get().(*TaobaoSimbaAdgroupDeleteAPIRequest)
}

// ReleaseTaobaoSimbaAdgroupDeleteAPIRequest 将 TaobaoSimbaAdgroupDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaAdgroupDeleteAPIRequest(v *TaobaoSimbaAdgroupDeleteAPIRequest) {
	v.Reset()
	poolTaobaoSimbaAdgroupDeleteAPIRequest.Put(v)
}
