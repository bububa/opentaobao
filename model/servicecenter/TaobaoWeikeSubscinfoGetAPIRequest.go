package servicecenter

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWeikeSubscinfoGetAPIRequest 需求订单查询接口 API请求
// taobao.weike.subscinfo.get
//
// 需求订单查询接口
type TaobaoWeikeSubscinfoGetAPIRequest struct {
	model.Params
	// 商家旺旺名称
	_sellerName string
	// 时间范围开始时间
	_startTime string
	// 时间范围结束时间
	_endTime string
	// 页码
	_pageNum int64
}

// NewTaobaoWeikeSubscinfoGetRequest 初始化TaobaoWeikeSubscinfoGetAPIRequest对象
func NewTaobaoWeikeSubscinfoGetRequest() *TaobaoWeikeSubscinfoGetAPIRequest {
	return &TaobaoWeikeSubscinfoGetAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWeikeSubscinfoGetAPIRequest) Reset() {
	r._sellerName = ""
	r._startTime = ""
	r._endTime = ""
	r._pageNum = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWeikeSubscinfoGetAPIRequest) GetApiMethodName() string {
	return "taobao.weike.subscinfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWeikeSubscinfoGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWeikeSubscinfoGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSellerName is SellerName Setter
// 商家旺旺名称
func (r *TaobaoWeikeSubscinfoGetAPIRequest) SetSellerName(_sellerName string) error {
	r._sellerName = _sellerName
	r.Set("seller_name", _sellerName)
	return nil
}

// GetSellerName SellerName Getter
func (r TaobaoWeikeSubscinfoGetAPIRequest) GetSellerName() string {
	return r._sellerName
}

// SetStartTime is StartTime Setter
// 时间范围开始时间
func (r *TaobaoWeikeSubscinfoGetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoWeikeSubscinfoGetAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 时间范围结束时间
func (r *TaobaoWeikeSubscinfoGetAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoWeikeSubscinfoGetAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetPageNum is PageNum Setter
// 页码
func (r *TaobaoWeikeSubscinfoGetAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r TaobaoWeikeSubscinfoGetAPIRequest) GetPageNum() int64 {
	return r._pageNum
}

var poolTaobaoWeikeSubscinfoGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWeikeSubscinfoGetRequest()
	},
}

// GetTaobaoWeikeSubscinfoGetRequest 从 sync.Pool 获取 TaobaoWeikeSubscinfoGetAPIRequest
func GetTaobaoWeikeSubscinfoGetAPIRequest() *TaobaoWeikeSubscinfoGetAPIRequest {
	return poolTaobaoWeikeSubscinfoGetAPIRequest.Get().(*TaobaoWeikeSubscinfoGetAPIRequest)
}

// ReleaseTaobaoWeikeSubscinfoGetAPIRequest 将 TaobaoWeikeSubscinfoGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoWeikeSubscinfoGetAPIRequest(v *TaobaoWeikeSubscinfoGetAPIRequest) {
	v.Reset()
	poolTaobaoWeikeSubscinfoGetAPIRequest.Put(v)
}
