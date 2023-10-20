package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterServicemonitormessageInfoAPIRequest 服务预警单查询接口 API请求
// tmall.servicecenter.servicemonitormessage.info
//
// 服务预警单查询接口,用于查询预警单最新状态
type TmallServicecenterServicemonitormessageInfoAPIRequest struct {
	model.Params
	// 预警单ID
	_id int64
}

// NewTmallServicecenterServicemonitormessageInfoRequest 初始化TmallServicecenterServicemonitormessageInfoAPIRequest对象
func NewTmallServicecenterServicemonitormessageInfoRequest() *TmallServicecenterServicemonitormessageInfoAPIRequest {
	return &TmallServicecenterServicemonitormessageInfoAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServicecenterServicemonitormessageInfoAPIRequest) Reset() {
	r._id = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterServicemonitormessageInfoAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.servicemonitormessage.info"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterServicemonitormessageInfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterServicemonitormessageInfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// 预警单ID
func (r *TmallServicecenterServicemonitormessageInfoAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TmallServicecenterServicemonitormessageInfoAPIRequest) GetId() int64 {
	return r._id
}

var poolTmallServicecenterServicemonitormessageInfoAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServicecenterServicemonitormessageInfoRequest()
	},
}

// GetTmallServicecenterServicemonitormessageInfoRequest 从 sync.Pool 获取 TmallServicecenterServicemonitormessageInfoAPIRequest
func GetTmallServicecenterServicemonitormessageInfoAPIRequest() *TmallServicecenterServicemonitormessageInfoAPIRequest {
	return poolTmallServicecenterServicemonitormessageInfoAPIRequest.Get().(*TmallServicecenterServicemonitormessageInfoAPIRequest)
}

// ReleaseTmallServicecenterServicemonitormessageInfoAPIRequest 将 TmallServicecenterServicemonitormessageInfoAPIRequest 放入 sync.Pool
func ReleaseTmallServicecenterServicemonitormessageInfoAPIRequest(v *TmallServicecenterServicemonitormessageInfoAPIRequest) {
	v.Reset()
	poolTmallServicecenterServicemonitormessageInfoAPIRequest.Put(v)
}
