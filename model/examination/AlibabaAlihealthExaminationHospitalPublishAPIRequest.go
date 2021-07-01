package examination

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthExaminationHospitalPublishAPIRequest
体检机构对接_门店发布／更新 API请求
alibaba.alihealth.examination.hospital.publish

第三方B端有新的门店发布，或者老的门店更新的时候，使用这个接口 */
type AlibabaAlihealthExaminationHospitalPublishAPIRequest struct {
	model.Params
	// 门店简介
	_detail string
	// 门店联系电话
	_tel string
	// 门店所属城市
	_cityName string
	// 门店城市code（国标）
	_cityCode string
	// 操作类型: publish=发布，update=更新
	_type string
	// 医院等级，三甲、
	_keyWord string
	// “须知”使用下面note_category字段
	_examNotice string
	// 门店位置经度高德 坐标系
	_pointX string
	// 门店位置纬度高德 坐标系
	_pointY string
	// 门店地址
	_address string
	// 工作时间
	_workTime string
	// 门店名称
	_hospitalName string
	// 门店code，机构保证唯一
	_hospitalCode string
	// 交通线路，通过\r\n 进行换行
	_routes string
	// http://images.aliyun.com/image?id=123
	_logo string
	// 是否支持在线报告。0:不支持;1:支持
	_onlineReport int64
	// 社会统一信用代码
	_socialCreditCode string
	// 线下报告获取说明（必填）
	_reportWay string
	// 线上体检报告几天出具（如果有电子报告必填）
	_reportWayOnline string
	// 环境图片(json字符串数组)，第一张是头图；（传图前先找运营同学要图片规范，别瞎传）
	_envImgsUrl string
	// 免费停车场,绿色VIP通道,免费早餐,3天出报告,1V1导检,接待引导,独家签约,专家会诊,当天出报告；多个逗号分隔
	_specialTagsCode string
	// 通知信息
	_notify string
	// 不同种类的预约须知；
	_noteCategory string
	// 经营模式 0自营模式、1平台模式
	_mode string
	// 门店与医院协议
	_agreement string
	// 营业执照
	_businessLicense string
	// 医疗经营许可
	_medicalLicense string
	// 类目:1=体检；2=核酸；3=上门；4=健康证；多个类目以逗号分割
	_category string
}

// New
