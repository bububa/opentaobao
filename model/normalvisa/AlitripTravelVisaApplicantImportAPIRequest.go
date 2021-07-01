package normalvisa

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripTravelVisaApplicantImportAPIRequest
签证申请人导入 API请求
alitrip.travel.visa.applicant.import

签证线下申请人导入接口。供商家将线下的签证申请人信息导入，进行签证线上化办理 */
type AlitripTravelVisaApplicantImportAPIRequest struct {
	model.Params
	// 国家id。目前支持越南(27027)
	_nationId int64
	// 证件照文件字节流
	_photoFile *model.File
	// 外部商家申请人id
	_outerApplyId string
	// 护照文件类型
	_passportFileType string
	// 护照文件字节流
	_passportFile *model.File
	// 证件照文件类型
	_photoFileType string
	// 申请人信息。字段注释：1.sex(性别),值:M/F;2.nationality(国籍),值:CHN(中国大陆),HKG(中国香港),MAC(中国澳门),USA(美国),CAN(加拿大);3.daibanTypeId(代办类型):1(越南一个月单次入境),2(越南一个月多次入境),3(越南三个月单次入境),4(越南三个月多次入境)
	_formDataJson string
}

// New
