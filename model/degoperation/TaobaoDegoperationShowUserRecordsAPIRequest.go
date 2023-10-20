package degoperation

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDegoperationShowUserRecordsAPIRequest 用户中奖记录 API请求
// taobao.degoperation.show.user.records
//
// 用户中奖记录
type TaobaoDegoperationShowUserRecordsAPIRequest struct {
	model.Params
	// 系统信息
	_degAccessToken string
	// 活动后台配置
	_degAppKey string
	// 活动后台配置
	_eventKey string
	// 第几页
	_pageNumber int64
	// 分页尺寸
	_pageSize int64
}

// NewTaobaoDegoperationShowUserRecordsRequest 初始化TaobaoDegoperationShowUserRecordsAPIRequest对象
func NewTaobaoDegoperationShowUserRecordsRequest() *TaobaoDegoperationShowUserRecordsAPIRequest {
	return &TaobaoDegoperationShowUserRecordsAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoDegoperationShowUserRecordsAPIRequest) Reset() {
	r._degAccessToken = ""
	r._degAppKey = ""
	r._eventKey = ""
	r._pageNumber = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoDegoperationShowUserRecordsAPIRequest) GetApiMethodName() string {
	return "taobao.degoperation.show.user.records"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoDegoperationShowUserRecordsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoDegoperationShowUserRecordsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDegAccessToken is DegAccessToken Setter
// 系统信息
func (r *TaobaoDegoperationShowUserRecordsAPIRequest) SetDegAccessToken(_degAccessToken string) error {
	r._degAccessToken = _degAccessToken
	r.Set("deg_access_token", _degAccessToken)
	return nil
}

// GetDegAccessToken DegAccessToken Getter
func (r TaobaoDegoperationShowUserRecordsAPIRequest) GetDegAccessToken() string {
	return r._degAccessToken
}

// SetDegAppKey is DegAppKey Setter
// 活动后台配置
func (r *TaobaoDegoperationShowUserRecordsAPIRequest) SetDegAppKey(_degAppKey string) error {
	r._degAppKey = _degAppKey
	r.Set("deg_app_key", _degAppKey)
	return nil
}

// GetDegAppKey DegAppKey Getter
func (r TaobaoDegoperationShowUserRecordsAPIRequest) GetDegAppKey() string {
	return r._degAppKey
}

// SetEventKey is EventKey Setter
// 活动后台配置
func (r *TaobaoDegoperationShowUserRecordsAPIRequest) SetEventKey(_eventKey string) error {
	r._eventKey = _eventKey
	r.Set("event_key", _eventKey)
	return nil
}

// GetEventKey EventKey Getter
func (r TaobaoDegoperationShowUserRecordsAPIRequest) GetEventKey() string {
	return r._eventKey
}

// SetPageNumber is PageNumber Setter
// 第几页
func (r *TaobaoDegoperationShowUserRecordsAPIRequest) SetPageNumber(_pageNumber int64) error {
	r._pageNumber = _pageNumber
	r.Set("page_number", _pageNumber)
	return nil
}

// GetPageNumber PageNumber Getter
func (r TaobaoDegoperationShowUserRecordsAPIRequest) GetPageNumber() int64 {
	return r._pageNumber
}

// SetPageSize is PageSize Setter
// 分页尺寸
func (r *TaobaoDegoperationShowUserRecordsAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoDegoperationShowUserRecordsAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolTaobaoDegoperationShowUserRecordsAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoDegoperationShowUserRecordsRequest()
	},
}

// GetTaobaoDegoperationShowUserRecordsRequest 从 sync.Pool 获取 TaobaoDegoperationShowUserRecordsAPIRequest
func GetTaobaoDegoperationShowUserRecordsAPIRequest() *TaobaoDegoperationShowUserRecordsAPIRequest {
	return poolTaobaoDegoperationShowUserRecordsAPIRequest.Get().(*TaobaoDegoperationShowUserRecordsAPIRequest)
}

// ReleaseTaobaoDegoperationShowUserRecordsAPIRequest 将 TaobaoDegoperationShowUserRecordsAPIRequest 放入 sync.Pool
func ReleaseTaobaoDegoperationShowUserRecordsAPIRequest(v *TaobaoDegoperationShowUserRecordsAPIRequest) {
	v.Reset()
	poolTaobaoDegoperationShowUserRecordsAPIRequest.Put(v)
}
