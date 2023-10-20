package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelraterelationshipwithrpgetAPIRequest 根据gid查询卖家下所有的rpId API请求
// taobao.xhotel.rate.relationshipwithrp.get
//
// 根据gid查询卖家下所有的rpId，可分页，默认展示第一页的数据
type TaobaoxhotelraterelationshipwithrpgetAPIRequest struct {
	model.Params
	// 宝贝的gid
	_gid int64
	// 页数，可根据此值展示某页的数据。不填默认为1
	_pageNo int64
}

// NewTaobaoxhotelraterelationshipwithrpgetRequest 初始化TaobaoxhotelraterelationshipwithrpgetAPIRequest对象
func NewTaobaoxhotelraterelationshipwithrpgetRequest() *TaobaoxhotelraterelationshipwithrpgetAPIRequest {
	return &TaobaoxhotelraterelationshipwithrpgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelraterelationshipwithrpgetAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.rate.relationshipwithrp.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelraterelationshipwithrpgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelraterelationshipwithrpgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGid is Gid Setter
// 宝贝的gid
func (r *TaobaoxhotelraterelationshipwithrpgetAPIRequest) SetGid(_gid int64) error {
	r._gid = _gid
	r.Set("gid", _gid)
	return nil
}

// GetGid Gid Getter
func (r TaobaoxhotelraterelationshipwithrpgetAPIRequest) GetGid() int64 {
	return r._gid
}

// SetPageNo is PageNo Setter
// 页数，可根据此值展示某页的数据。不填默认为1
func (r *TaobaoxhotelraterelationshipwithrpgetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoxhotelraterelationshipwithrpgetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}
