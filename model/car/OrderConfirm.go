package car

import (
	"sync"
)

// OrderConfirm 结构体
type OrderConfirm struct {
	// 打标能力集合 核酸检测KEY
	Feature []string `json:"feature,omitempty" xml:"feature>string,omitempty"`
	// 确认时间
	ConfirmTime string `json:"confirm_time,omitempty" xml:"confirm_time,omitempty"`
	// 车辆描述,车辆颜色,候车地点等
	DriverCarDesc string `json:"driver_car_desc,omitempty" xml:"driver_car_desc,omitempty"`
	// 司机车型名称
	DriverCarName string `json:"driver_car_name,omitempty" xml:"driver_car_name,omitempty"`
	// 司机车牌号
	DriverCarNo string `json:"driver_car_no,omitempty" xml:"driver_car_no,omitempty"`
	// 司机名称
	DriverName string `json:"driver_name,omitempty" xml:"driver_name,omitempty"`
	// 司机联系方式。格式：加号+国家区号+空格+当地电话号码
	DriverTel string `json:"driver_tel,omitempty" xml:"driver_tel,omitempty"`
	// 阿里旅行订单ID
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 服务商标识
	ProviderId string `json:"provider_id,omitempty" xml:"provider_id,omitempty"`
	// 服务商订单号
	ThirdOrderId string `json:"third_order_id,omitempty" xml:"third_order_id,omitempty"`
	// 可选，卖家id
	SellerId string `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 司机图片，公网可访问链接
	DriverPic string `json:"driver_pic,omitempty" xml:"driver_pic,omitempty"`
	// 车型图片，公网可访问链接
	CarPic string `json:"car_pic,omitempty" xml:"car_pic,omitempty"`
	// 司机虚拟号
	DriverTrumpetPhone string `json:"driver_trumpet_phone,omitempty" xml:"driver_trumpet_phone,omitempty"`
	// 司机身份证号
	DriverIdNumber string `json:"driver_id_number,omitempty" xml:"driver_id_number,omitempty"`
	// 下游供应商logo图片(高德专用)
	SubPic string `json:"sub_pic,omitempty" xml:"sub_pic,omitempty"`
	// 下游供应商名称(高德专用)
	SubTitle string `json:"sub_title,omitempty" xml:"sub_title,omitempty"`
	// 下游供应商唯一标识(高德专用)
	SubKey string `json:"sub_key,omitempty" xml:"sub_key,omitempty"`
	// 车辆颜色
	DriverCarColor string `json:"driver_car_color,omitempty" xml:"driver_car_color,omitempty"`
	// 车品牌
	CarBrand string `json:"car_brand,omitempty" xml:"car_brand,omitempty"`
	// 车型名称
	CarModel string `json:"car_model,omitempty" xml:"car_model,omitempty"`
	// 车龄
	CarAge string `json:"car_age,omitempty" xml:"car_age,omitempty"`
	// 性别
	DeriverGender string `json:"deriver_gender,omitempty" xml:"deriver_gender,omitempty"`
	// 0应答1改派
	ConfirmType int64 `json:"confirm_type,omitempty" xml:"confirm_type,omitempty"`
	// 0:接送机 1：实时打车 2：租车(不传值默认为0)
	UseType int64 `json:"use_type,omitempty" xml:"use_type,omitempty"`
	// 飞猪车型id(30:出租车 31:经济型 32:舒适型 33:商务型 34:豪华型)
	CarTypeId int64 `json:"car_type_id,omitempty" xml:"car_type_id,omitempty"`
	// 本次用车是否支持 司机实时位置回传。若为true，则飞猪平台在用车实际开始时将从服务商处实时查询司机位置
	SupportRealTimePoi bool `json:"support_real_time_poi,omitempty" xml:"support_real_time_poi,omitempty"`
}

var poolOrderConfirm = sync.Pool{
	New: func() any {
		return new(OrderConfirm)
	},
}

// GetOrderConfirm() 从对象池中获取OrderConfirm
func GetOrderConfirm() *OrderConfirm {
	return poolOrderConfirm.Get().(*OrderConfirm)
}

// ReleaseOrderConfirm 释放OrderConfirm
func ReleaseOrderConfirm(v *OrderConfirm) {
	v.Feature = v.Feature[:0]
	v.ConfirmTime = ""
	v.DriverCarDesc = ""
	v.DriverCarName = ""
	v.DriverCarNo = ""
	v.DriverName = ""
	v.DriverTel = ""
	v.OrderId = ""
	v.ProviderId = ""
	v.ThirdOrderId = ""
	v.SellerId = ""
	v.DriverPic = ""
	v.CarPic = ""
	v.DriverTrumpetPhone = ""
	v.DriverIdNumber = ""
	v.SubPic = ""
	v.SubTitle = ""
	v.SubKey = ""
	v.DriverCarColor = ""
	v.CarBrand = ""
	v.CarModel = ""
	v.CarAge = ""
	v.DeriverGender = ""
	v.ConfirmType = 0
	v.UseType = 0
	v.CarTypeId = 0
	v.SupportRealTimePoi = false
	poolOrderConfirm.Put(v)
}
