package servicecenter

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFuwuSpBillreordAddAPIRequest 内购服务确认单明细上传接口 API请求
// taobao.fuwu.sp.billreord.add
//
// isv能通过该接口上传确认单明细数据
type TaobaoFuwuSpBillreordAddAPIRequest struct {
	model.Params
	// 确认单的账单明细
	_paramBillRecordDTO *BillRecordDto
}

// NewTaobaoFuwuSpBillreordAddRequest 初始化TaobaoFuwuSpBillreordAddAPIRequest对象
func NewTaobaoFuwuSpBillreordAddRequest() *TaobaoFuwuSpBillreordAddAPIRequest {
	return &TaobaoFuwuSpBillreordAddAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFuwuSpBillreordAddAPIRequest) Reset() {
	r._paramBillRecordDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFuwuSpBillreordAddAPIRequest) GetApiMethodName() string {
	return "taobao.fuwu.sp.billreord.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFuwuSpBillreordAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFuwuSpBillreordAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBillRecordDTO is ParamBillRecordDTO Setter
// 确认单的账单明细
func (r *TaobaoFuwuSpBillreordAddAPIRequest) SetParamBillRecordDTO(_paramBillRecordDTO *BillRecordDto) error {
	r._paramBillRecordDTO = _paramBillRecordDTO
	r.Set("param_bill_record_d_t_o", _paramBillRecordDTO)
	return nil
}

// GetParamBillRecordDTO ParamBillRecordDTO Getter
func (r TaobaoFuwuSpBillreordAddAPIRequest) GetParamBillRecordDTO() *BillRecordDto {
	return r._paramBillRecordDTO
}

var poolTaobaoFuwuSpBillreordAddAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFuwuSpBillreordAddRequest()
	},
}

// GetTaobaoFuwuSpBillreordAddRequest 从 sync.Pool 获取 TaobaoFuwuSpBillreordAddAPIRequest
func GetTaobaoFuwuSpBillreordAddAPIRequest() *TaobaoFuwuSpBillreordAddAPIRequest {
	return poolTaobaoFuwuSpBillreordAddAPIRequest.Get().(*TaobaoFuwuSpBillreordAddAPIRequest)
}

// ReleaseTaobaoFuwuSpBillreordAddAPIRequest 将 TaobaoFuwuSpBillreordAddAPIRequest 放入 sync.Pool
func ReleaseTaobaoFuwuSpBillreordAddAPIRequest(v *TaobaoFuwuSpBillreordAddAPIRequest) {
	v.Reset()
	poolTaobaoFuwuSpBillreordAddAPIRequest.Put(v)
}
