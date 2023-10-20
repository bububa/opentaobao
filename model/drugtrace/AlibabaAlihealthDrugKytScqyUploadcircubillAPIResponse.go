package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytscqyuploadcircubillAPIResponse 生成批发单据上传 API返回值
// alibaba.alihealth.drug.kyt.scqy.uploadcircubill
//
// 生产批发单据上传（非零售企业使用），包括101, &#34;生产入库&#34;；102, &#34;采购入库&#34;；103, &#34;退货入库&#34;；104, &#34;调拨入库&#34;；106, &#34;零头入库&#34;；107, &#34;供应入库&#34;；108, &#34;召回入库&#34;；110,&#34;赠品入库&#34;；111,&#34;盘盈入库&#34;；112,&#34;报废入库&#34;；113,&#34;其他入库&#34;
// 201, &#34;销售出库&#34;；202, &#34;退货出库&#34;；203, &#34;调拨出库&#34;；204, &#34;返工出库&#34;；205, &#34;销毁出库&#34;；206, &#34;抽检出库&#34;；207, &#34;直调出库&#34;；208, &#34;生产出库&#34;；209, &#34;供应出库&#34;；211, &#34;召回出库&#34;；212,&#34;赠品出库&#34;；214,&#34;盘亏出库&#34;；215,&#34;损坏出库&#34;；216,&#34;报废出库&#34;；217,&#34;其他出库&#34;；237, &#34;直调退货&#34;。
type AlibabaalihealthdrugkytscqyuploadcircubillAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthdrugkytscqyuploadcircubillAPIResponseModel
}

// AlibabaalihealthdrugkytscqyuploadcircubillAPIResponseModel is 生成批发单据上传 成功返回结果
type AlibabaalihealthdrugkytscqyuploadcircubillAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_scqy_uploadcircubill_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 返回结果
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回结果
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 是否成功(true 成功 ,false失败)
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}
