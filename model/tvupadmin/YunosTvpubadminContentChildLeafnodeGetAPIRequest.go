package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentChildLeafnodeGetAPIRequest 获取少儿大厅二级类目列表 API请求
// yunos.tvpubadmin.content.child.leafnode.get
//
// 获取少儿大厅二级类目列表的接口
type YunosTvpubadminContentChildLeafnodeGetAPIRequest struct {
	model.Params
}

// NewYunosTvpubadminContentChildLeafnodeGetRequest 初始化YunosTvpubadminContentChildLeafnodeGetAPIRequest对象
func NewYunosTvpubadminContentChildLeafnodeGetRequest() *YunosTvpubadminContentChildLeafnodeGetAPIRequest {
	return &YunosTvpubadminContentChildLeafnodeGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminContentChildLeafnodeGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentChildLeafnodeGetAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.child.leafnode.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentChildLeafnodeGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminContentChildLeafnodeGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolYunosTvpubadminContentChildLeafnodeGetAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminContentChildLeafnodeGetRequest()
	},
}

// GetYunosTvpubadminContentChildLeafnodeGetRequest 从 sync.Pool 获取 YunosTvpubadminContentChildLeafnodeGetAPIRequest
func GetYunosTvpubadminContentChildLeafnodeGetAPIRequest() *YunosTvpubadminContentChildLeafnodeGetAPIRequest {
	return poolYunosTvpubadminContentChildLeafnodeGetAPIRequest.Get().(*YunosTvpubadminContentChildLeafnodeGetAPIRequest)
}

// ReleaseYunosTvpubadminContentChildLeafnodeGetAPIRequest 将 YunosTvpubadminContentChildLeafnodeGetAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminContentChildLeafnodeGetAPIRequest(v *YunosTvpubadminContentChildLeafnodeGetAPIRequest) {
	v.Reset()
	poolYunosTvpubadminContentChildLeafnodeGetAPIRequest.Put(v)
}
