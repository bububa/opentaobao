package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOcTradetraceAlertsGetAPIRequest 异常订单信息获取 API请求
// taobao.oc.tradetrace.alerts.get
//
// 提供订单预警模块的数据查询接口
type TaobaoOcTradetraceAlertsGetAPIRequest struct {
	model.Params
	// 异常类型代码：发货回写淘宝异常=1，商家系统漏单提醒=2，发货超时提醒=3，物流寄送超时=4，买家签收超时=5，物流中转异常=6，退货寄回异常=7，订单追回提醒=8"。
	_abnormalType int64
	// 返回数据的页码
	_pageNo int64
	// 返回数据每页包含的记录数目
	_pageSize int64
}

// NewTaobaoOcTradetraceAlertsGetRequest 初始化TaobaoOcTradetraceAlertsGetAPIRequest对象
func NewTaobaoOcTradetraceAlertsGetRequest() *TaobaoOcTradetraceAlertsGetAPIRequest {
	return &TaobaoOcTradetraceAlertsGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOcTradetraceAlertsGetAPIRequest) GetApiMethodName() string {
	return "taobao.oc.tradetrace.alerts.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOcTradetraceAlertsGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AbnormalType Setter
// 异常类型代码：发货回写淘宝异常=1，商家系统漏单提醒=2，发货超时提醒=3，物流寄送超时=4，买家签收超时=5，物流中转异常=6，退货寄回异常=7，订单追回提醒=8"。
func (r *TaobaoOcTradetraceAlertsGetAPIRequest) SetAbnormalType(_abnormalType int64) error {
	r._abnormalType = _abnormalType
	r.Set("abnormal_type", _abnormalType)
	return nil
}

// Get AbnormalType Getter
func (r TaobaoOcTradetraceAlertsGetAPIRequest) GetAbnormalType() int64 {
	return r._abnormalType
}

// Set is PageNo Setter
// 返回数据的页码
func (r *TaobaoOcTradetraceAlertsGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// Get PageNo Getter
func (r TaobaoOcTradetraceAlertsGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// Set is PageSize Setter
// 返回数据每页包含的记录数目
func (r *TaobaoOcTradetraceAlertsGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r TaobaoOcTradetraceAlertsGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
