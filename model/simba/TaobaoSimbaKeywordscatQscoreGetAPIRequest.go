package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbakeywordscatqscoregetAPIRequest 取得一个推广组的所有关键词和类目出价的质量得分 API请求
// taobao.simba.keywordscat.qscore.get
//
// 取得一个推广组的所有关键词和类目出价的质量得分列表
type TaobaosimbakeywordscatqscoregetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广组Id
	_adgroupId int64
}

// NewTaobaosimbakeywordscatqscoregetRequest 初始化TaobaosimbakeywordscatqscoregetAPIRequest对象
func NewTaobaosimbakeywordscatqscoregetRequest() *TaobaosimbakeywordscatqscoregetAPIRequest {
	return &TaobaosimbakeywordscatqscoregetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbakeywordscatqscoregetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.keywordscat.qscore.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbakeywordscatqscoregetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbakeywordscatqscoregetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaosimbakeywordscatqscoregetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbakeywordscatqscoregetAPIRequest) GetNick() string {
	return r._nick
}

// SetAdgroupId is AdgroupId Setter
// 推广组Id
func (r *TaobaosimbakeywordscatqscoregetAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaosimbakeywordscatqscoregetAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}
