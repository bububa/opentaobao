package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangscitemdeleteAPIResponse 货品删除 API返回值
// alibaba.dchain.aoxiang.scitem.delete
//
// 货品删除
type AlibabadchainaoxiangscitemdeleteAPIResponse struct {
	model.CommonResponse
	AlibabadchainaoxiangscitemdeleteAPIResponseModel
}

// AlibabadchainaoxiangscitemdeleteAPIResponseModel is 货品删除 成功返回结果
type AlibabadchainaoxiangscitemdeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_scitem_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	DeleteScItemResponse *DeleteScItemResponse `json:"delete_sc_item_response,omitempty" xml:"delete_sc_item_response,omitempty"`
}
