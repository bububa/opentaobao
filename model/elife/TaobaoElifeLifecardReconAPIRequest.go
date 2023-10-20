package elife

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoElifeLifecardReconAPIRequest 查询对账文件地址接口 API请求
// taobao.elife.lifecard.recon
//
// 查询对账文件地址接口
type TaobaoElifeLifecardReconAPIRequest struct {
	model.Params
	// 对账日期(YYYYMMDD)
	_opDate string
}

// NewTaobaoElifeLifecardReconRequest 初始化TaobaoElifeLifecardReconAPIRequest对象
func NewTaobaoElifeLifecardReconRequest() *TaobaoElifeLifecardReconAPIRequest {
	return &TaobaoElifeLifecardReconAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoElifeLifecardReconAPIRequest) Reset() {
	r._opDate = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoElifeLifecardReconAPIRequest) GetApiMethodName() string {
	return "taobao.elife.lifecard.recon"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoElifeLifecardReconAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoElifeLifecardReconAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOpDate is OpDate Setter
// 对账日期(YYYYMMDD)
func (r *TaobaoElifeLifecardReconAPIRequest) SetOpDate(_opDate string) error {
	r._opDate = _opDate
	r.Set("op_date", _opDate)
	return nil
}

// GetOpDate OpDate Getter
func (r TaobaoElifeLifecardReconAPIRequest) GetOpDate() string {
	return r._opDate
}

var poolTaobaoElifeLifecardReconAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoElifeLifecardReconRequest()
	},
}

// GetTaobaoElifeLifecardReconRequest 从 sync.Pool 获取 TaobaoElifeLifecardReconAPIRequest
func GetTaobaoElifeLifecardReconAPIRequest() *TaobaoElifeLifecardReconAPIRequest {
	return poolTaobaoElifeLifecardReconAPIRequest.Get().(*TaobaoElifeLifecardReconAPIRequest)
}

// ReleaseTaobaoElifeLifecardReconAPIRequest 将 TaobaoElifeLifecardReconAPIRequest 放入 sync.Pool
func ReleaseTaobaoElifeLifecardReconAPIRequest(v *TaobaoElifeLifecardReconAPIRequest) {
	v.Reset()
	poolTaobaoElifeLifecardReconAPIRequest.Put(v)
}
