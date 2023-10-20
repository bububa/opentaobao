package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbaadgroupsbyadgroupidsgetAPIRequest 批量得到推广组 API请求
// taobao.simba.adgroupsbyadgroupids.get
//
// 批量得到推广组
type TaobaosimbaadgroupsbyadgroupidsgetAPIRequest struct {
	model.Params
	// 推广组Id列表
	_adgroupIds []string
	// 主人昵称
	_nick string
	// 页尺寸，最大200，如果入参adgroup_ids有传入值，则page_size和page_no值不起作用。如果adgrpup_ids为空而campaign_id有值，此时page_size和page_no值才是返回的页数据大小和页码
	_pageSize int64
	// 页码，从1开始
	_pageNo int64
}

// NewTaobaosimbaadgroupsbyadgroupidsgetRequest 初始化TaobaosimbaadgroupsbyadgroupidsgetAPIRequest对象
func NewTaobaosimbaadgroupsbyadgroupidsgetRequest() *TaobaosimbaadgroupsbyadgroupidsgetAPIRequest {
	return &TaobaosimbaadgroupsbyadgroupidsgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbaadgroupsbyadgroupidsgetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.adgroupsbyadgroupids.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbaadgroupsbyadgroupidsgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbaadgroupsbyadgroupidsgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAdgroupIds is AdgroupIds Setter
// 推广组Id列表
func (r *TaobaosimbaadgroupsbyadgroupidsgetAPIRequest) SetAdgroupIds(_adgroupIds []string) error {
	r._adgroupIds = _adgroupIds
	r.Set("adgroup_ids", _adgroupIds)
	return nil
}

// GetAdgroupIds AdgroupIds Getter
func (r TaobaosimbaadgroupsbyadgroupidsgetAPIRequest) GetAdgroupIds() []string {
	return r._adgroupIds
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaosimbaadgroupsbyadgroupidsgetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbaadgroupsbyadgroupidsgetAPIRequest) GetNick() string {
	return r._nick
}

// SetPageSize is PageSize Setter
// 页尺寸，最大200，如果入参adgroup_ids有传入值，则page_size和page_no值不起作用。如果adgrpup_ids为空而campaign_id有值，此时page_size和page_no值才是返回的页数据大小和页码
func (r *TaobaosimbaadgroupsbyadgroupidsgetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaosimbaadgroupsbyadgroupidsgetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNo is PageNo Setter
// 页码，从1开始
func (r *TaobaosimbaadgroupsbyadgroupidsgetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaosimbaadgroupsbyadgroupidsgetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}
