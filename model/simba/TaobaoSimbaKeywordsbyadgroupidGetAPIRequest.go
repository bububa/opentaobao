package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaKeywordsbyadgroupidGetAPIRequest 取得一个推广组的所有关键词 API请求
// taobao.simba.keywordsbyadgroupid.get
//
// 取得一个推广组的所有关键词
type TaobaoSimbaKeywordsbyadgroupidGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广组Id
	_adgroupId int64
}

// NewTaobaoSimbaKeywordsbyadgroupidGetRequest 初始化TaobaoSimbaKeywordsbyadgroupidGetAPIRequest对象
func NewTaobaoSimbaKeywordsbyadgroupidGetRequest() *TaobaoSimbaKeywordsbyadgroupidGetAPIRequest {
	return &TaobaoSimbaKeywordsbyadgroupidGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaKeywordsbyadgroupidGetAPIRequest) Reset() {
	r._nick = ""
	r._adgroupId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaKeywordsbyadgroupidGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.keywordsbyadgroupid.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaKeywordsbyadgroupidGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaKeywordsbyadgroupidGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaoSimbaKeywordsbyadgroupidGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaKeywordsbyadgroupidGetAPIRequest) GetNick() string {
	return r._nick
}

// SetAdgroupId is AdgroupId Setter
// 推广组Id
func (r *TaobaoSimbaKeywordsbyadgroupidGetAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaoSimbaKeywordsbyadgroupidGetAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

var poolTaobaoSimbaKeywordsbyadgroupidGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaKeywordsbyadgroupidGetRequest()
	},
}

// GetTaobaoSimbaKeywordsbyadgroupidGetRequest 从 sync.Pool 获取 TaobaoSimbaKeywordsbyadgroupidGetAPIRequest
func GetTaobaoSimbaKeywordsbyadgroupidGetAPIRequest() *TaobaoSimbaKeywordsbyadgroupidGetAPIRequest {
	return poolTaobaoSimbaKeywordsbyadgroupidGetAPIRequest.Get().(*TaobaoSimbaKeywordsbyadgroupidGetAPIRequest)
}

// ReleaseTaobaoSimbaKeywordsbyadgroupidGetAPIRequest 将 TaobaoSimbaKeywordsbyadgroupidGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaKeywordsbyadgroupidGetAPIRequest(v *TaobaoSimbaKeywordsbyadgroupidGetAPIRequest) {
	v.Reset()
	poolTaobaoSimbaKeywordsbyadgroupidGetAPIRequest.Put(v)
}
