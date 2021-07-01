package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaRetailDeviceVendingRegisterAPIRequest
贩卖机设备注册 API请求
alibaba.retail.device.vending.register

贩卖机注册 */
type AlibabaRetailDeviceVendingRegisterAPIRequest struct {
	model.Params
	// 设备名称
	_deviceName string
	// 设备地址
	_address string
	// 对接过阿里atm传入
	_deviceSn string
	// 业务编码，联系对接人申请
	_bizCode string
	// 外部设备编号
	_deviceUuid string
	// 设备类型
	_deviceModel string
	// COMMUNITY：小区,SCHOOL：学校,OFFICE：写字楼,SHOPPING_MALL：商场,AIRPORT：机场,SUBWAY：地铁,HOSPITAL：医院,PLAYGROUNDS：游乐场所,FACTORY：工厂,VIEWPOINT：旅游景点,OTHERS：其他
	_scene string
	// 场地名称，根据场地类型来，如：学校名称，商场名称。如果不传系统会根据address传入计算，为确保准确性请传入并确保address完整
	_siteName string
	// 楼栋信息。如果不传系统会根据address传入计算，为确保准确性请传入并确保address完整
	_floor string
	// 层。如果不传系统会根据address传入计算，为确保准确性请传入并确保address完整
	_layer string
	// 室内地址描述。如果不传系统会根据address传入计算，为确保准确性请传入并确保address完整
	_location string
}

// New
