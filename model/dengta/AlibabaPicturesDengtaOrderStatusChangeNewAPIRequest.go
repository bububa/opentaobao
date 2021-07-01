package dengta

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaPicturesDengtaOrderStatusChangeNewAPIRequest
天下秀订单状态变更通知 API请求
alibaba.pictures.dengta.order.status.change.new

天下秀订单状态变更通知 */
type AlibabaPicturesDengtaOrderStatusChangeNewAPIRequest struct {
	model.Params
	// 拒绝原因
	_remark string
	// 新状态
	_status int64
	// 变更时间
	_changeTime string
	// ims订单编号
	_imsOrderId string
	// task 编号
	_aliTaskId string
	// 发布内容
	_taskContent string
	// 发布图片url列表
	_taskPic string
	// 扩展字段。json结构
	_extJson string
}

// NewAlibabaPicturesDengtaOrderStatusChangeNewRequest 初始化AlibabaPicturesDengtaOrderStatusChangeNewAPIRequest对象
func NewAlibabaPicturesDengtaOrderStatusChangeNewRequest() *AlibabaPicturesDengtaOrderStatusChangeNewAPIRequest {
	return &AlibabaPicturesDengtaOrderStatusChangeNewAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaPicturesDengtaOrderStatusChangeNewAPIRequest) GetApiMethodName() string {
	return "alibaba.pictures.dengta.order.status.change.new"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaPicturesDengtaOrderStatusChangeNewAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Remark Setter
// 拒绝原因
func (r *AlibabaPicturesDengtaOrderStatusChangeNewAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// Get Remark Getter
func (r AlibabaPicturesDengtaOrderStatusChangeNewAPIRequest) GetRemark() string {
	return r._remark
}

// Set is Status Setter
// 新状态
func (r *AlibabaPicturesDengtaOrderStatusChangeNewAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// Get Status Getter
func (r AlibabaPicturesDengtaOrderStatusChangeNewAPIRequest) GetStatus() int64 {
	return r._status
}

// Set is ChangeTime Setter
// 变更时间
func (r *AlibabaPicturesDengtaOrderStatusChangeNewAPIRequest) SetChangeTime(_changeTime string) error {
	r._changeTime = _changeTime
	r.Set("change_time", _changeTime)
	return nil
}

// Get ChangeTime Getter
func (r AlibabaPicturesDengtaOrderStatusChangeNewAPIRequest) GetChangeTime() string {
	return r._changeTime
}

// Set is ImsOrderId Setter
// ims订单编号
func (r *AlibabaPicturesDengtaOrderStatusChangeNewAPIRequest) SetImsOrderId(_imsOrderId string) error {
	r._imsOrderId = _imsOrderId
	r.Set("ims_order_id", _imsOrderId)
	return nil
}

// Get ImsOrderId Getter
func (r AlibabaPicturesDengtaOrderStatusChangeNewAPIRequest) GetImsOrderId() string {
	return r._imsOrderId
}

// Set is AliTaskId Setter
// task 编号
func (r *AlibabaPicturesDengtaOrderStatusChangeNewAPIRequest) SetAliTaskId(_aliTaskId string) error {
	r._aliTaskId = _aliTaskId
	r.Set("ali_task_id", _aliTaskId)
	return nil
}

// Get AliTaskId Getter
func (r AlibabaPicturesDengtaOrderStatusChangeNewAPIRequest) GetAliTaskId() string {
	return r._aliTaskId
}

// Set is TaskContent Setter
// 发布内容
func (r *AlibabaPicturesDengtaOrderStatusChangeNewAPIRequest) SetTaskContent(_taskContent string) error {
	r._taskContent = _taskContent
	r.Set("task_content", _taskContent)
	return nil
}

// Get TaskContent Getter
func (r AlibabaPicturesDengtaOrderStatusChangeNewAPIRequest) GetTaskContent() string {
	return r._taskContent
}

// Set is TaskPic Setter
// 发布图片url列表
func (r *AlibabaPicturesDengtaOrderStatusChangeNewAPIRequest) SetTaskPic(_taskPic string) error {
	r._taskPic = _taskPic
	r.Set("task_pic", _taskPic)
	return nil
}

// Get TaskPic Getter
func (r AlibabaPicturesDengtaOrderStatusChangeNewAPIRequest) GetTaskPic() string {
	return r._taskPic
}

// Set is ExtJson Setter
// 扩展字段。json结构
func (r *AlibabaPicturesDengtaOrderStatusChangeNewAPIRequest) SetExtJson(_extJson string) error {
	r._extJson = _extJson
	r.Set("ext_json", _extJson)
	return nil
}

// Get ExtJson Getter
func (r AlibabaPicturesDengtaOrderStatusChangeNewAPIRequest) GetExtJson() string {
	return r._extJson
}
