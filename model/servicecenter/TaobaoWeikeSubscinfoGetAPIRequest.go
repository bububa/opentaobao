package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoweikesubscinfogetAPIRequest 需求订单查询接口 API请求
// taobao.weike.subscinfo.get
//
// 需求订单查询接口
type TaobaoweikesubscinfogetAPIRequest struct {
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

// NewTaobaoweikesubscinfogetRequest 初始化TaobaoweikesubscinfogetAPIRequest对象
func NewTaobaoweikesubscinfogetRequest() *TaobaoweikesubscinfogetAPIRequest {
	return &TaobaoweikesubscinfogetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoweikesubscinfogetAPIRequest) GetApiMethodName() string {
	return "taobao.weike.subscinfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoweikesubscinfogetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoweikesubscinfogetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSellerName is SellerName Setter
// 商家旺旺名称
func (r *TaobaoweikesubscinfogetAPIRequest) SetSellerName(_sellerName string) error {
	r._sellerName = _sellerName
	r.Set("seller_name", _sellerName)
	return nil
}

// GetSellerName SellerName Getter
func (r TaobaoweikesubscinfogetAPIRequest) GetSellerName() string {
	return r._sellerName
}

// SetStartTime is StartTime Setter
// 时间范围开始时间
func (r *TaobaoweikesubscinfogetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoweikesubscinfogetAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 时间范围结束时间
func (r *TaobaoweikesubscinfogetAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoweikesubscinfogetAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetPageNum is PageNum Setter
// 页码
func (r *TaobaoweikesubscinfogetAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r TaobaoweikesubscinfogetAPIRequest) GetPageNum() int64 {
	return r._pageNum
}
