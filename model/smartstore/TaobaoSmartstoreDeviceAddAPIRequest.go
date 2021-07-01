package smartstore

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSmartstoreDeviceAddAPIRequest
智慧门店设备创建 API请求
taobao.smartstore.device.add

智慧门店设备创建 */
type TaobaoSmartstoreDeviceAddAPIRequest struct {
	model.Params
	// mac地址
	_mac string
	// 设备在店内的位置，以文字描述
	_indoorPosition string
	// 设备名称
	_deviceName string
	// 门店ID
	_storeId int64
	// 操作系统类型：WINDOWS("WINDOWS", "WINDOWS"),     ANDROID("ANDROID", "ANDROID"),     IOS("IOS", "IOS"),     LINUX("LINUX", "LINUX"),     OTHER("OTHER", "OTHER");
	_osType string
	// 设备类型：     CAMERA("CAMERA", "客流摄像头"),     SHELF("SHELF", "云货架"),     MAKEUP_MIRROR("MAKEUP_MIRROR", "试妆镜"),     FITTING_MIRROR("FITTING_MIRROR", "试衣镜"),     VENDOR("VENDOR", "售货机"),     SAMPLE_MACHINE("SAMPLE_MACHINE","派样机"),     DOLL_MACHINE("DOLL_MACHINE", "娃娃机"),     INTERACTIVE_PHOTO("INTERACTIVE_PHOTO", "互动拍照"),     INTERACTIVE_GAME("INTERACTIVE_GAME", "互动游戏"),     USHER_SCREEN("USHER_SCREEN", "智慧迎宾屏"),     DRESSING("DRESSING", "闪电换装"),     MAGIC_MIRROR("MAGIC_MIRROR", "百搭魔镜"),     SHOES_FITTING_MIRROR("SHOES_FITTING_MIRROR", "试鞋镜"),     SKIN_DETECTION("SKIN_DETECTION", "肌肤测试仪"),     FOOT_DETECTION("FOOT_DETECTION", "测脚仪"),     RFID_SENSOR("RFID_SENSOR", "RFID"),touch_machine("touch_machine","导购一体屏")
	_deviceType string
	// 商家自定义设备编码
	_outerCode string
}

// New
