package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthMedicalHospitalSyncAPIRequest
阿里将康支付宝挂号数据医院回传接口 API请求
alibaba.alihealth.medical.hospital.sync

阿里将康支付宝挂号数据医院回传接口 */
type AlibabaAlihealthMedicalHospitalSyncAPIRequest struct {
	model.Params
	// top保存入参
	_saveRequest *CommonRequest4Top
}

// New
