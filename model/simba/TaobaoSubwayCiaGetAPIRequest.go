package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubwayCiaGetAPIRequest 查询单元智能出价信息 API请求
// taobao.subway.cia.get
//
// 查询单元智能出价信息
type TaobaoSubwayCiaGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广组Id
	_adgroupId int64
}

// NewTaobaoSubwayCiaGetRequest 初始化TaobaoSubwayCiaGetAPIRequest对象
func NewTaobaoSubwayCiaGetRequest() *TaobaoSubwayCiaGetAPIRequest {
	return &TaobaoSubwayCiaGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSubwayCiaGetAPIRequest) Reset() {
	r._nick = ""
	r._adgroupId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSubwayCiaGetAPIRequest) GetApiMethodName() string {
	return "taobao.subway.cia.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSubwayCiaGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSubwayCiaGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaoSubwayCiaGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSubwayCiaGetAPIRequest) GetNick() string {
	return r._nick
}

// SetAdgroupId is AdgroupId Setter
// 推广组Id
func (r *TaobaoSubwayCiaGetAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaoSubwayCiaGetAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

var poolTaobaoSubwayCiaGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSubwayCiaGetRequest()
	},
}

// GetTaobaoSubwayCiaGetRequest 从 sync.Pool 获取 TaobaoSubwayCiaGetAPIRequest
func GetTaobaoSubwayCiaGetAPIRequest() *TaobaoSubwayCiaGetAPIRequest {
	return poolTaobaoSubwayCiaGetAPIRequest.Get().(*TaobaoSubwayCiaGetAPIRequest)
}

// ReleaseTaobaoSubwayCiaGetAPIRequest 将 TaobaoSubwayCiaGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSubwayCiaGetAPIRequest(v *TaobaoSubwayCiaGetAPIRequest) {
	v.Reset()
	poolTaobaoSubwayCiaGetAPIRequest.Put(v)
}
