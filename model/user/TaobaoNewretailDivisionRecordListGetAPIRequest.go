package user

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoNewretailDivisionRecordListGetAPIRequest 导购分佣明细列表 API请求
// taobao.newretail.division.record.list.get
//
// 提供分页查询导购分佣明细的能力
type TaobaoNewretailDivisionRecordListGetAPIRequest struct {
	model.Params
	// 入参
	_param *TopDivisionRecordReqDto
}

// NewTaobaoNewretailDivisionRecordListGetRequest 初始化TaobaoNewretailDivisionRecordListGetAPIRequest对象
func NewTaobaoNewretailDivisionRecordListGetRequest() *TaobaoNewretailDivisionRecordListGetAPIRequest {
	return &TaobaoNewretailDivisionRecordListGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoNewretailDivisionRecordListGetAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoNewretailDivisionRecordListGetAPIRequest) GetApiMethodName() string {
	return "taobao.newretail.division.record.list.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoNewretailDivisionRecordListGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoNewretailDivisionRecordListGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *TaobaoNewretailDivisionRecordListGetAPIRequest) SetParam(_param *TopDivisionRecordReqDto) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TaobaoNewretailDivisionRecordListGetAPIRequest) GetParam() *TopDivisionRecordReqDto {
	return r._param
}

var poolTaobaoNewretailDivisionRecordListGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoNewretailDivisionRecordListGetRequest()
	},
}

// GetTaobaoNewretailDivisionRecordListGetRequest 从 sync.Pool 获取 TaobaoNewretailDivisionRecordListGetAPIRequest
func GetTaobaoNewretailDivisionRecordListGetAPIRequest() *TaobaoNewretailDivisionRecordListGetAPIRequest {
	return poolTaobaoNewretailDivisionRecordListGetAPIRequest.Get().(*TaobaoNewretailDivisionRecordListGetAPIRequest)
}

// ReleaseTaobaoNewretailDivisionRecordListGetAPIRequest 将 TaobaoNewretailDivisionRecordListGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoNewretailDivisionRecordListGetAPIRequest(v *TaobaoNewretailDivisionRecordListGetAPIRequest) {
	v.Reset()
	poolTaobaoNewretailDivisionRecordListGetAPIRequest.Put(v)
}
