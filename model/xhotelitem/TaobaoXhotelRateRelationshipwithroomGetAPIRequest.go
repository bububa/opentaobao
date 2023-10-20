package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelRateRelationshipwithroomGetAPIRequest 查询rpId API请求
// taobao.xhotel.rate.relationshipwithroom.get
//
// 某个卖家根据rpId查询所有的gid，可分页，不填分页信息则默认显示第一页。
type TaobaoXhotelRateRelationshipwithroomGetAPIRequest struct {
	model.Params
	// rpId
	_rpId int64
	// 页数
	_pageNo int64
}

// NewTaobaoXhotelRateRelationshipwithroomGetRequest 初始化TaobaoXhotelRateRelationshipwithroomGetAPIRequest对象
func NewTaobaoXhotelRateRelationshipwithroomGetRequest() *TaobaoXhotelRateRelationshipwithroomGetAPIRequest {
	return &TaobaoXhotelRateRelationshipwithroomGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelRateRelationshipwithroomGetAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.rate.relationshipwithroom.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelRateRelationshipwithroomGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelRateRelationshipwithroomGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRpId is RpId Setter
// rpId
func (r *TaobaoXhotelRateRelationshipwithroomGetAPIRequest) SetRpId(_rpId int64) error {
	r._rpId = _rpId
	r.Set("rp_id", _rpId)
	return nil
}

// GetRpId RpId Getter
func (r TaobaoXhotelRateRelationshipwithroomGetAPIRequest) GetRpId() int64 {
	return r._rpId
}

// SetPageNo is PageNo Setter
// 页数
func (r *TaobaoXhotelRateRelationshipwithroomGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoXhotelRateRelationshipwithroomGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}
