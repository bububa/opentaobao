package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkerUpdateAPIRequest 修改工人信息 API请求
// tmall.servicecenter.worker.update
//
// 修改工人信息。该接口为多个业务公用，部分字段可忽略。对于电器预约安装业务，同一个服务商，通过工人姓名+手机号+biz_type 保证唯一性。工人已存在才可以修改。
// 错误码如下
// 100000, 系统错误
// 100001, 工人信息校验失败
// 100002, 用户校验失败
// 100003, 操作失败
// 10004,工人信息为空
// 10005,服务商id为空或者服务商名称为空
// 10006, 工人不存在
// 10007, 工人已存在
// 10008, 缺少工人姓名
// 10009, 缺少工人电话
// 10010, 网点不存在
// 11000, category_id 无效
// 11001, biz_type 无效
// 20001,已查询到最后一页
type TmallServicecenterWorkerUpdateAPIRequest struct {
	model.Params
	// 工人信息
	_worker *WorkerDto
}

// NewTmallServicecenterWorkerUpdateRequest 初始化TmallServicecenterWorkerUpdateAPIRequest对象
func NewTmallServicecenterWorkerUpdateRequest() *TmallServicecenterWorkerUpdateAPIRequest {
	return &TmallServicecenterWorkerUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServicecenterWorkerUpdateAPIRequest) Reset() {
	r._worker = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkerUpdateAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.worker.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkerUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterWorkerUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorker is Worker Setter
// 工人信息
func (r *TmallServicecenterWorkerUpdateAPIRequest) SetWorker(_worker *WorkerDto) error {
	r._worker = _worker
	r.Set("worker", _worker)
	return nil
}

// GetWorker Worker Getter
func (r TmallServicecenterWorkerUpdateAPIRequest) GetWorker() *WorkerDto {
	return r._worker
}

var poolTmallServicecenterWorkerUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServicecenterWorkerUpdateRequest()
	},
}

// GetTmallServicecenterWorkerUpdateRequest 从 sync.Pool 获取 TmallServicecenterWorkerUpdateAPIRequest
func GetTmallServicecenterWorkerUpdateAPIRequest() *TmallServicecenterWorkerUpdateAPIRequest {
	return poolTmallServicecenterWorkerUpdateAPIRequest.Get().(*TmallServicecenterWorkerUpdateAPIRequest)
}

// ReleaseTmallServicecenterWorkerUpdateAPIRequest 将 TmallServicecenterWorkerUpdateAPIRequest 放入 sync.Pool
func ReleaseTmallServicecenterWorkerUpdateAPIRequest(v *TmallServicecenterWorkerUpdateAPIRequest) {
	v.Reset()
	poolTmallServicecenterWorkerUpdateAPIRequest.Put(v)
}
