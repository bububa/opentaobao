package traveltrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripTravelVisaApplicantUpdateAPIRequest
飞猪度假-普通签证-申请人进度推进接口 API请求
alitrip.travel.visa.applicant.update

普通签证订单的申请人进度推进接口，用于商家代用户填写申请人基本信息 或 推进单个申请人的签证进度。 */
type AlitripTravelVisaApplicantUpdateAPIRequest struct {
	model.Params
	// 必填，子订单id
	_subOrderId string
	// 必填，操作类型。1-上传新申请人基本信息（商家代填申请人），2-更新已有申请人基本信息，3-更新已有申请人的签证进度，4-商家代传申请人信息和材料(使馆直连订单)
	_operType int64
	// 特殊必填，申请人基本信息（证件号，姓名，手机号）列表。当operType为1或2或4时必填
	_applicantInfos []NormalVisaApplicantInfo
	// 特殊必填，签证申请人进度推进操作（目前只支持对单个申请人状态进行推进）。当operType为3时必填
	_applicantOp *NormalVisaApplicantOperation
	// 特殊必填，pdf文件字节流。用于上传电子签pdf结果 或者 预约面试信pdf文件。
	_fileBytes *model.File
	// 特殊必填，文件字节流，用于上传证件照，必须和photoType同时传
	_photoBytes *model.File
	// 证件照文件类型
	_photoType string
	// 特殊必填，文件字节流，用于上传护照，必须和passportType同时传
	_passportBytes *model.File
	// 护照文件类型
	_passportType string
	// 酒店预订文件类型
	_hotelBookingFormType string
	// 特殊必填，文件字节流，用于上传酒店预订，必须和hotelBookingFormType同时传
	_hotelBookingFormBytes *model.File
	// 机票预订文件类型
	_flightBookingFormType string
	// 特殊必填，文件字节流，用于上传机票预订，必须和flightBookingFormType同时传
	_flightBookingFormBytes *model.File
	// 特殊必填，更多材料
	_documentInfos []NormalVisaDocumentInfo
}

// New
