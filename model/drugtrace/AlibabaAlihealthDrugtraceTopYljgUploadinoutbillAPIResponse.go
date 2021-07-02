package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIResponse 出入库单据上传 API返回值
// alibaba.alihealth.drugtrace.top.yljg.uploadinoutbill
//
// 零售企业上传出入库信息，包括采购入库（102），退货入库（103），供应入库（107）,退货出库（202），销毁出库（205），抽检出库（206）， 供应出库（209）,
// 不包括对个人的零售出库，疫苗接种，领药出库。
type AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIResponseModel
}

// AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIResponseModel is 出入库单据上传 成功返回结果
type AlibabaAlihealthDrugtraceTopYljgUploadinoutbillAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugtrace_top_yljg_uploadinoutbill_response"`
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
