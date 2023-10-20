package fenxiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoProductcatsGetAPIRequest 查询产品线列表 API请求
// taobao.fenxiao.productcats.get
//
// 查询供应商的所有产品线数据。根据登陆用户来查询，不需要其他入参
type TaobaoFenxiaoProductcatsGetAPIRequest struct {
	model.Params
	// 返回字段列表
	_fields string
}

// NewTaobaoFenxiaoProductcatsGetRequest 初始化TaobaoFenxiaoProductcatsGetAPIRequest对象
func NewTaobaoFenxiaoProductcatsGetRequest() *TaobaoFenxiaoProductcatsGetAPIRequest {
	return &TaobaoFenxiaoProductcatsGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFenxiaoProductcatsGetAPIRequest) Reset() {
	r._fields = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoProductcatsGetAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.productcats.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoProductcatsGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFenxiaoProductcatsGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 返回字段列表
func (r *TaobaoFenxiaoProductcatsGetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoFenxiaoProductcatsGetAPIRequest) GetFields() string {
	return r._fields
}

var poolTaobaoFenxiaoProductcatsGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFenxiaoProductcatsGetRequest()
	},
}

// GetTaobaoFenxiaoProductcatsGetRequest 从 sync.Pool 获取 TaobaoFenxiaoProductcatsGetAPIRequest
func GetTaobaoFenxiaoProductcatsGetAPIRequest() *TaobaoFenxiaoProductcatsGetAPIRequest {
	return poolTaobaoFenxiaoProductcatsGetAPIRequest.Get().(*TaobaoFenxiaoProductcatsGetAPIRequest)
}

// ReleaseTaobaoFenxiaoProductcatsGetAPIRequest 将 TaobaoFenxiaoProductcatsGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoFenxiaoProductcatsGetAPIRequest(v *TaobaoFenxiaoProductcatsGetAPIRequest) {
	v.Reset()
	poolTaobaoFenxiaoProductcatsGetAPIRequest.Put(v)
}
