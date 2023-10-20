package koubeimall

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaokoubeimallcommonstorepageAPIResponse 分页查询综合体内的门店列表信息 API返回值
// taobao.koubei.mall.common.store.page
//
// 分页查询综合体内的门店列表信息
type TaobaokoubeimallcommonstorepageAPIResponse struct {
	model.CommonResponse
	TaobaokoubeimallcommonstorepageAPIResponseModel
}

// TaobaokoubeimallcommonstorepageAPIResponseModel is 分页查询综合体内的门店列表信息 成功返回结果
type TaobaokoubeimallcommonstorepageAPIResponseModel struct {
	XMLName xml.Name `xml:"koubei_mall_common_store_page_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// API接口返回的result模型
	Result *TaobaokoubeimallcommonstorepageResult `json:"result,omitempty" xml:"result,omitempty"`
}
