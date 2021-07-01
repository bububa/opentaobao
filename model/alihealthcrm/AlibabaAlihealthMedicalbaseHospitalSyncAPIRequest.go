package alihealthcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest
互联网医院批量导入接口 API请求
alibaba.alihealth.medicalbase.hospital.sync

互联网医院isv批量通过接口批量导入 */
type AlibabaAlihealthMedicalbaseHospitalSyncAPIRequest struct {
	model.Params
	// 是否需要用户授权
	_isAuth string
	// 主院区纬度
	_lat string
	// 主院区经度
	_lon string
	// 主院区地址
	_hosAddress string
	// 主院区的联系电话
	_telephone string
	// 院区名称
	_regionName string
	// 是否公立医院（Y／N）
	_isPublic string
	// 标签
	_serviceInfo string
	// 自定义科室
	_special string
	// 生活号或者服务窗url
	_serviceWindowUrl string
	// 医院简介url
	_descriptionUrl string
	// 是否支持医保（Y/N）
	_isInsurance string
	// 医院等级
	_grade string
	// 综合(general)、专科（special）
	_category string
	// 医院简称
	_shortName string
	// 医院pid
	_pid string
	// 机构编码
	_unifyCode string
	// 所在城市code
	_cityCode string
	// 营业执照上的医院全称
	_hosName string
	// 公司名称
	_companyName string
	// 支付宝BD的姓名
	_aliInterfaceMan string
	// 邮箱地址
	_email string
	// 联系人
	_technicalMan string
	// 联系手机
	_phone string
	// 单医院（main）／ 平台（platform）
	_hosType string
	// 服务项列表
	_functions string
	// isv库里面的hosCode
	_isvHosCode string
	// 投放阵地alipay aliyy uc quark
	_deliveryChannel string
}

// New
