package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkscpublisherinfogetAPIRequest 淘宝客-公用-私域用户备案信息查询 API请求
// taobao.tbk.sc.publisher.info.get
//
// 查询已生成的渠道id或会员运营id的相关信息。
type TaobaotbkscpublisherinfogetAPIRequest struct {
	model.Params
	// 备案的场景：common（通用备案），etao（一淘备案），minietao（一淘小程序备案），offlineShop（线下门店备案），offlinePerson（线下个人备案）。如不填默认common。查询会员信息只需填写common即可
	_relationapp string
	// 会员独占 - 会员运营ID
	_specialid string
	// 淘宝客外部用户标记，如自身系统账户ID；微信ID等
	_externalid string
	// 渠道独占 - 渠道关系ID
	_relationid int64
	// 第几页，下标从0开始
	_pageno int64
	// 每页大小
	_pagesize int64
	// 类型，必选 1:渠道信息；2:会员信息
	_infotype int64
	// 1-微信、2-微博、3-抖音、4-快手、5-QQ，0-其他；默认为0
	_externaltype int64
}

// NewTaobaotbkscpublisherinfogetRequest 初始化TaobaotbkscpublisherinfogetAPIRequest对象
func NewTaobaotbkscpublisherinfogetRequest() *TaobaotbkscpublisherinfogetAPIRequest {
	return &TaobaotbkscpublisherinfogetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkscpublisherinfogetAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.publisher.info.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkscpublisherinfogetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkscpublisherinfogetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRelationapp is Relationapp Setter
// 备案的场景：common（通用备案），etao（一淘备案），minietao（一淘小程序备案），offlineShop（线下门店备案），offlinePerson（线下个人备案）。如不填默认common。查询会员信息只需填写common即可
func (r *TaobaotbkscpublisherinfogetAPIRequest) SetRelationapp(_relationapp string) error {
	r._relationapp = _relationapp
	r.Set("relation_app", _relationapp)
	return nil
}

// GetRelationapp Relationapp Getter
func (r TaobaotbkscpublisherinfogetAPIRequest) GetRelationapp() string {
	return r._relationapp
}

// SetSpecialid is Specialid Setter
// 会员独占 - 会员运营ID
func (r *TaobaotbkscpublisherinfogetAPIRequest) SetSpecialid(_specialid string) error {
	r._specialid = _specialid
	r.Set("special_id", _specialid)
	return nil
}

// GetSpecialid Specialid Getter
func (r TaobaotbkscpublisherinfogetAPIRequest) GetSpecialid() string {
	return r._specialid
}

// SetExternalid is Externalid Setter
// 淘宝客外部用户标记，如自身系统账户ID；微信ID等
func (r *TaobaotbkscpublisherinfogetAPIRequest) SetExternalid(_externalid string) error {
	r._externalid = _externalid
	r.Set("external_id", _externalid)
	return nil
}

// GetExternalid Externalid Getter
func (r TaobaotbkscpublisherinfogetAPIRequest) GetExternalid() string {
	return r._externalid
}

// SetRelationid is Relationid Setter
// 渠道独占 - 渠道关系ID
func (r *TaobaotbkscpublisherinfogetAPIRequest) SetRelationid(_relationid int64) error {
	r._relationid = _relationid
	r.Set("relation_id", _relationid)
	return nil
}

// GetRelationid Relationid Getter
func (r TaobaotbkscpublisherinfogetAPIRequest) GetRelationid() int64 {
	return r._relationid
}

// SetPageno is Pageno Setter
// 第几页，下标从0开始
func (r *TaobaotbkscpublisherinfogetAPIRequest) SetPageno(_pageno int64) error {
	r._pageno = _pageno
	r.Set("page_no", _pageno)
	return nil
}

// GetPageno Pageno Getter
func (r TaobaotbkscpublisherinfogetAPIRequest) GetPageno() int64 {
	return r._pageno
}

// SetPagesize is Pagesize Setter
// 每页大小
func (r *TaobaotbkscpublisherinfogetAPIRequest) SetPagesize(_pagesize int64) error {
	r._pagesize = _pagesize
	r.Set("page_size", _pagesize)
	return nil
}

// GetPagesize Pagesize Getter
func (r TaobaotbkscpublisherinfogetAPIRequest) GetPagesize() int64 {
	return r._pagesize
}

// SetInfotype is Infotype Setter
// 类型，必选 1:渠道信息；2:会员信息
func (r *TaobaotbkscpublisherinfogetAPIRequest) SetInfotype(_infotype int64) error {
	r._infotype = _infotype
	r.Set("info_type", _infotype)
	return nil
}

// GetInfotype Infotype Getter
func (r TaobaotbkscpublisherinfogetAPIRequest) GetInfotype() int64 {
	return r._infotype
}

// SetExternaltype is Externaltype Setter
// 1-微信、2-微博、3-抖音、4-快手、5-QQ，0-其他；默认为0
func (r *TaobaotbkscpublisherinfogetAPIRequest) SetExternaltype(_externaltype int64) error {
	r._externaltype = _externaltype
	r.Set("external_type", _externaltype)
	return nil
}

// GetExternaltype Externaltype Getter
func (r TaobaotbkscpublisherinfogetAPIRequest) GetExternaltype() int64 {
	return r._externaltype
}
