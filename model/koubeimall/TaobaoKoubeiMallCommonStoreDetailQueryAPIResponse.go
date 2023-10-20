package koubeimall

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaokoubeimallcommonstoredetailqueryAPIResponse 查询综合体内的门店详细信息 API返回值
// taobao.koubei.mall.common.store.detail.query
//
// 查询口碑综合体内的门店详情信息
type TaobaokoubeimallcommonstoredetailqueryAPIResponse struct {
	model.CommonResponse
	TaobaokoubeimallcommonstoredetailqueryAPIResponseModel
}

// TaobaokoubeimallcommonstoredetailqueryAPIResponseModel is 查询综合体内的门店详细信息 成功返回结果
type TaobaokoubeimallcommonstoredetailqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"koubei_mall_common_store_detail_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TaobaokoubeimallcommonstoredetailqueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
