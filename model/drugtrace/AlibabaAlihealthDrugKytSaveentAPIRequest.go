package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytSaveentAPIRequest
新增往来单位企业 API请求
alibaba.alihealth.drug.kyt.saveent

新增往来单位企业记录 */
type AlibabaAlihealthDrugKytSaveentAPIRequest struct {
	model.Params
	// 添加企业唯一标识
	_refEntId string
	// 新增企业信息
	_addEntReq *AddEntReqDto
	// 图片数据流。图片大小务必控制在2M以内
	_licPictureByte *model.File
}

// New
