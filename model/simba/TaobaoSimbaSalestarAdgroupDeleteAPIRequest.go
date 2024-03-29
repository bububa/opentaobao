package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaSalestarAdgroupDeleteAPIRequest (新)销量明星删除推广单元接口 API请求
// taobao.simba.salestar.adgroup.delete
//
// 删除一个推广组
type TaobaoSimbaSalestarAdgroupDeleteAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广组Id
	_adgroupId int64
}

// NewTaobaoSimbaSalestarAdgroupDeleteRequest 初始化TaobaoSimbaSalestarAdgroupDeleteAPIRequest对象
func NewTaobaoSimbaSalestarAdgroupDeleteRequest() *TaobaoSimbaSalestarAdgroupDeleteAPIRequest {
	return &TaobaoSimbaSalestarAdgroupDeleteAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaSalestarAdgroupDeleteAPIRequest) Reset() {
	r._nick = ""
	r._adgroupId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaSalestarAdgroupDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.simba.salestar.adgroup.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaSalestarAdgroupDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaSalestarAdgroupDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaoSimbaSalestarAdgroupDeleteAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaSalestarAdgroupDeleteAPIRequest) GetNick() string {
	return r._nick
}

// SetAdgroupId is AdgroupId Setter
// 推广组Id
func (r *TaobaoSimbaSalestarAdgroupDeleteAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaoSimbaSalestarAdgroupDeleteAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

var poolTaobaoSimbaSalestarAdgroupDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaSalestarAdgroupDeleteRequest()
	},
}

// GetTaobaoSimbaSalestarAdgroupDeleteRequest 从 sync.Pool 获取 TaobaoSimbaSalestarAdgroupDeleteAPIRequest
func GetTaobaoSimbaSalestarAdgroupDeleteAPIRequest() *TaobaoSimbaSalestarAdgroupDeleteAPIRequest {
	return poolTaobaoSimbaSalestarAdgroupDeleteAPIRequest.Get().(*TaobaoSimbaSalestarAdgroupDeleteAPIRequest)
}

// ReleaseTaobaoSimbaSalestarAdgroupDeleteAPIRequest 将 TaobaoSimbaSalestarAdgroupDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaSalestarAdgroupDeleteAPIRequest(v *TaobaoSimbaSalestarAdgroupDeleteAPIRequest) {
	v.Reset()
	poolTaobaoSimbaSalestarAdgroupDeleteAPIRequest.Put(v)
}
