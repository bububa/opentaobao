package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaKeywordsQscoreGetAPIRequest 取得一个推广组的所有关键词的质量得分或者根据关键词Id列表取得一组关键词的质量得分 API请求
// taobao.simba.keywords.qscore.get
//
// 取得一个推广组的所有关键词的质量得分列表
type TaobaoSimbaKeywordsQscoreGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广组Id
	_adgroupId int64
}

// NewTaobaoSimbaKeywordsQscoreGetRequest 初始化TaobaoSimbaKeywordsQscoreGetAPIRequest对象
func NewTaobaoSimbaKeywordsQscoreGetRequest() *TaobaoSimbaKeywordsQscoreGetAPIRequest {
	return &TaobaoSimbaKeywordsQscoreGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaKeywordsQscoreGetAPIRequest) Reset() {
	r._nick = ""
	r._adgroupId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaKeywordsQscoreGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.keywords.qscore.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaKeywordsQscoreGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaKeywordsQscoreGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaoSimbaKeywordsQscoreGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaKeywordsQscoreGetAPIRequest) GetNick() string {
	return r._nick
}

// SetAdgroupId is AdgroupId Setter
// 推广组Id
func (r *TaobaoSimbaKeywordsQscoreGetAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaoSimbaKeywordsQscoreGetAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

var poolTaobaoSimbaKeywordsQscoreGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaKeywordsQscoreGetRequest()
	},
}

// GetTaobaoSimbaKeywordsQscoreGetRequest 从 sync.Pool 获取 TaobaoSimbaKeywordsQscoreGetAPIRequest
func GetTaobaoSimbaKeywordsQscoreGetAPIRequest() *TaobaoSimbaKeywordsQscoreGetAPIRequest {
	return poolTaobaoSimbaKeywordsQscoreGetAPIRequest.Get().(*TaobaoSimbaKeywordsQscoreGetAPIRequest)
}

// ReleaseTaobaoSimbaKeywordsQscoreGetAPIRequest 将 TaobaoSimbaKeywordsQscoreGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaKeywordsQscoreGetAPIRequest(v *TaobaoSimbaKeywordsQscoreGetAPIRequest) {
	v.Reset()
	poolTaobaoSimbaKeywordsQscoreGetAPIRequest.Put(v)
}
