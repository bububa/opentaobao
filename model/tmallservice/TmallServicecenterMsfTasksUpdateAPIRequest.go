package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterMsfTasksUpdateAPIRequest 喵师傅工人任务批量完成接口 API请求
// tmall.servicecenter.msf.tasks.update
//
// 喵师傅工人任务批量完成接口
type TmallServicecenterMsfTasksUpdateAPIRequest struct {
	model.Params
	// 子订单号列表。最多100个
	_bizOrderIds []int64
	// 周期序号。必须与子订单列表对应
	_seqs []int64
	// 服务编码
	_serviceCode string
	// 调用来源
	_source string
	// 工人手机号
	_workerMobile int64
}

// NewTmallServicecenterMsfTasksUpdateRequest 初始化TmallServicecenterMsfTasksUpdateAPIRequest对象
func NewTmallServicecenterMsfTasksUpdateRequest() *TmallServicecenterMsfTasksUpdateAPIRequest {
	return &TmallServicecenterMsfTasksUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterMsfTasksUpdateAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.msf.tasks.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterMsfTasksUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBizOrderIds is BizOrderIds Setter
// 子订单号列表。最多100个
func (r *TmallServicecenterMsfTasksUpdateAPIRequest) SetBizOrderIds(_bizOrderIds []int64) error {
	r._bizOrderIds = _bizOrderIds
	r.Set("biz_order_ids", _bizOrderIds)
	return nil
}

// GetBizOrderIds BizOrderIds Getter
func (r TmallServicecenterMsfTasksUpdateAPIRequest) GetBizOrderIds() []int64 {
	return r._bizOrderIds
}

// SetSeqs is Seqs Setter
// 周期序号。必须与子订单列表对应
func (r *TmallServicecenterMsfTasksUpdateAPIRequest) SetSeqs(_seqs []int64) error {
	r._seqs = _seqs
	r.Set("seqs", _seqs)
	return nil
}

// GetSeqs Seqs Getter
func (r TmallServicecenterMsfTasksUpdateAPIRequest) GetSeqs() []int64 {
	return r._seqs
}

// SetServiceCode is ServiceCode Setter
// 服务编码
func (r *TmallServicecenterMsfTasksUpdateAPIRequest) SetServiceCode(_serviceCode string) error {
	r._serviceCode = _serviceCode
	r.Set("service_code", _serviceCode)
	return nil
}

// GetServiceCode ServiceCode Getter
func (r TmallServicecenterMsfTasksUpdateAPIRequest) GetServiceCode() string {
	return r._serviceCode
}

// SetSource is Source Setter
// 调用来源
func (r *TmallServicecenterMsfTasksUpdateAPIRequest) SetSource(_source string) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// GetSource Source Getter
func (r TmallServicecenterMsfTasksUpdateAPIRequest) GetSource() string {
	return r._source
}

// SetWorkerMobile is WorkerMobile Setter
// 工人手机号
func (r *TmallServicecenterMsfTasksUpdateAPIRequest) SetWorkerMobile(_workerMobile int64) error {
	r._workerMobile = _workerMobile
	r.Set("worker_mobile", _workerMobile)
	return nil
}

// GetWorkerMobile WorkerMobile Getter
func (r TmallServicecenterMsfTasksUpdateAPIRequest) GetWorkerMobile() int64 {
	return r._workerMobile
}
