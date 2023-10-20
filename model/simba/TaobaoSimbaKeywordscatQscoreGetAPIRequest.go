package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaKeywordscatQscoreGetAPIRequest 取得一个推广组的所有关键词和类目出价的质量得分 API请求
// taobao.simba.keywordscat.qscore.get
//
// 取得一个推广组的所有关键词和类目出价的质量得分列表
type TaobaoSimbaKeywordscatQscoreGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广组Id
	_adgroupId int64
}

// NewTaobaoSimbaKeywordscatQscoreGetRequest 初始化TaobaoSimbaKeywordscatQscoreGetAPIRequest对象
func NewTaobaoSimbaKeywordscatQscoreGetRequest() *TaobaoSimbaKeywordscatQscoreGetAPIRequest {
	return &TaobaoSimbaKeywordscatQscoreGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaKeywordscatQscoreGetAPIRequest) Reset() {
	r._nick = ""
	r._adgroupId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaKeywordscatQscoreGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.keywordscat.qscore.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaKeywordscatQscoreGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaKeywordscatQscoreGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaoSimbaKeywordscatQscoreGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaKeywordscatQscoreGetAPIRequest) GetNick() string {
	return r._nick
}

// SetAdgroupId is AdgroupId Setter
// 推广组Id
func (r *TaobaoSimbaKeywordscatQscoreGetAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaoSimbaKeywordscatQscoreGetAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

var poolTaobaoSimbaKeywordscatQscoreGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaKeywordscatQscoreGetRequest()
	},
}

// GetTaobaoSimbaKeywordscatQscoreGetRequest 从 sync.Pool 获取 TaobaoSimbaKeywordscatQscoreGetAPIRequest
func GetTaobaoSimbaKeywordscatQscoreGetAPIRequest() *TaobaoSimbaKeywordscatQscoreGetAPIRequest {
	return poolTaobaoSimbaKeywordscatQscoreGetAPIRequest.Get().(*TaobaoSimbaKeywordscatQscoreGetAPIRequest)
}

// ReleaseTaobaoSimbaKeywordscatQscoreGetAPIRequest 将 TaobaoSimbaKeywordscatQscoreGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaKeywordscatQscoreGetAPIRequest(v *TaobaoSimbaKeywordscatQscoreGetAPIRequest) {
	v.Reset()
	poolTaobaoSimbaKeywordscatQscoreGetAPIRequest.Put(v)
}
