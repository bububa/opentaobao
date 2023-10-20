package tblogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsinstanttracesearchAPIRequest 物流详情查询 API请求
// taobao.logistics.instant.trace.search
//
// 物流详情查询
type TaobaologisticsinstanttracesearchAPIRequest struct {
	model.Params
	// 运单号
	_mailNo string
	// 子订单列表
	_subTid string
	// 是否拆单
	_isSplit int64
	// 交易单号
	_tid int64
}

// NewTaobaologisticsinstanttracesearchRequest 初始化TaobaologisticsinstanttracesearchAPIRequest对象
func NewTaobaologisticsinstanttracesearchRequest() *TaobaologisticsinstanttracesearchAPIRequest {
	return &TaobaologisticsinstanttracesearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticsinstanttracesearchAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.instant.trace.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticsinstanttracesearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticsinstanttracesearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMailNo is MailNo Setter
// 运单号
func (r *TaobaologisticsinstanttracesearchAPIRequest) SetMailNo(_mailNo string) error {
	r._mailNo = _mailNo
	r.Set("mail_no", _mailNo)
	return nil
}

// GetMailNo MailNo Getter
func (r TaobaologisticsinstanttracesearchAPIRequest) GetMailNo() string {
	return r._mailNo
}

// SetSubTid is SubTid Setter
// 子订单列表
func (r *TaobaologisticsinstanttracesearchAPIRequest) SetSubTid(_subTid string) error {
	r._subTid = _subTid
	r.Set("sub_tid", _subTid)
	return nil
}

// GetSubTid SubTid Getter
func (r TaobaologisticsinstanttracesearchAPIRequest) GetSubTid() string {
	return r._subTid
}

// SetIsSplit is IsSplit Setter
// 是否拆单
func (r *TaobaologisticsinstanttracesearchAPIRequest) SetIsSplit(_isSplit int64) error {
	r._isSplit = _isSplit
	r.Set("is_split", _isSplit)
	return nil
}

// GetIsSplit IsSplit Getter
func (r TaobaologisticsinstanttracesearchAPIRequest) GetIsSplit() int64 {
	return r._isSplit
}

// SetTid is Tid Setter
// 交易单号
func (r *TaobaologisticsinstanttracesearchAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaologisticsinstanttracesearchAPIRequest) GetTid() int64 {
	return r._tid
}
