package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugLsydSaveentAPIRequest
零售药店往来单位新增 API请求
alibaba.alihealth.drug.lsyd.saveent

新增往来单位企业记录 */
type AlibabaAlihealthDrugLsydSaveentAPIRequest struct {
	model.Params
	// 添加企业唯一标识
	_refEntId string
	// 新增企业信息
	_addEntReq *AddEntReqDto
	// 图片数据流。图片大小务必控制在2M以内
	_licPictureByte *model.File
}

// New
