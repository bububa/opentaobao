package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbaserchcrowdpricebatchupdateAPIRequest 单品推广搜索人群修改溢价 API请求
// taobao.simba.serchcrowd.price.batch.update
//
// 单品推广搜索人群修改溢价, 不支持跨推广单元修改
type TaobaosimbaserchcrowdpricebatchupdateAPIRequest struct {
	model.Params
	// 需要修改出价的人群包id,批量传入的时候用,分割
	_adgroupCrowdIds []string
	// 子帐号nick
	_subNick string
	// 被操作者的淘宝昵称
	_nick string
	// 推广单元id
	_adgroupId int64
	// 人群溢价比例，溢价范围[5,300]
	_discount int64
}

// NewTaobaosimbaserchcrowdpricebatchupdateRequest 初始化TaobaosimbaserchcrowdpricebatchupdateAPIRequest对象
func NewTaobaosimbaserchcrowdpricebatchupdateRequest() *TaobaosimbaserchcrowdpricebatchupdateAPIRequest {
	return &TaobaosimbaserchcrowdpricebatchupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbaserchcrowdpricebatchupdateAPIRequest) GetApiMethodName() string {
	return "taobao.simba.serchcrowd.price.batch.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbaserchcrowdpricebatchupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbaserchcrowdpricebatchupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAdgroupCrowdIds is AdgroupCrowdIds Setter
// 需要修改出价的人群包id,批量传入的时候用,分割
func (r *TaobaosimbaserchcrowdpricebatchupdateAPIRequest) SetAdgroupCrowdIds(_adgroupCrowdIds []string) error {
	r._adgroupCrowdIds = _adgroupCrowdIds
	r.Set("adgroup_crowd_ids", _adgroupCrowdIds)
	return nil
}

// GetAdgroupCrowdIds AdgroupCrowdIds Getter
func (r TaobaosimbaserchcrowdpricebatchupdateAPIRequest) GetAdgroupCrowdIds() []string {
	return r._adgroupCrowdIds
}

// SetSubNick is SubNick Setter
// 子帐号nick
func (r *TaobaosimbaserchcrowdpricebatchupdateAPIRequest) SetSubNick(_subNick string) error {
	r._subNick = _subNick
	r.Set("sub_nick", _subNick)
	return nil
}

// GetSubNick SubNick Getter
func (r TaobaosimbaserchcrowdpricebatchupdateAPIRequest) GetSubNick() string {
	return r._subNick
}

// SetNick is Nick Setter
// 被操作者的淘宝昵称
func (r *TaobaosimbaserchcrowdpricebatchupdateAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbaserchcrowdpricebatchupdateAPIRequest) GetNick() string {
	return r._nick
}

// SetAdgroupId is AdgroupId Setter
// 推广单元id
func (r *TaobaosimbaserchcrowdpricebatchupdateAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaosimbaserchcrowdpricebatchupdateAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

// SetDiscount is Discount Setter
// 人群溢价比例，溢价范围[5,300]
func (r *TaobaosimbaserchcrowdpricebatchupdateAPIRequest) SetDiscount(_discount int64) error {
	r._discount = _discount
	r.Set("discount", _discount)
	return nil
}

// GetDiscount Discount Getter
func (r TaobaosimbaserchcrowdpricebatchupdateAPIRequest) GetDiscount() int64 {
	return r._discount
}
