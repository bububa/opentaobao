package smartstore

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoNrrsAcceptanceTaskUpdateResultAPIRequest 更新验收任务支付宝结果 API请求
// taobao.nrrs.acceptance.task.updateResult
//
// 智慧门店商家验收任务检查相关接口-更新支付宝的验收结果。
type TaobaoNrrsAcceptanceTaskUpdateResultAPIRequest struct {
	model.Params
	// 系统自动生成
	_alipayResultList []AlipayCheckResult
	// 任务ID
	_taskId string
}

// NewTaobaoNrrsAcceptanceTaskUpdateResultRequest 初始化TaobaoNrrsAcceptanceTaskUpdateResultAPIRequest对象
func NewTaobaoNrrsAcceptanceTaskUpdateResultRequest() *TaobaoNrrsAcceptanceTaskUpdateResultAPIRequest {
	return &TaobaoNrrsAcceptanceTaskUpdateResultAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoNrrsAcceptanceTaskUpdateResultAPIRequest) GetApiMethodName() string {
	return "taobao.nrrs.acceptance.task.updateResult"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoNrrsAcceptanceTaskUpdateResultAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetAlipayResultList is AlipayResultList Setter
// 系统自动生成
func (r *TaobaoNrrsAcceptanceTaskUpdateResultAPIRequest) SetAlipayResultList(_alipayResultList []AlipayCheckResult) error {
	r._alipayResultList = _alipayResultList
	r.Set("alipay_result_list", _alipayResultList)
	return nil
}

// GetAlipayResultList AlipayResultList Getter
func (r TaobaoNrrsAcceptanceTaskUpdateResultAPIRequest) GetAlipayResultList() []AlipayCheckResult {
	return r._alipayResultList
}

// SetTaskId is TaskId Setter
// 任务ID
func (r *TaobaoNrrsAcceptanceTaskUpdateResultAPIRequest) SetTaskId(_taskId string) error {
	r._taskId = _taskId
	r.Set("task_id", _taskId)
	return nil
}

// GetTaskId TaskId Getter
func (r TaobaoNrrsAcceptanceTaskUpdateResultAPIRequest) GetTaskId() string {
	return r._taskId
}
