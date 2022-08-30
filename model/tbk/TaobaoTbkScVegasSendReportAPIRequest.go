package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScVegasSendReportAPIRequest 淘宝客-服务商-查询红包发放个数 API请求
// taobao.tbk.sc.vegas.send.report
//
// 服务商使用。可通过此接口查询对应推广者账号下的红包发放个数。
type TaobaoTbkScVegasSendReportAPIRequest struct {
	model.Params
	// 统计日期
	_bizDate string
	// 渠道关系id
	_relationId int64
	// 红包活动id：1462
	_activityId int64
	// 页码
	_pageNo int64
	// 每页大小
	_pageSize int64
}

// NewTaobaoTbkScVegasSendReportRequest 初始化TaobaoTbkScVegasSendReportAPIRequest对象
func NewTaobaoTbkScVegasSendReportRequest() *TaobaoTbkScVegasSendReportAPIRequest {
	return &TaobaoTbkScVegasSendReportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkScVegasSendReportAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.vegas.send.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkScVegasSendReportAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBizDate is BizDate Setter
// 统计日期
func (r *TaobaoTbkScVegasSendReportAPIRequest) SetBizDate(_bizDate string) error {
	r._bizDate = _bizDate
	r.Set("biz_date", _bizDate)
	return nil
}

// GetBizDate BizDate Getter
func (r TaobaoTbkScVegasSendReportAPIRequest) GetBizDate() string {
	return r._bizDate
}

// SetRelationId is RelationId Setter
// 渠道关系id
func (r *TaobaoTbkScVegasSendReportAPIRequest) SetRelationId(_relationId int64) error {
	r._relationId = _relationId
	r.Set("relation_id", _relationId)
	return nil
}

// GetRelationId RelationId Getter
func (r TaobaoTbkScVegasSendReportAPIRequest) GetRelationId() int64 {
	return r._relationId
}

// SetActivityId is ActivityId Setter
// 红包活动id：1462
func (r *TaobaoTbkScVegasSendReportAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TaobaoTbkScVegasSendReportAPIRequest) GetActivityId() int64 {
	return r._activityId
}

// SetPageNo is PageNo Setter
// 页码
func (r *TaobaoTbkScVegasSendReportAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoTbkScVegasSendReportAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页大小
func (r *TaobaoTbkScVegasSendReportAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoTbkScVegasSendReportAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
