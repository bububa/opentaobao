package xhotelitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelRateRelationshipwithrpGetAPIRequest 根据gid查询卖家下所有的rpId API请求
// taobao.xhotel.rate.relationshipwithrp.get
//
// 根据gid查询卖家下所有的rpId，可分页，默认展示第一页的数据
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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelRateRelationshipwithrpGetAPIRequest) Reset() {
	r._gid = 0
	r._pageNo = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelRateRelationshipwithrpGetAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.rate.relationshipwithrp.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelRateRelationshipwithrpGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelRateRelationshipwithrpGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGid is Gid Setter
// 宝贝的gid
func (r *TaobaoXhotelRateRelationshipwithrpGetAPIRequest) SetGid(_gid int64) error {
	r._gid = _gid
	r.Set("gid", _gid)
	return nil
}

// GetGid Gid Getter
func (r TaobaoXhotelRateRelationshipwithrpGetAPIRequest) GetGid() int64 {
	return r._gid
}

// SetPageNo is PageNo Setter
// 页数，可根据此值展示某页的数据。不填默认为1
func (r *TaobaoXhotelRateRelationshipwithrpGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoXhotelRateRelationshipwithrpGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

var poolTaobaoXhotelRateRelationshipwithrpGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelRateRelationshipwithrpGetRequest()
	},
}

// GetTaobaoXhotelRateRelationshipwithrpGetRequest 从 sync.Pool 获取 TaobaoXhotelRateRelationshipwithrpGetAPIRequest
func GetTaobaoXhotelRateRelationshipwithrpGetAPIRequest() *TaobaoXhotelRateRelationshipwithrpGetAPIRequest {
	return poolTaobaoXhotelRateRelationshipwithrpGetAPIRequest.Get().(*TaobaoXhotelRateRelationshipwithrpGetAPIRequest)
}

// ReleaseTaobaoXhotelRateRelationshipwithrpGetAPIRequest 将 TaobaoXhotelRateRelationshipwithrpGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelRateRelationshipwithrpGetAPIRequest(v *TaobaoXhotelRateRelationshipwithrpGetAPIRequest) {
	v.Reset()
	poolTaobaoXhotelRateRelationshipwithrpGetAPIRequest.Put(v)
}
