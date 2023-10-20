package fenxiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoGradesGetAPIRequest 分销商等级查询 API请求
// taobao.fenxiao.grades.get
//
// 根据供应商ID，查询他的分销商等级信息
type TaobaoFenxiaoGradesGetAPIRequest struct {
	model.Params
}

// NewTaobaoFenxiaoGradesGetRequest 初始化TaobaoFenxiaoGradesGetAPIRequest对象
func NewTaobaoFenxiaoGradesGetRequest() *TaobaoFenxiaoGradesGetAPIRequest {
	return &TaobaoFenxiaoGradesGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFenxiaoGradesGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoGradesGetAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.grades.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoGradesGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFenxiaoGradesGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTaobaoFenxiaoGradesGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFenxiaoGradesGetRequest()
	},
}

// GetTaobaoFenxiaoGradesGetRequest 从 sync.Pool 获取 TaobaoFenxiaoGradesGetAPIRequest
func GetTaobaoFenxiaoGradesGetAPIRequest() *TaobaoFenxiaoGradesGetAPIRequest {
	return poolTaobaoFenxiaoGradesGetAPIRequest.Get().(*TaobaoFenxiaoGradesGetAPIRequest)
}

// ReleaseTaobaoFenxiaoGradesGetAPIRequest 将 TaobaoFenxiaoGradesGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoFenxiaoGradesGetAPIRequest(v *TaobaoFenxiaoGradesGetAPIRequest) {
	v.Reset()
	poolTaobaoFenxiaoGradesGetAPIRequest.Put(v)
}
