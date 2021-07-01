package fenxiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFenxiaoDealerRequisitionorderAgreeAPIResponse
供应商/分销商通过采购申请/经销采购单申请 API返回值
taobao.fenxiao.dealer.requisitionorder.agree

供应商或分销商通过采购申请/经销采购单审核 */
type TaobaoFenxiaoDealerRequisitionorderAgreeAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoDealerRequisitionorderAgreeAPIResponseModel
}

// TaobaoFenxiaoDealerRequisitionorderAgreeAPIResponseModel is 供应商/分销商通过采购申请/经销采购单申请 成功返回结果
type TaobaoFenxiaoDealerRequisitionorderAgreeAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_dealer_requisitionorder_agree_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作是否成功。true：成功；false：失败。
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
