package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaSerchcrowdGetAPIRequest 根据推广单元id获取搜索溢价人群 API请求
// taobao.simba.serchcrowd.get
//
// 根据推广单元id获取搜索溢价人群
type TaobaoSimbaSerchcrowdGetAPIRequest struct {
	model.Params
	// 被操作者的淘宝昵称
	_nick string
	// 推广单元id
	_adgroupId int64
}

// NewTaobaoSimbaSerchcrowdGetRequest 初始化TaobaoSimbaSerchcrowdGetAPIRequest对象
func NewTaobaoSimbaSerchcrowdGetRequest() *TaobaoSimbaSerchcrowdGetAPIRequest {
	return &TaobaoSimbaSerchcrowdGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaSerchcrowdGetAPIRequest) Reset() {
	r._nick = ""
	r._adgroupId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaSerchcrowdGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.serchcrowd.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaSerchcrowdGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaSerchcrowdGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 被操作者的淘宝昵称
func (r *TaobaoSimbaSerchcrowdGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaSerchcrowdGetAPIRequest) GetNick() string {
	return r._nick
}

// SetAdgroupId is AdgroupId Setter
// 推广单元id
func (r *TaobaoSimbaSerchcrowdGetAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaoSimbaSerchcrowdGetAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

var poolTaobaoSimbaSerchcrowdGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaSerchcrowdGetRequest()
	},
}

// GetTaobaoSimbaSerchcrowdGetRequest 从 sync.Pool 获取 TaobaoSimbaSerchcrowdGetAPIRequest
func GetTaobaoSimbaSerchcrowdGetAPIRequest() *TaobaoSimbaSerchcrowdGetAPIRequest {
	return poolTaobaoSimbaSerchcrowdGetAPIRequest.Get().(*TaobaoSimbaSerchcrowdGetAPIRequest)
}

// ReleaseTaobaoSimbaSerchcrowdGetAPIRequest 将 TaobaoSimbaSerchcrowdGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaSerchcrowdGetAPIRequest(v *TaobaoSimbaSerchcrowdGetAPIRequest) {
	v.Reset()
	poolTaobaoSimbaSerchcrowdGetAPIRequest.Put(v)
}
