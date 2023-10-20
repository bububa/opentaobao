package tblogistics

import (
	"sync"
)

// ReportExceptionRequest 结构体
type ReportExceptionRequest struct {
	// 异常包裹图片链接
	PackageImage []string `json:"package_image,omitempty" xml:"package_image>string,omitempty"`
	// 异常类型
	ErrorType []string `json:"error_type,omitempty" xml:"error_type>string,omitempty"`
	// 物流主体
	LogisticsOwner string `json:"logistics_owner,omitempty" xml:"logistics_owner,omitempty"`
	// 物流服务商ID
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 服务商仓code
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 异常包裹运单号
	MailNo string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
	// 异常作业单号
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
	// 作业单类型
	OrderType string `json:"order_type,omitempty" xml:"order_type,omitempty"`
	// 异常包裹WMS单据号
	OuterOrderCode string `json:"outer_order_code,omitempty" xml:"outer_order_code,omitempty"`
	// 异常包裹快递公司code
	TmsCpCode string `json:"tms_cp_code,omitempty" xml:"tms_cp_code,omitempty"`
	// 异常操作时间
	OperateTime string `json:"operate_time,omitempty" xml:"operate_time,omitempty"`
	// 异常包裹寄件人手机号
	SenderMobile string `json:"sender_mobile,omitempty" xml:"sender_mobile,omitempty"`
	// 异常发生的仓内作业节点
	WmsOperateNode string `json:"wms_operate_node,omitempty" xml:"wms_operate_node,omitempty"`
	// 异常操作节点
	ErrorOperateStatus string `json:"error_operate_status,omitempty" xml:"error_operate_status,omitempty"`
	// 异常说明
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 包裹货主
	PackageOwnerCode string `json:"package_owner_code,omitempty" xml:"package_owner_code,omitempty"`
	// 拓展属性
	ExtendProps string `json:"extend_props,omitempty" xml:"extend_props,omitempty"`
}

var poolReportExceptionRequest = sync.Pool{
	New: func() any {
		return new(ReportExceptionRequest)
	},
}

// GetReportExceptionRequest() 从对象池中获取ReportExceptionRequest
func GetReportExceptionRequest() *ReportExceptionRequest {
	return poolReportExceptionRequest.Get().(*ReportExceptionRequest)
}

// ReleaseReportExceptionRequest 释放ReportExceptionRequest
func ReleaseReportExceptionRequest(v *ReportExceptionRequest) {
	v.PackageImage = v.PackageImage[:0]
	v.ErrorType = v.ErrorType[:0]
	v.LogisticsOwner = ""
	v.SupplierId = ""
	v.WarehouseCode = ""
	v.MailNo = ""
	v.OrderCode = ""
	v.OrderType = ""
	v.OuterOrderCode = ""
	v.TmsCpCode = ""
	v.OperateTime = ""
	v.SenderMobile = ""
	v.WmsOperateNode = ""
	v.ErrorOperateStatus = ""
	v.ErrorMsg = ""
	v.PackageOwnerCode = ""
	v.ExtendProps = ""
	poolReportExceptionRequest.Put(v)
}
