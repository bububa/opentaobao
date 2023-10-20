package wlb

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlbitemupdateAPIResponse 物流宝商品修改 API返回值
// taobao.wlb.item.update
//
// 修改物流宝商品信息
type TaobaowlbitemupdateAPIResponse struct {
	model.CommonResponse
	TaobaowlbitemupdateAPIResponseModel
}

// TaobaowlbitemupdateAPIResponseModel is 物流宝商品修改 成功返回结果
type TaobaowlbitemupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_item_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 修改时间
	GmtModified bool `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
}
