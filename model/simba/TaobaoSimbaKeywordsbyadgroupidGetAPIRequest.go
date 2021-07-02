package simba

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaKeywordsbyadgroupidGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.keywordsbyadgroupid.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaKeywordsbyadgroupidGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
