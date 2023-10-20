package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentShowEditAPIRequest 媒资节目信息修改 API请求
// yunos.tvpubadmin.content.show.edit
//
// 供迎客松修改媒资节目信息
type YunosTvpubadminContentShowEditAPIRequest struct {
	model.Params
	// 请求入参，JSON格式
	_data string
}

// NewYunosTvpubadminContentShowEditRequest 初始化YunosTvpubadminContentShowEditAPIRequest对象
func NewYunosTvpubadminContentShowEditRequest() *YunosTvpubadminContentShowEditAPIRequest {
	return &YunosTvpubadminContentShowEditAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminContentShowEditAPIRequest) Reset() {
	r._data = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentShowEditAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.show.edit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentShowEditAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminContentShowEditAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetData is Data Setter
// 请求入参，JSON格式
func (r *YunosTvpubadminContentShowEditAPIRequest) SetData(_data string) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// GetData Data Getter
func (r YunosTvpubadminContentShowEditAPIRequest) GetData() string {
	return r._data
}

var poolYunosTvpubadminContentShowEditAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminContentShowEditRequest()
	},
}

// GetYunosTvpubadminContentShowEditRequest 从 sync.Pool 获取 YunosTvpubadminContentShowEditAPIRequest
func GetYunosTvpubadminContentShowEditAPIRequest() *YunosTvpubadminContentShowEditAPIRequest {
	return poolYunosTvpubadminContentShowEditAPIRequest.Get().(*YunosTvpubadminContentShowEditAPIRequest)
}

// ReleaseYunosTvpubadminContentShowEditAPIRequest 将 YunosTvpubadminContentShowEditAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminContentShowEditAPIRequest(v *YunosTvpubadminContentShowEditAPIRequest) {
	v.Reset()
	poolYunosTvpubadminContentShowEditAPIRequest.Put(v)
}
