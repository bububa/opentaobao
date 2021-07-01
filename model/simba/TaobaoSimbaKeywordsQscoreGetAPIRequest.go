package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaKeywordsQscoreGetAPIRequest
取得一个推广组的所有关键词的质量得分或者根据关键词Id列表取得一组关键词的质量得分 API请求
taobao.simba.keywords.qscore.get

取得一个推广组的所有关键词的质量得分列表 */
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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaKeywordsQscoreGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.keywords.qscore.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaKeywordsQscoreGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Nick Setter
// 主人昵称
func (r *TaobaoSimbaKeywordsQscoreGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// Get Nick Getter
func (r TaobaoSimbaKeywordsQscoreGetAPIRequest) GetNick() string {
	return r._nick
}

// Set is AdgroupId Setter
// 推广组Id
func (r *TaobaoSimbaKeywordsQscoreGetAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// Get AdgroupId Getter
func (r TaobaoSimbaKeywordsQscoreGetAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}
