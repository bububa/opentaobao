package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthExaminationReserveConfirmAPIRequest
体检机构对接_体检套餐预定确认 API请求
alibaba.alihealth.examination.reserve.confirm

向体检机构确认用户购买的体检套餐信息 */
type AlibabaAlihealthExaminationReserveConfirmAPIRequest struct {
	model.Params
	// 商户唯一码
	_merchantCode string
	// 体检人姓名
	_name string
	// 阿里健康预约唯一标识
	_reserveNumber string
	// 性别(0-男;1-女;)
	_gender string
	// 出生日期
	_birthday string
	// 预约时间
	_reserveDate string
	// 体检套餐编码
	_packageCode string
	// 婚否(0-未婚; 1-已婚)
	_married string
	// 店铺ID
	_storeId string
	// 电话号码
	_phone string
	// 证件类型(0-身份证; 1-护照; 2-军官证)
	_certType string
	// 证件号
	_certNumber string
	// 所属公司
	_company string
	// 所属部门
	_department string
	// 报告邮寄地址
	_address string
	// 加项列表
	_addItems []AddItem
	// 加项包列表
	_addPacks []AddPack
	// 0没报告1有报告
	_havaReport string
	// 员工号
	_employeeNumber string
	// 服务类型，ONSITE_SERVICE（到店检测）、DOOR_TO_DOOR_SERVICE（上门检测）
	_serviceType string
	// 上门服务的上门地址
	_serviceAddress *AddAddress
	// 预约时间段开始
	_reserveTimeStart string
	// 预约时间段截止
	_reserveTimeEnd string
}

// New
