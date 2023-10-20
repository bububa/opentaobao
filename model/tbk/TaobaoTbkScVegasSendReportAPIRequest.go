package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkscvegassendreportAPIRequest 淘宝客-服务商-查询红包发放个数 API请求
// taobao.tbk.sc.vegas.send.report
//
// 服务商使用。可通过此接口查询对应推广者账号下的红包发放个数。
type TaobaotbkscvegassendreportAPIRequest struct {
	model.Params
	// 统计日期
	_bizDate string
	// 媒体推广pid
	_pid string
	// 查询维度，不填写默认是pid维度
	_rptDim string
	// 渠道关系id
	_relationId int64
	// 已下线，后续不需要填写
	_activityId int64
	// 页码
	_pageNo int64
	// 每页大小
	_pageSize int64
	// 查询红包类型，1-超级红包，2-福利购，3-签到红包，4-福利直降，不传时默认查询超级红包数据
	_activityCategory int64
}

// NewTaobaotbkscvegassendreportRequest 初始化TaobaotbkscvegassendreportAPIRequest对象
func NewTaobaotbkscvegassendreportRequest() *TaobaotbkscvegassendreportAPIRequest {
	return &TaobaotbkscvegassendreportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkscvegassendreportAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.vegas.send.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkscvegassendreportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkscvegassendreportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizDate is BizDate Setter
// 统计日期
func (r *TaobaotbkscvegassendreportAPIRequest) SetBizDate(_bizDate string) error {
	r._bizDate = _bizDate
	r.Set("biz_date", _bizDate)
	return nil
}

// GetBizDate BizDate Getter
func (r TaobaotbkscvegassendreportAPIRequest) GetBizDate() string {
	return r._bizDate
}

// SetPid is Pid Setter
// 媒体推广pid
func (r *TaobaotbkscvegassendreportAPIRequest) SetPid(_pid string) error {
	r._pid = _pid
	r.Set("pid", _pid)
	return nil
}

// GetPid Pid Getter
func (r TaobaotbkscvegassendreportAPIRequest) GetPid() string {
	return r._pid
}

// SetRptDim is RptDim Setter
// 查询维度，不填写默认是pid维度
func (r *TaobaotbkscvegassendreportAPIRequest) SetRptDim(_rptDim string) error {
	r._rptDim = _rptDim
	r.Set("rpt_dim", _rptDim)
	return nil
}

// GetRptDim RptDim Getter
func (r TaobaotbkscvegassendreportAPIRequest) GetRptDim() string {
	return r._rptDim
}

// SetRelationId is RelationId Setter
// 渠道关系id
func (r *TaobaotbkscvegassendreportAPIRequest) SetRelationId(_relationId int64) error {
	r._relationId = _relationId
	r.Set("relation_id", _relationId)
	return nil
}

// GetRelationId RelationId Getter
func (r TaobaotbkscvegassendreportAPIRequest) GetRelationId() int64 {
	return r._relationId
}

// SetActivityId is ActivityId Setter
// 已下线，后续不需要填写
func (r *TaobaotbkscvegassendreportAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TaobaotbkscvegassendreportAPIRequest) GetActivityId() int64 {
	return r._activityId
}

// SetPageNo is PageNo Setter
// 页码
func (r *TaobaotbkscvegassendreportAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaotbkscvegassendreportAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页大小
func (r *TaobaotbkscvegassendreportAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaotbkscvegassendreportAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetActivityCategory is ActivityCategory Setter
// 查询红包类型，1-超级红包，2-福利购，3-签到红包，4-福利直降，不传时默认查询超级红包数据
func (r *TaobaotbkscvegassendreportAPIRequest) SetActivityCategory(_activityCategory int64) error {
	r._activityCategory = _activityCategory
	r.Set("activity_category", _activityCategory)
	return nil
}

// GetActivityCategory ActivityCategory Getter
func (r TaobaotbkscvegassendreportAPIRequest) GetActivityCategory() int64 {
	return r._activityCategory
}
