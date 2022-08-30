package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugtraceTopLsydUploadinoutbillAPIResponse 出入库单据上传 API返回值
// alibaba.alihealth.drugtrace.top.lsyd.uploadinoutbill
//
// 零售企业上传出入库信息，包括102, "采购入库"；103, "退货入库"；104, "调拨入库"；107, "供应入库"；108, "召回入库"；110,"赠品入库"；111,"盘盈入库"；112,"报废入库"；113,"其他入库"
// 201, "销售出库"；202, "退货出库"；203, "调拨出库"；204, "返工出库"；205, "销毁出库"；206, "抽检出库"；207, "直调出库"；209, "供应出库"；211, "召回出库"；212,"赠品出库"；214,"盘亏出库"；215,"损坏出库"；216,"报废出库"；217,"其他出库"；237, "直调退货"。
// 不包括对个人的零售出库，疫苗接种，领药出库。
type AlibabaAlihealthDrugtraceTopLsydUploadinoutbillAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugtraceTopLsydUploadinoutbillAPIResponseModel
}

// AlibabaAlihealthDrugtraceTopLsydUploadinoutbillAPIResponseModel is 出入库单据上传 成功返回结果
type AlibabaAlihealthDrugtraceTopLsydUploadinoutbillAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugtrace_top_lsyd_uploadinoutbill_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 返回编码(BILL_DECODE_ERROR 单据转码失败  BILL_FILE_NAME_DUPLICATE_UPLOAD 文件名重复)
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 是否成功(true 成功 false 失败)
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}
