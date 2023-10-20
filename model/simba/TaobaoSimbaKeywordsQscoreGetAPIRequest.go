package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbakeywordsqscoregetAPIRequest 取得一个推广组的所有关键词的质量得分或者根据关键词Id列表取得一组关键词的质量得分 API请求
// taobao.simba.keywords.qscore.get
//
// 取得一个推广组的所有关键词的质量得分列表
type TaobaosimbakeywordsqscoregetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广组Id
	_adgroupId int64
}

// NewTaobaosimbakeywordsqscoregetRequest 初始化TaobaosimbakeywordsqscoregetAPIRequest对象
func NewTaobaosimbakeywordsqscoregetRequest() *TaobaosimbakeywordsqscoregetAPIRequest {
	return &TaobaosimbakeywordsqscoregetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbakeywordsqscoregetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.keywords.qscore.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbakeywordsqscoregetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbakeywordsqscoregetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaosimbakeywordsqscoregetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbakeywordsqscoregetAPIRequest) GetNick() string {
	return r._nick
}

// SetAdgroupId is AdgroupId Setter
// 推广组Id
func (r *TaobaosimbakeywordsqscoregetAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaosimbakeywordsqscoregetAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}
