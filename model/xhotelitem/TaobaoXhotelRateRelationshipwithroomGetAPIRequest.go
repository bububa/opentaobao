package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelraterelationshipwithroomgetAPIRequest 查询rpId API请求
// taobao.xhotel.rate.relationshipwithroom.get
//
// 某个卖家根据rpId查询所有的gid，可分页，不填分页信息则默认显示第一页。
type TaobaoxhotelraterelationshipwithroomgetAPIRequest struct {
	model.Params
	// rpId
	_rpId int64
	// 页数
	_pageNo int64
}

// NewTaobaoxhotelraterelationshipwithroomgetRequest 初始化TaobaoxhotelraterelationshipwithroomgetAPIRequest对象
func NewTaobaoxhotelraterelationshipwithroomgetRequest() *TaobaoxhotelraterelationshipwithroomgetAPIRequest {
	return &TaobaoxhotelraterelationshipwithroomgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelraterelationshipwithroomgetAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.rate.relationshipwithroom.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelraterelationshipwithroomgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelraterelationshipwithroomgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRpId is RpId Setter
// rpId
func (r *TaobaoxhotelraterelationshipwithroomgetAPIRequest) SetRpId(_rpId int64) error {
	r._rpId = _rpId
	r.Set("rp_id", _rpId)
	return nil
}

// GetRpId RpId Getter
func (r TaobaoxhotelraterelationshipwithroomgetAPIRequest) GetRpId() int64 {
	return r._rpId
}

// SetPageNo is PageNo Setter
// 页数
func (r *TaobaoxhotelraterelationshipwithroomgetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoxhotelraterelationshipwithroomgetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}
