package fenxiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoDistributorsGetAPIRequest 获取分销商信息 API请求
// taobao.fenxiao.distributors.get
//
// 查询和当前登录供应商有合作关系的分销商的信息
type TaobaoFenxiaoDistributorsGetAPIRequest struct {
	model.Params
	// 分销商用户名列表。多个之间以“,”分隔;最多支持50个分销商用户名。
	_nicks string
}

// NewTaobaoFenxiaoDistributorsGetRequest 初始化TaobaoFenxiaoDistributorsGetAPIRequest对象
func NewTaobaoFenxiaoDistributorsGetRequest() *TaobaoFenxiaoDistributorsGetAPIRequest {
	return &TaobaoFenxiaoDistributorsGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFenxiaoDistributorsGetAPIRequest) Reset() {
	r._nicks = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoDistributorsGetAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.distributors.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoDistributorsGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFenxiaoDistributorsGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNicks is Nicks Setter
// 分销商用户名列表。多个之间以“,”分隔;最多支持50个分销商用户名。
func (r *TaobaoFenxiaoDistributorsGetAPIRequest) SetNicks(_nicks string) error {
	r._nicks = _nicks
	r.Set("nicks", _nicks)
	return nil
}

// GetNicks Nicks Getter
func (r TaobaoFenxiaoDistributorsGetAPIRequest) GetNicks() string {
	return r._nicks
}

var poolTaobaoFenxiaoDistributorsGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFenxiaoDistributorsGetRequest()
	},
}

// GetTaobaoFenxiaoDistributorsGetRequest 从 sync.Pool 获取 TaobaoFenxiaoDistributorsGetAPIRequest
func GetTaobaoFenxiaoDistributorsGetAPIRequest() *TaobaoFenxiaoDistributorsGetAPIRequest {
	return poolTaobaoFenxiaoDistributorsGetAPIRequest.Get().(*TaobaoFenxiaoDistributorsGetAPIRequest)
}

// ReleaseTaobaoFenxiaoDistributorsGetAPIRequest 将 TaobaoFenxiaoDistributorsGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoFenxiaoDistributorsGetAPIRequest(v *TaobaoFenxiaoDistributorsGetAPIRequest) {
	v.Reset()
	poolTaobaoFenxiaoDistributorsGetAPIRequest.Put(v)
}
