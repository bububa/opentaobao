package jst

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOcTradetraceAlertsGetAPIRequest) Reset() {
	r._abnormalType = 0
	r._pageNo = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOcTradetraceAlertsGetAPIRequest) GetApiMethodName() string {
	return "taobao.oc.tradetrace.alerts.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOcTradetraceAlertsGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOcTradetraceAlertsGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAbnormalType is AbnormalType Setter
// 异常类型代码：发货回写淘宝异常=1，商家系统漏单提醒=2，发货超时提醒=3，物流寄送超时=4，买家签收超时=5，物流中转异常=6，退货寄回异常=7，订单追回提醒=8&#34;。
func (r *TaobaoOcTradetraceAlertsGetAPIRequest) SetAbnormalType(_abnormalType int64) error {
	r._abnormalType = _abnormalType
	r.Set("abnormal_type", _abnormalType)
	return nil
}

// GetAbnormalType AbnormalType Getter
func (r TaobaoOcTradetraceAlertsGetAPIRequest) GetAbnormalType() int64 {
	return r._abnormalType
}

// SetPageNo is PageNo Setter
// 返回数据的页码
func (r *TaobaoOcTradetraceAlertsGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoOcTradetraceAlertsGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 返回数据每页包含的记录数目
func (r *TaobaoOcTradetraceAlertsGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoOcTradetraceAlertsGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolTaobaoOcTradetraceAlertsGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOcTradetraceAlertsGetRequest()
	},
}

// GetTaobaoOcTradetraceAlertsGetRequest 从 sync.Pool 获取 TaobaoOcTradetraceAlertsGetAPIRequest
func GetTaobaoOcTradetraceAlertsGetAPIRequest() *TaobaoOcTradetraceAlertsGetAPIRequest {
	return poolTaobaoOcTradetraceAlertsGetAPIRequest.Get().(*TaobaoOcTradetraceAlertsGetAPIRequest)
}

// ReleaseTaobaoOcTradetraceAlertsGetAPIRequest 将 TaobaoOcTradetraceAlertsGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoOcTradetraceAlertsGetAPIRequest(v *TaobaoOcTradetraceAlertsGetAPIRequest) {
	v.Reset()
	poolTaobaoOcTradetraceAlertsGetAPIRequest.Put(v)
}
