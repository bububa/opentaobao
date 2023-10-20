package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbakeywordsbyadgroupidgetAPIRequest 取得一个推广组的所有关键词 API请求
// taobao.simba.keywordsbyadgroupid.get
//
// 取得一个推广组的所有关键词
type TaobaosimbakeywordsbyadgroupidgetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广组Id
	_adgroupId int64
}

// NewTaobaosimbakeywordsbyadgroupidgetRequest 初始化TaobaosimbakeywordsbyadgroupidgetAPIRequest对象
func NewTaobaosimbakeywordsbyadgroupidgetRequest() *TaobaosimbakeywordsbyadgroupidgetAPIRequest {
	return &TaobaosimbakeywordsbyadgroupidgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbakeywordsbyadgroupidgetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.keywordsbyadgroupid.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbakeywordsbyadgroupidgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbakeywordsbyadgroupidgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaosimbakeywordsbyadgroupidgetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbakeywordsbyadgroupidgetAPIRequest) GetNick() string {
	return r._nick
}

// SetAdgroupId is AdgroupId Setter
// 推广组Id
func (r *TaobaosimbakeywordsbyadgroupidgetAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaosimbakeywordsbyadgroupidgetAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}
