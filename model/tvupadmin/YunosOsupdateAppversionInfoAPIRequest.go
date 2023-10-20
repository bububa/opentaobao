package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosOsupdateAppversionInfoAPIRequest 获取应用升级详情 API请求
// yunos.osupdate.appversion.info
//
// 获取应用升级详情
type YunosOsupdateAppversionInfoAPIRequest struct {
	model.Params
	// 升级任务ID
	_id int64
}

// NewYunosOsupdateAppversionInfoRequest 初始化YunosOsupdateAppversionInfoAPIRequest对象
func NewYunosOsupdateAppversionInfoRequest() *YunosOsupdateAppversionInfoAPIRequest {
	return &YunosOsupdateAppversionInfoAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosOsupdateAppversionInfoAPIRequest) Reset() {
	r._id = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosOsupdateAppversionInfoAPIRequest) GetApiMethodName() string {
	return "yunos.osupdate.appversion.info"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosOsupdateAppversionInfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosOsupdateAppversionInfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// 升级任务ID
func (r *YunosOsupdateAppversionInfoAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r YunosOsupdateAppversionInfoAPIRequest) GetId() int64 {
	return r._id
}

var poolYunosOsupdateAppversionInfoAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosOsupdateAppversionInfoRequest()
	},
}

// GetYunosOsupdateAppversionInfoRequest 从 sync.Pool 获取 YunosOsupdateAppversionInfoAPIRequest
func GetYunosOsupdateAppversionInfoAPIRequest() *YunosOsupdateAppversionInfoAPIRequest {
	return poolYunosOsupdateAppversionInfoAPIRequest.Get().(*YunosOsupdateAppversionInfoAPIRequest)
}

// ReleaseYunosOsupdateAppversionInfoAPIRequest 将 YunosOsupdateAppversionInfoAPIRequest 放入 sync.Pool
func ReleaseYunosOsupdateAppversionInfoAPIRequest(v *YunosOsupdateAppversionInfoAPIRequest) {
	v.Reset()
	poolYunosOsupdateAppversionInfoAPIRequest.Put(v)
}
