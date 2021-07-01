package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelRateRelationshipwithrpGetAPIRequest
根据gid查询卖家下所有的rpId API请求
taobao.xhotel.rate.relationshipwithrp.get

根据gid查询卖家下所有的rpId，可分页，默认展示第一页的数据 */
type TaobaoXhotelRateRelationshipwithrpGetAPIRequest struct {
	model.Params
	// 宝贝的gid
	_gid int64
	// 页数，可根据此值展示某页的数据。不填默认为1
	_pageNo int64
}

// NewTaobaoXhotelRateRelationshipwithrpGetRequest 初始化TaobaoXhotelRateRelationshipwithrpGetAPIRequest对象
func NewTaobaoXhotelRateRelationshipwithrpGetRequest() *TaobaoXhotelRateRelationshipwithrpGetAPIRequest {
	return &TaobaoXhotelRateRelationshipwithrpGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelRateRelationshipwithrpGetAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.rate.relationshipwithrp.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelRateRelationshipwithrpGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Gid Setter
// 宝贝的gid
func (r *TaobaoXhotelRateRelationshipwithrpGetAPIRequest) SetGid(_gid int64) error {
	r._gid = _gid
	r.Set("gid", _gid)
	return nil
}

// Get Gid Getter
func (r TaobaoXhotelRateRelationshipwithrpGetAPIRequest) GetGid() int64 {
	return r._gid
}

// Set is PageNo Setter
// 页数，可根据此值展示某页的数据。不填默认为1
func (r *TaobaoXhotelRateRelationshipwithrpGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// Get PageNo Getter
func (r TaobaoXhotelRateRelationshipwithrpGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}
