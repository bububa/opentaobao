package omniorder

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderGuideDataGetAPIRequest 获取全渠道导购产品数据 API请求
// taobao.omniorder.guide.data.get
//
// 获取全渠道导购产品，目前包括随心购、随身购扫码、加购和交易数据。
type TaobaoOmniorderGuideDataGetAPIRequest struct {
	model.Params
	// detail_smg_scan: 扫码购扫码明细;detail_smg_cart: 扫码购加购明细;detail_smg_order: 扫码购订单明细;detail_sxg_search: 随心购搜索明细;detail_sxg_view_item: 随心购商品浏览明细;detail_sxg_cart: 随心购加购明细;detail_sxg_order: 随心购订单明细
	_type string
	// 拉取数据开始时间
	_startTime string
	// 页码，从1开始
	_pageNo int64
	// 每页数量，不能大于1000
	_pageSize int64
}

// NewTaobaoOmniorderGuideDataGetRequest 初始化TaobaoOmniorderGuideDataGetAPIRequest对象
func NewTaobaoOmniorderGuideDataGetRequest() *TaobaoOmniorderGuideDataGetAPIRequest {
	return &TaobaoOmniorderGuideDataGetAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOmniorderGuideDataGetAPIRequest) Reset() {
	r._type = ""
	r._startTime = ""
	r._pageNo = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderGuideDataGetAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.guide.data.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderGuideDataGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOmniorderGuideDataGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetType is Type Setter
// detail_smg_scan: 扫码购扫码明细;detail_smg_cart: 扫码购加购明细;detail_smg_order: 扫码购订单明细;detail_sxg_search: 随心购搜索明细;detail_sxg_view_item: 随心购商品浏览明细;detail_sxg_cart: 随心购加购明细;detail_sxg_order: 随心购订单明细
func (r *TaobaoOmniorderGuideDataGetAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaoOmniorderGuideDataGetAPIRequest) GetType() string {
	return r._type
}

// SetStartTime is StartTime Setter
// 拉取数据开始时间
func (r *TaobaoOmniorderGuideDataGetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoOmniorderGuideDataGetAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetPageNo is PageNo Setter
// 页码，从1开始
func (r *TaobaoOmniorderGuideDataGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoOmniorderGuideDataGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页数量，不能大于1000
func (r *TaobaoOmniorderGuideDataGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoOmniorderGuideDataGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolTaobaoOmniorderGuideDataGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOmniorderGuideDataGetRequest()
	},
}

// GetTaobaoOmniorderGuideDataGetRequest 从 sync.Pool 获取 TaobaoOmniorderGuideDataGetAPIRequest
func GetTaobaoOmniorderGuideDataGetAPIRequest() *TaobaoOmniorderGuideDataGetAPIRequest {
	return poolTaobaoOmniorderGuideDataGetAPIRequest.Get().(*TaobaoOmniorderGuideDataGetAPIRequest)
}

// ReleaseTaobaoOmniorderGuideDataGetAPIRequest 将 TaobaoOmniorderGuideDataGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoOmniorderGuideDataGetAPIRequest(v *TaobaoOmniorderGuideDataGetAPIRequest) {
	v.Reset()
	poolTaobaoOmniorderGuideDataGetAPIRequest.Put(v)
}
